package module

type LogFile struct {
	Path     string `yaml:"path"`
	Host     string `yaml:"host"`
	LogTrans chan []byte
}
