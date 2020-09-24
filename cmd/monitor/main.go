package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"temperature-monitor/internal/monitor/config"
	"temperature-monitor/internal/monitor/notification"
)

var configFile string
var configuration *config.Config = &config.Config{}
var rootCmd = &cobra.Command{
	Use:   "monitor",
	Short: "monitor sends temperature notifications to mqtt",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting temperature monitor")
		mqttClient, err := config.CreateMqttClient(&configuration.Mqtt)
		if err != nil {
			panic(err)
		}
		notifier := notification.NewMqttSender(mqttClient, configuration.Mqtt.Topic)
		notifier.Notify(notification.Data{
			DeviceName:  configuration.DeviceName,
			Temperature: 0,
		})
	},
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	viper.Unmarshal(configuration)
	fmt.Println("Reading configuration from", viper.ConfigFileUsed())
}

func main() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "Configuration file")
	rootCmd.MarkFlagRequired("config")
	cobra.OnInitialize(initConfig)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
