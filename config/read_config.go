package config

import (
	"sync"

	"fmt"
	"io/ioutil"

	"wordfrequency/gopkg.in/yaml.v2"

	"github.com/doovemax/log_trace/module"

	"github.com/ogier/pflag"
)

type ConfigAll struct {
	Hosts   map[string]module.Host      `yaml:"host"`
	LogFile map[string][]module.LogFile `yaml:"logfile"`
}

var (
	Conf       ConfigAll
	ConfigPath string
	mux        = sync.RWMutex{}
)

func Init() {
	pflag.StringVarP(&ConfigPath, "config", "f", "./config.yml", "指定配置文件(default: ./config.yml)")
	pflag.Parse()
	ParseConfig(ConfigPath)

}

func ParseConfig(configFile string) (err error) {
	fw, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	mux.Lock()
	err = yaml.Unmarshal(fw, &Conf)
	mux.Unlock()

	fmt.Printf("%+v\n", Conf)

	return
}
