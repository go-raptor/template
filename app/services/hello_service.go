package services

import (
	"github.com/go-raptor/raptor/v4"
)

type HelloService struct {
	raptor.Service
}

func (hs *HelloService) Greeting() string {
	return "Hello, World!"
}
