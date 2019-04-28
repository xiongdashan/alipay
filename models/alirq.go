package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
	"golang.org/x/text/encoding/simplifiedchinese"

	"github.com/Unknwon/goconfig"
	"github.com/otwdev/galaxylib"
)

type AliRQ struct {
	DataFile string
	Account  string
	//cookie     string
	param      string
	startTime  string
	endTime    string
	queryURL   string
	jar        http.CookieJar //*galaxylib.Jar
	alicookies []*http.Cookie
	ctoken     string
	username   string
	password   string
	retry      int
	billUserID string
}

func NewAliRQ(file string) *AliRQ {
	return &AliRQ{
		DataFile: file,
	}
}

//var ctoken string

func (a *AliRQ) loadConfg() {

	fmt.Println(a.Account)

	cfg, err := goconfig.LoadConfigFile(a.DataFile)
	if err != nil {
		galaxylib.GalaxyLogger.Error(err)
		return
	}

	//a.cookie = cfg.MustValue("rq", "cookie")
	a.param = cfg.MustValue("rq", "params")
	a.startTime = cfg.MustValue("rq", "startTime")
	a.endTime = cfg.MustValue("rq", "endTime")
	a.queryURL = cfg.MustValue("rq", "url")
	a.username = cfg.MustValue("rq", "username")
	a.password = cfg.MustValue("rq", "password")

}

func (a *AliRQ) RQ() {

	if a.jar == nil {

		a.loadConfg()

		//a.jar.SetCookies()

		//login := Login()

		a.alicookies = Login(a.username, a.password) //a.jar.BulkEdit(a.cookie)

		for _, v := range a.alicookies {
			if v.Name == "ctoken" {
				a.ctoken = v.Value
				//break
			}
			if v.Name == "ali_apache_tracktmp" {
				billID := strings.Split(v.Value, "=")
				a.billUserID = billID[1]
			}
		}
	}

	startTime := time.Now().Format("2006-01-02 00:00:00")
	endTime := time.Now().Add(24 * time.Hour).Format("2006-01-02 00:00:00")

	if len(a.startTime) > 0 {
		startTime = a.startTime
	}

	if len(a.endTime) > 0 {
		endTime = a.endTime
	}

	params, _ := url.ParseQuery(a.param)

	params.Set("startDateInput", startTime)
	params.Set("endDateInput", endTime)
	params.Set("billUserId", a.billUserID)

	// for _, c := range a.alicookies {

	// 	if c.Name == "ctoken" {
	params.Set("ctoken", a.ctoken)
	// 	}
	// }

	rq, _ := http.NewRequest("POST", a.queryURL, strings.NewReader(params.Encode()))

	section, _ := galaxylib.GalaxyCfgFile.GetSection("rqheader")

	for key, s := range section {
		rq.Header.Add(strings.TrimSpace(key), strings.TrimSpace(s))
	}

	if a.jar == nil {
		jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

		a.jar = jar

		a.jar.SetCookies(rq.URL, a.alicookies)
	}

	//a.jar.SetCookies(rq.URL, a.alicookies)
	c := http.Client{
		Jar: a.jar,
	}

	rs, err := c.Do(rq)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rs.Body.Close()

	buf, _ := ioutil.ReadAll(rs.Body)

	for _, c := range rs.Cookies() {

		if c.Name == "ctoken" {
			a.ctoken = c.Value
			fmt.Println(a.ctoken)
			break
		}
	}

	buf, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(buf)

	fmt.Println(string(buf))

	var trade *AliTradeItem

	json.Unmarshal(buf, &trade)

	if trade == nil {
		//data := galaxylib.DefaultGalaxyTools.Bytes2CHString(buf)

		fmt.Println(string(buf))
		return
	}

	if trade.Stat == "deny" {
		galaxylib.GalaxyLogger.Warnln(fmt.Sprintf("账号：%s超时", a.Account))

		if a.retry < 5 {

			a.retry++

			a.jar = nil

			a.alicookies = Login(a.username, a.password)

			a.RQ()
		}

		return
	}

	trade.Account = a.Account

	tradeBuf, _ := json.Marshal(trade)

	//galaxylib.DefaultGalaxyTools.Bytes2CHString(tradeBuf)
	//fmt.Println(string(tradeBuf))

	a.sendToAPI(tradeBuf)

	galaxylib.GalaxyLogger.Infoln(fmt.Sprintf("%s账号-数据：%d-查询条件:%s", a.Account, trade.Result.Summary.ExpendSum.Count, params.Encode()))

}

func (a *AliRQ) sendToAPI(t []byte) {

	apiURL := galaxylib.GalaxyCfgFile.MustValue("api", "url")
	res, err := http.Post(apiURL, "application/json", bytes.NewBuffer(t))
	if err != nil {
		galaxylib.GalaxyLogger.Errorln(err)
		return
	}
	defer res.Body.Close()

	//galaxylib.GalaxyLogger.Infoln(galaxylib.DefaultGalaxyTools.ResponseToString(res.Body))
}
