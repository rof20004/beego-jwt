package routers

import (
	"strings"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/rof20004/beego-jwt/controllers"
)

// FilterAuth -> verify if client is authenticated
var FilterAuth = func(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/token") {
		return
	}

	j := &controllers.JwtController{}
	b := j.IsTokenValid(ctx.Request)
	if !b {
		message, _ := json.Marshal("Not authenticated")
		ctx.Output.SetStatus(403)
		ctx.Output.Body(message)
	}
}

func init() {
	beego.InsertFilter("/", beego.BeforeRouter, FilterAuth)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/token", &controllers.JwtController{}, "post:Auth")
}
