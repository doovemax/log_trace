package module

type ServerConfig struct {
	ListenPort string `yaml:"listenport"`
	BindAddr   string `yaml:"bindaddr"`
}
