package server

const (
	ENDPOINT_USER_LOGIN = "user/login"
)

func v1(endpoint string) string {
	return cfg.Host + "/v1/" + endpoint
}
