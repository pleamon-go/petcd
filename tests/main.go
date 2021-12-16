package main

import (
	"log"
	"time"

	"git.pleamon.com/p/pconfig"
	"git.pleamon.com/p/petcd"
)

type Config struct {
	Endpoints   []string
	DialTimeout time.Duration
}

var config *Config

func init() {
	pconfig.LoadYml("conf/config.yml", &config)
}

func main() {
	e, err := petcd.New(config.Endpoints...)
	if err != nil {
		panic(err)
	}
	putResp, err := e.Set("/hello/world_1", "test_1_1")
	if err != nil {
		panic(err)
	}
	log.Println("set world_1", putResp)
	putResp, err = e.Set("/hello/world_2", "test_2_2")
	if err != nil {
		panic(err)
	}
	log.Println("set world_2", putResp)
	setResp, err := e.Get("/hello/world_1")
	if err != nil {
		panic(err)
	}
	log.Println("get world_1", setResp.Kvs)
	setResp, err = e.GetPrefix("/hello/world")
	if err != nil {
		panic(err)
	}
	log.Println("get world_1", setResp.Kvs)
}
