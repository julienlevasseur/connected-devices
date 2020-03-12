package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "plugs",
	Short: "A tiny programm to command plugs.",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

// Execute is the root function for Cobra.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func run() {
	// Init GPIO inputs
	//inputs.InitGPIO()

	duration := viper.GetInt("duration")

	fmt.Println(duration)

	for {
		select {
		case <-time.After(time.Duration(duration) * time.Second):
			fmt.Println("timed out")
			os.Exit(0)
		}
	}
}
