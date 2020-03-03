package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
)

var defaultConf = []byte(`
id: 5553321
core:
  environment: "staging"
  mode: "subversion"
api:
  port: ":8088"
  health_port: ":1488"
  metric_uri: "/metrics"
  health_uri: "/healthz"
  use_auth: true
  auth_login: "testuser"
  auth_password: "secret"
log:
  level: "info"
svn:
  svn_user: ""
  svn_password: ""
  svn_base_url: ""
git:
  git_token: ""
  git_base_url: ""
native:
  native_base_path: "/opt/app/config/"
  native_base_config: "/opt/app/application.properties"
auth:
  username: "Tom"
  password: "s99881"
`)

// Create private data struct to hold config options.
type config struct {
	Hostname string `yaml:"hostname"`
	Port     string `yaml:"port"`
	ID       string `yaml:"id"`
	Auth     auth   `yaml:"auth"`
}

type auth struct {
	username string `yaml:"username"`
	password string `yaml:"password"`
}

func LoadConf(confPath string) (config, error) {
	var conf config

	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix("go")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if confPath != "" {
		content, err := ioutil.ReadFile(confPath)

		if err != nil {
			return conf, err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return conf, err
		}
	} else {
		// Search config in home directory with name ".gorush" (without extension).
		// viper.AddConfigPath("/etc/go-config/server/")
		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
				return conf, err
			}
		}
	}
	conf.Auth.username = viper.GetString("auth.username")
	conf.Auth.password = viper.GetString("auth.password")
	conf.ID = viper.GetString("ID")

	return conf, nil

}

func main() {

	data, _ := LoadConf("")
	// fmt.Printf("username: %s", data.username)
	// fmt.Printf("password: %s", data.password)
	fmt.Printf("Username: %s\n", data.Auth.username)
	fmt.Printf("ID: %s\n", data.ID)
}
