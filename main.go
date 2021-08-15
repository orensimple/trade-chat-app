package main

import (
	"fmt"

	"github.com/prometheus/common/log"

	"github.com/orensimple/trade-chat-app/internal/app/adapter"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("mongodb_host", "127.0.0.1")
	viper.SetDefault("mongodb_port", "27017")
	viper.SetDefault("mongodb_dbname", "trade")
	viper.SetDefault("app_port", "80")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error(err)
	}

	r := adapter.Router()
	port := viper.Get("app_port")
	err = r.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Error(err)
	}
}
