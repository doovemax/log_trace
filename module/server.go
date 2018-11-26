package module

type ServerConfig struct {
	ListenPort string `yaml:"ListenPort"`
	BindAddr   string `yaml:"BindAddr"`
}
