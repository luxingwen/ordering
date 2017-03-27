package controllers

import (
	"ordering/models"

	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	token := this.GetString("token", "")
	fmt.Println("Prepare--->token", token)
	if token != models.Token {
		this.StopRun()
	}
}
