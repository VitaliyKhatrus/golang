package main

import (
	"fmt"

	"github.com/spf13/viper"
)

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

// Create a new config instance.
var (
	conf *config
)

// Read the config file from the current directory and marshal
// into the conf config struct.
func getConf() *config {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}

// Initialization routine.
// func init() {
// 	// Retrieve config options.
// 	conf = getConf()
// }

// Main program.
func main() {
	// Print the config options from the new conf struct instance.
	// fmt.Printf("Hostname: %#v\n", conf.Hostname)
	// fmt.Printf("Port: %v\n", conf.Port)
	// fmt.Printf("ID: %v\n", conf.ID)
	// fmt.Printf("Username: %v\n", conf.Auth)

	fmt.Println(getConf())
}
