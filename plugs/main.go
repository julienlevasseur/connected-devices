package main

import (
	"github.com/julienlevasseur/connected-devices/plugs/cmd"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("cd_plugs") // will be uppercased automatically
	viper.SetDefault("loglevel", "Info")
}

func main() {
	initConfig()
	cmd.Execute()
}
