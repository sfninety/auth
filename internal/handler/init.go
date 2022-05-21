package handler

import (
	"github.com/sfninety/iris"
)

type Config struct {
	JwtSigningKey string
}

var (
	cfg *Config
)

func Init(router *iris.Router) {

}
