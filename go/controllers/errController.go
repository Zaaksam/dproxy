package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type ErrController struct {
	beego.Controller
}

func (c *ErrController) Error404() {
	// c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	// c.Ctx.Output.SetStatus(404)
	// c.Ctx.Output.Body([]byte("404"))
	http.Error(c.Ctx.ResponseWriter, "", 404)
}

func (c *ErrController) Error500() {
	// c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	// c.Ctx.Output.SetStatus(500)
	// c.Ctx.Output.Body([]byte("500"))
	http.Error(c.Ctx.ResponseWriter, "", 500)
}
