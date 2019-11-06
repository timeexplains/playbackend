package main

import (
	"fmt"
	"github.com/astaxie/beego/plugins/cors"
)
import "github.com/astaxie/beego"
import _ "god/routers"
import "god/models"

func  main()  {
	fmt.Println("HELLO GO")
	fmt.Println(models.Fibnonacci_add(3,4))


	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Run()
}

