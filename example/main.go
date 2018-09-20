package main

import (
	"time"

	"github.com/astaxie/beego"
	"go.uber.org/zap"
	".."
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	logger, _ := zap.NewProduction()
	beezap.InitBeeZapMiddleware(logger, time.RFC3339, true, true)
	beego.Router("/", &MainController{})
	beego.Run(":8090")
}