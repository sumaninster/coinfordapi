package main

import (
	_ "coinford_api/routers"
	"coinford_api/configs"
	//"coinford_api/process"

	"github.com/astaxie/beego"
)

func init() {
    configs.Init()
}

func api() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func main() {
	//go process.ProcessOrders()
	//btcprocess := process.BTCProcess{}
	//go btcprocess.Process()
	//fmt.Println(configs.RandString(30))
	//ethprocess := process.ETHProcess{}
	//go ethprocess.Process()
	api()
}