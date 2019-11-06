package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
		c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	}

	bodystr, err := ioutil.ReadFile("front/index.html");

	if( err != nil ) {
		c.Ctx.WriteString("server is wrong")
		return
	}

	c.Ctx.Output.Body(bodystr)
}