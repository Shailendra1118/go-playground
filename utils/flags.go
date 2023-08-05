package utils

import (
	"bytes"
	"flag"
	"fmt"

	"os"

	"github.com/spf13/viper"
)

func main() {
	//this is defining
	//name := flag.String("name", "Devi Johns", "Please provide name")
	//isAdmin := flag.Bool("admin", false, "Specify if the user is an admin")

	flag.Parse()
	//fmt.Println("name is: ", *name)
	//fmt.Println("isAdmin:", *isAdmin)

	if err := loadConfig(); err != nil {
		fmt.Printf("error while loading config %v", err)
	}
}

var defaultConfig = []byte(`
release_service:
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
	os.Setenv("RELEASE_SERVICE_API_BASE_URL", "http://no-example.com")
	fmt.Println("loading...")

	v := viper.New()
	v.SetConfigType("yaml")

	data := defaultConfig
	//v.SetEnvPrefix("release_service_api")
	v.AutomaticEnv()
	var err error
	if err = v.ReadConfig(bytes.NewBuffer(data)); err != nil {
		fmt.Println("error: ", err)
		return err
	}

	fmt.Println("After reading config... ")

	cfgs := map[string]interface{}{
		"release_service.api":     &APIConfig{},
		"release_service.foo":     &FooConfig{},
		"release_service.secrets": &SecretConfig{},
	}

	fmt.Println("setting env variable..")

	fmt.Println("Getting keys", v.GetString("release_service.api.base_url"))
	//os.Setenv("RELEASE-SERVICE_API_BASE_URL", "localhost:9090")
	//fmt.Println("Getting keys after setting env", v.GetString("release-service.api.base_url"))
	fmt.Println("cfgs", cfgs)

	for key, cfg := range cfgs {
		fmt.Printf("marshling key %v and cfg %v\n", key, cfg)
		if err = v.UnmarshalKey(key, cfg); err != nil {
			return fmt.Errorf("unable to decode %s into struct due to %v", key, err)
		}
	}

	fmt.Println(cfgs["release_service.api"].(*APIConfig))
	fmt.Println(cfgs["release_service.foo"].(*FooConfig))
	fmt.Println(cfgs["release_service.secrets"].(*SecretConfig))
	return nil
}
