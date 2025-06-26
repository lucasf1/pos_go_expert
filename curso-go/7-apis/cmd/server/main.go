package main

import "github.com/lucasf1/goexpert/7-apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
