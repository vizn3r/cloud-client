package http

type Conf struct {
	ServerHost string `yaml:"server_host"`
}

var cfg *Conf

func Init(conf *Conf) {
	cfg = conf
}
