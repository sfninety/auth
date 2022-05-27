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

	config = flag.String("config", "./local/config.yaml", "config file")
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

	cfg = &Config{}
	err = yaml.Unmarshal(b, cfg)
	log.Printf("\n ########## \n configuration loaded \n connection string: %v \n signing key: %v \n ##########", cfg.Datastore.ConnectionString, cfg.Handler.JwtSigningKey)
	return err
}

func attachRoutes(router *iris.Router) {
	router.GET("/health-check", healthCheck)
	router.POST("/reg-begin", Register)
	router.POST("/request-new-otp", RequestNewOTP)
	router.PUT("/verify-otp", VerifyOTP)
	router.POST("/finish-reg", FinishRegistration)
}

func healthCheck(r iris.Request) iris.Response {
	return r.ResponseWithCode("OK", 200)
}
