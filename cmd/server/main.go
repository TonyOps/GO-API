package main

import "github.com/tonyops/GO-API/configs"

func main() {
	config, _ := configs.LoadConfig("cmd/server")
	println(config.DBDriver)
}
