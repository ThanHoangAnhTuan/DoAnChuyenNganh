package initializer

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
)

func LoadConfig() {
	// file config in /configs/local.yaml
	viper := viper.New()

	viper.SetConfigName("local")      // name of config file (without extension)
	viper.SetConfigType("yaml")       // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./configs/") // path to look for the config file in
	err := viper.ReadInConfig()       // Find and read the config file
	if err != nil {                   // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// map file config to struct
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v\n", err)
	}
}
