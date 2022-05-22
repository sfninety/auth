package handler

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/sfninety/auth/internal/datastore"
	"github.com/sfninety/iris"
	"gopkg.in/yaml.v2"
)

type HandlerConfig struct {
	JwtSigningKey string `yaml:"jwt_signing_key"`
}

type Config struct {
	Handler   *HandlerConfig    `yaml:"handler"`
	Datastore *datastore.Config `yaml:"datastore"`
}

var (
	cfg *Config

	config = flag.String("config", "/etc/config.yaml", "config file")
)

func Init(router *iris.Router) {
	err := loadConfig()
	if err != nil {
		panic(err)
	}

	attachRoutes(router)

	datastore.Init(cfg.Datastore)
}

func loadConfig() error {
	b, err := ioutil.ReadFile(*config)
	if err != nil {
		log.Printf("error reading config file: %v", err)
		return err
	}

	err = yaml.Unmarshal(b, cfg)
	return err
}

func attachRoutes(router *iris.Router) {
}
