package initialize

import (
	"fmt"

	"github.com/anle/codebase/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")

	viper.SetConfigName("production")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Error reading config")
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to unmarshal config")
	}

}
