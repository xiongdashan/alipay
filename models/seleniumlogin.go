package models

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/tebeka/selenium"
)

func Login(username, password string) []*http.Cookie {
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 4444))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	//wd.Wait()

	if err := wd.Get("https://auth.alipay.com/login/index.htm"); err != nil {
		fmt.Println(err)
		return nil
	}

	tabelm := getElm(wd, selenium.ByID, "J-loginMethod-tabs")
	lis, err := tabelm.FindElements(selenium.ByTagName, "li")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = lis[1].Click()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	u := getElm(wd, selenium.ByID, "J-input-user")
	u.Clear()

	//time.Sleep(3 * time.Second)

	// err = u.SendKeys("skyteam@zailushang.cn")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	slowInput(u, username)

	//time.Sleep(3 * time.Second)

	p := getElm(wd, selenium.ByID, "password_rsainput")
	p.Clear()

	// err = p.SendKeys("cgb101025")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	slowInput(p, password)

	//time.Sleep(3 * time.Second)

	// c := getElm(wd, selenium.ByID, "J-input-checkcode")

	// c.Clear()

	// //c.SendKeys("xdsl")

	// time.Sleep(3 * time.Second)

	b := getElm(wd, selenium.ByID, "J-login-btn")

	err = b.Click()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	time.Sleep(3 * time.Second)

	err = wd.Get("https://mbillexprod.alipay.com/enterprise/fundAccountDetail.htm")

	if err != nil {
		return nil
	}

	cookies, err := wd.GetCookies()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var retCookies []*http.Cookie

	for _, ck := range cookies {
		htpCookie := &http.Cookie{
			Name:   ck.Name,
			Value:  ck.Value,
			Domain: ck.Domain,
			Path:   ck.Path,
		}
		retCookies = append(retCookies, htpCookie)
		fmt.Printf("%s--%s\n", ck.Name, ck.Value)
	}

	return retCookies
	//fmt.Println(h)

}

func slowInput(elm selenium.WebElement, val string) {

	if elm == nil {
		return
	}

	for _, v := range val {

		elm.SendKeys(string(v))

		i := rand.Intn(500)

		//fmt.Println(i)

		time.Sleep(time.Duration(i) * time.Millisecond)
	}
}

func getElm(wd selenium.WebDriver, by, val string) selenium.WebElement {
	if wd == nil {
		fmt.Println("elm null.....")
		return nil
	}
	elm, err := wd.FindElement(by, val)
	if err != nil {
		fmt.Printf("%s---%v", val, err)
		return nil
	}
	return elm
}
