package main

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"qixiang/email"
	"strings"
)

type Conf struct {
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
	Host     string `yaml:"host"`
}

var conf *Conf
var target *To

type From struct {
	City string   `json:"city"`
	User []string `json:"user"`
}

type To struct {
	From []From `json:"from"`
}

func init() {
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configFile, &conf)
	if err != nil {
		panic(err)
	}

	jsonByte, err := ioutil.ReadFile("recipient.json")
	if err != nil {
		panic(err)
	}

	var to To
	err = json.Unmarshal(jsonByte, &to)
	if err != nil {
		panic(err)
	}

	target = &to

	//router.Star()

}

func main() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 35 9 * * ?", func() {
		for _, v := range target.From {
			email.SendMail(v.City, conf.UserName, conf.PassWord, conf.Host, strings.Join(v.User, ";"))
		}
	})
	fmt.Printf("service start... \n")
	c.Start()
	select {}
}
