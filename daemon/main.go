package main

import (
	_ "embed"

	"github.com/vizn3r/cloud/daemon/server"
	"github.com/vizn3r/cloud/lib/conf"
	"github.com/vizn3r/cloud/lib/logger"
)

type GlobalConfig struct {
	Server *server.ServerConfig `yaml:"server"`
}

var (
	log = logger.New("MAIN", logger.Cyan)
	cfg *GlobalConfig
)

//go:embed config.test.yaml
var rawConfig []byte

func main() {
	if err := conf.LoadFromBytes[GlobalConfig](rawConfig, "yaml"); err != nil {
		log.Fatal(err)
	}

	cfg = conf.Get[GlobalConfig]()

	server.Init(conf.Get[GlobalConfig]().Server)

	log.Print(cfg.Server.User.Email)

	if err := server.Authenticate(cfg.Server.User.Email, "test"); err != nil {
		log.Fatal(err)
	}

	log.Info("Hello world")
	log.Close()
}
