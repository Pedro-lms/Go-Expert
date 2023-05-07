package main

import "github.com/Pedro-lms/goexpert/9-APIs/configs"

func main() {
	config, err := configs.LoadConfig(".")
	println(config.DBDriver)
	if err != nil {
		panic(err)
	}
}
