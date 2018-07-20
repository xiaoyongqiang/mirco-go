package config

type MysqlConf struct {
	Host      string `default:"192.168.7.205"`
	Port      int    `default:"4000"`
	User      string `default:"root"`
	Pass      string `default:"123456"`
	Name      string `default:"swing_baby"`
	Charset   string `default:"utf8"`
	IdleConns int    `default:"200"`
	OpenConns int    `default:"500"`
}
