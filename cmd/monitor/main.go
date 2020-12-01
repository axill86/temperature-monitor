package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"temperature-monitor/internal/monitor/config"
	"temperature-monitor/internal/monitor/notification"
	"temperature-monitor/internal/monitor/temperature"
)

var configFile string
var configuration *config.Config = &config.Config{}
var rootCmd = &cobra.Command{
	Use:   "monitor",
	Short: "monitor sends temperature notifications to mqtt",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting temperature monitor")

		sender, err := notification.NewMqttSender(
			[]byte(configuration.Mqtt.RootCa),
			configuration.Mqtt.Username,
			configuration.Mqtt.Password,
			configuration.Mqtt.BrokerUrl,
			configuration.Mqtt.Topic)

		reader, err := temperature.NewIntervalReader(configuration.Interval,
			configuration.I2C.Bus,
			configuration.I2C.Address,
			func(measurement temperature.Measurement) {
				log.Printf("Measured  %v", measurement)
				err = sender.Notify(notification.Data{
					Temperature: measurement.Temp,
					Humidity:    measurement.Humidity,
					Timestamp:   measurement.Timestamp,
				})
				if err != nil {
					log.Println("Seeing error...", err)
				}
			})
		if err != nil {
			panic(err)
		}
		defer reader.Stop()
		reader.Start()
		//interruption channel
		s := make(chan os.Signal)
		signal.Notify(s, syscall.SIGTERM, syscall.SIGINT)
		//waiting till process finishes
		<-s
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
