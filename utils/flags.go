package main

import (
	"bytes"
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	//this is defining
	name := flag.String("name", "Devi Johns", "Please provide name")
	isAdmin := flag.Bool("admin", false, "Specify if the user is an admin")

	flag.Parse()
	fmt.Println("name is: ", *name)
	fmt.Println("isAdmin:", *isAdmin)

	if err := loadConfig(); err != nil {
		fmt.Printf("error while loading config %v", err)
	}
}

var defaultConfig = []byte(`
release-service:
 api:
  base_url: http://localhost:8080
  access_token: token
  timeout: 10
 foo:
  bar: foobar
  again: yetagain
  yet_again: oh_again
 secrets:
   details: moredetails
   path: src/other/deploy
`)

type APIConfig struct {
	BaseUrl      string `mapstructure:"base_url"`
	Access_Token string `mapstructure:"access_token"`
	Timeout      string `mapstructure:"timeout"`
}

type FooConfig struct {
	Bar      string `mapstructure:"bar"`
	Again    string `mapstructure:"again"`
	YetAgain string `mapstructure:"yet_again"`
}

type SecretConfig struct {
	Details string `mapstructure:"details"`
	Path    string `mapstructure:"path"`
}

func loadConfig() error {
	fmt.Println("loading...")

	v := viper.New()
	v.SetConfigType("yaml")

	//data := defaultConfig
	var err error
	if err = v.ReadConfig(bytes.NewBuffer(defaultConfig)); err != nil {
		fmt.Println("error: ", err)
		return err
	}

	fmt.Printf("After reading config... ")

	cfgs := map[string]interface{}{
		"release-service.api":     &APIConfig{},
		"release-service.foo":     &FooConfig{},
		"release-service.secrets": &SecretConfig{},
	}

	for key, cfg := range cfgs {
		if err = v.UnmarshalKey(key, cfg); err != nil {
			return fmt.Errorf("unable to decode %s into struct due to %v", key, err)
		}
	}

	fmt.Println(cfgs["release-service.api"].(*APIConfig))
	fmt.Println(cfgs["release-service.foo"].(*FooConfig))
	fmt.Println(cfgs["release-service.secrets"].(*SecretConfig))
	return nil
}
