package routers

import "github.com/astaxie/beego"
import "god/controllers"

func init()  {
	beego.Router("/",&controllers.MainController{})
	beego.Router("/user",&controllers.UserController{},"Get:GetUser")
}

