package helper

import (
	"github.com/spf13/viper"
)

func NewViper() {
	viper.SetConfigName("gowatch")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		// if viper.GetString("ENV") != "PROD" {
		// 	panic(fmt.Errorf("fatal error config file: %w", err))
		// }
	}
}
