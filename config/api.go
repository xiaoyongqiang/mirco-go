package config

// ApiConf is stored the api listen host and port
type ApiConf struct {
	Host string `default:"api.beibei1.butup.me"`
	Port int    `default:"8080"`
}
