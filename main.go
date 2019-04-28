package main

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/otwdev/alipay/models"
	"github.com/otwdev/galaxylib"
)

//var alicookie *cookiejar.Jar

func main() {

	galaxylib.DefaultGalaxyConfig.InitConfig()
	galaxylib.DefaultGalaxyLog.ConfigLogger()

	//bufio.NewReader(os.Stdin).ReadLine()

	// aliRQ := &models.AliRQ{}

	// for {
	// 	aliRQ.RQ()
	// 	time.Sleep(1 * time.Minute)
	// }

	//login()

	//models.Login()

	crawl()

}

func crawl() {
	c := make(chan int)

	filepath.Walk("./account", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, "ini") {
			name := strings.TrimSuffix(info.Name(), ".ini")

			go func(p, n string) {
				execRq(p, n)
			}(path, name)

		}
		return nil
	})

	<-c
}

func execRq(file, name string) {

	aliRQ := &models.AliRQ{
		DataFile: file,
	}
	aliRQ.Account = name

	for {

		aliRQ.RQ()

		time.Sleep(1 * time.Minute)
	}

}
