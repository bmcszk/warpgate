package main

import (
	socks5 "github.com/armon/go-socks5"
	"github.com/spf13/viper"
)

type WarpgateConfig struct {
	Serve  ServeConfig
	Socks5 socks5.Config
}

type ServeConfig struct {
	Network string
	Address string
}

func main() {

	viper.SetConfigName("warpgate")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config WarpgateConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	server, err := socks5.New(&config.Socks5)
	if err != nil {
		panic(err)
	}

	if err := server.ListenAndServe(config.Serve.Network, config.Serve.Address); err != nil {
		panic(err)
	}

}
