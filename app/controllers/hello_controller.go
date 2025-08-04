package controllers

import (
	"github.com/go-raptor/raptor/v4"
	"github.com/go-raptor/template/app/services"
)

type HelloController struct {
	raptor.Controller

	Hello *services.HelloService
}

func (hc *HelloController) Greet(c *raptor.Context) error {
	greeting := map[string]string{
		"message": hc.Hello.Greeting(),
	}

	return c.Data(greeting)
}
