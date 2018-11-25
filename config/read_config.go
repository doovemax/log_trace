package config

import "github.com/ogier/pflag"

var ConfigPath string

func Init() {
	pflag.StringVarP(&ConfigPath, "config", "f", "./config.yml", "指定配置文件(default: ./config.yml)")
	pflag.Parse()

}
