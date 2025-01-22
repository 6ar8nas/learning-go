package api

type APIConfig struct {
	Host string
	Port string
}

var Env = func() *APIConfig {
	return &APIConfig{
		Host: "api",
		Port: "8080",
	}
}()
