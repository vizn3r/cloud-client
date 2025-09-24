package server

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/vizn3r/go-lib/logger"
)

type ServerConfig struct {
	Host string `yaml:"host"`
	User struct {
		Email string `yaml:"email"`
	} `yaml:"user"`
}

var (
	log = logger.New("SRVR", logger.Purple)
	cfg *ServerConfig
)

func Init(conf *ServerConfig) {
	cfg = conf
}

func Authenticate(email, password string) error {
	client := &http.Client{}
	data := "{\"email\":\"" + email + "\",\"password\":\"" + password + "\"}"
	reqData := strings.NewReader(data)

	req, err := http.NewRequest("POST", v1(ENDPOINT_USER_LOGIN), reqData)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(data)))

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	resData := string(bytes)

	log.Info("Response: ", resData)

	log.Close()
	return nil
}
