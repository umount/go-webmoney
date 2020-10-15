package main

import (
	"fmt"
	"github.com/C0nstantin/go-webmoney"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
)

func main() {
	date, err := ioutil.ReadFile("examples/conf.toml")
	if err != nil {
		log.Fatal(err)
	}

	config, _ := toml.Load(string(date))
	wmClient := webmoney.WmClient{
		Wmid: config.Get("client1.Wmid").(string),
		Key:  config.Get("client1.Key").(string),
		Pass: config.Get("client1.Pass").(string),
	}

	test := webmoney.TestSignRequest{
		Wmid: "128756507061",
		Plan: "-",
		Sign: "-",
	}

	result, err := wmClient.TestSign(test)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%#v", result)
}
