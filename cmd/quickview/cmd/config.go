package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configKeyAddress = "address"
)

var config struct {
	Address string `mapstructure:"address"`
}

func init() {
	cobra.OnInitialize(func() {
		err := viper.Unmarshal(&config)

		if err != nil {
			panic(err)
		}
	})
}
