package config

import (
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

type mallConf struct {
	Title  string `toml:"title"`
	Server struct {
		Port string `toml:"port"`
	} `toml:"server"`
	Token struct {
		Secret string `toml:"secret"`
	} `toml:"token"`
	Datebases struct {
		Alpha struct {
			Host          string `toml:"host"`
			ConnectionMax int    `toml:"connection_max"`
			User          string `toml:"user"`
			Password      string `toml:"password"`
			DbName        string `toml:"db_name"`
			Ssl           string `toml:"ssl"`
			Schema        string `toml:"schema"`
		} `toml:"alpha"`
		Beta struct {
			Host          string `toml:"host"`
			ConnectionMax int    `toml:"connection_max"`
			User          string `toml:"user"`
			Password      string `toml:"password"`
			DbName        string `toml:"db_name"`
			Ssl           string `toml:"ssl"`
			Schema        string `toml:"schema"`
		} `toml:"beta"`
		Release struct {
			Host          string `toml:"host"`
			ConnectionMax int    `toml:"connection_max"`
			User          string `toml:"user"`
			Password      string `toml:"password"`
			DbName        string `toml:"db_name"`
			Ssl           string `toml:"ssl"`
			Schema        string `toml:"schema"`
		} `toml:"release"`
	} `toml:"datebases"`
}

var (
	conf    *mallConf
	once    sync.Once
	cfgLock sync.RWMutex
)

func New() *mallConf {
	once.Do(ParseConfig)
	cfgLock.RLock()
	defer cfgLock.RUnlock()
	return conf
}

// 解析toml文件
func ParseConfig() {
	path, err := filepath.Abs("./config/mall.toml")
	if err != nil {
		panic(err)
	}
	config := new(mallConf)
	cfgLock.Lock()
	defer cfgLock.Unlock()
	if _, err := toml.DecodeFile(path, config); err != nil {
		panic(err)
	}
	conf = config
	return
}
