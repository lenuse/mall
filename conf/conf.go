package conf

import (
	"basic_mall/utils"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

type DBConf struct {
	Database     string `yaml:"database"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"pwd"`
	DBName       string `yaml:"dbname"`
	Log          bool   `yaml:"log"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
}

type JWTConf struct {
	JwtSecret string `yaml:"jwt_secret"`
}

type LogConf struct {
	LogFilePath string `yaml:"log_file_path,omitempty"`
}

type Conf struct {
	DBConf
	JWTConf
	LogConf
}

var conf Conf
var once sync.Once

func init() {
	New()
}

func New() Conf {
	once.Do(func() {
		fileInfo, err := ioutil.ReadFile("conf.yaml")
		utils.CheckErr(err)
		err = yaml.Unmarshal(fileInfo, &conf)
		utils.CheckErr(err)
	})
	return conf
}

func GetHost() string {
	hostMap := map[string]string{
		"postgres": "host={host} port={port} user={user} password={password} dbname={dbname} sslmode=disable",
		"mysql":    "{user}:{password}@tcp({host}:{port})/{dbname}?charset=utf8&parseTime=True&loc=Local",
	}
	var hostURL string
	if host, ok := hostMap[conf.Database]; ok {
		hostURL = host
	} else {
		fmt.Println("没有找到改驱动")
		os.Exit(1)
	}
	hostURL = strings.Replace(hostURL, "{user}", conf.User, -1)
	hostURL = strings.Replace(hostURL, "{password}", conf.Password, -1)
	hostURL = strings.Replace(hostURL, "{host}", conf.Host, -1)
	hostURL = strings.Replace(hostURL, "{port}", conf.Port, -1)
	hostURL = strings.Replace(hostURL, "{dbname}", conf.DBName, -1)
	return hostURL
}

func GetJwtSecret() []byte {
	return []byte(conf.JwtSecret)
}
