package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
)


type UserController struct{
	beego.Controller
}

type User struct {
	Name string `json:"name"`
	Age  int32	`json:"age"`
}
type UserList struct {
	List []User `json:"items"`
}

func (ctrl *UserController) GetUser() {
	res := UserList{}
	for i:=0; i< 10; i++ {
		user := User{
			Name: "owen",
			Age:  10,
		}
		res.List = append(res.List,user)
	}

	buf,_ := json.Marshal(res)
	log.Printf(string(buf))

	ctrl.Data["json"] = res
	ctrl.ServeJSON()
}