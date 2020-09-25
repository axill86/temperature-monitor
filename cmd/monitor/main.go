package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"temperature-monitor/internal/monitor/config"
	"temperature-monitor/internal/monitor/notification"
	"temperature-monitor/internal/monitor/temperature"
	"time"
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
		ticker := time.NewTicker(configuration.Interval)
		defer ticker.Stop()
		s := make(chan os.Signal)
		d := make(chan notification.Data)
		signal.Notify(s, syscall.SIGTERM, syscall.SIGINT)
		sender := notification.NewMqttSender(mqttClient, configuration.Mqtt.Topic)
		reader := temperature.NewReader()
		go notify(d, sender)
		go read(ticker.C, d, reader)
		//waiting till process finishes
		<-s
	},
}

func read(c <-chan time.Time, out chan<- notification.Data, reader temperature.Reader) {
	defer close(out)
	for t := range c {
		temperature := reader.Read()
		fmt.Println("Read temperature", temperature)
		out <- notification.Data{Timestamp: t, Temperature: temperature}
	}
}
func notify(c <-chan notification.Data, sender notification.Sender) {
	for data := range c {
		sender.Notify(data)
	}
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
