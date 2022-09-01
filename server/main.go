package main

import (
	"akexc.com/vueginlr-server/model"
	"akexc.com/vueginlr-server/router"
)

func main() {
	model.InitDB()
	err := router.InitRouter().Run(":5000")
	if err != nil {
		return
	}
}
