package main

import "github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/08-apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
