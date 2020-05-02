package main

import (
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/mitchellh/mapstructure"
	"github.com/personal-website-go/conf"
	"github.com/personal-website-go/dep"
	"github.com/personal-website-go/router"
)

func init() {
	initConfig()
	dep.InitDep()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		router.AddNamespace()
	}
	beego.Run()
}

func initConfig() {
	cwd, _ := os.Getwd()
	configFilePath := cwd + "/resources/config.json"
	err := beego.LoadAppConfig("json", configFilePath)
	if err != nil {
		log.Panicf("Error in loading config from file : %v", err.Error())
	}
	initBeegoConfig()
	initAppConfig()
}

func initBeegoConfig() {
	beegoConfig, err := beego.AppConfig.DIY("beegoConfig")
	if err != nil {
		log.Panicf("App config not found : %v", err.Error())
	}
	beegoConfiguration := &beego.Config{}
	mapstructure.Decode(beegoConfig, beegoConfiguration)
	beego.BConfig = beegoConfiguration
}

func initAppConfig() {
	appConfig, err := beego.AppConfig.DIY("appConfig")
	if err != nil {
		log.Panicf("App config not found : %v", err.Error())
	}
	appConfiguration := &conf.AppConfig{}
	mapstructure.Decode(appConfig, appConfiguration)
	conf.RegisterConfig(appConfiguration)
	if err != nil {
		log.Panicf("Error in initialising config : %v", err.Error())
	}
}
