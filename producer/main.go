package main

import (
	"producer/controller"
	"producer/services"
	"strings"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	eventProducer := services.NewEventProducer(producer)
	accountService := services.NewAccountService(eventProducer)
	accountController := controller.NewAccountController(accountService)

	app := fiber.New()

	app.Post("/openAccount", accountController.OpenAccount)
	app.Post("/depositFund", accountController.Deposit)
	app.Post("/withdrawFund", accountController.Withdraw)
	app.Post("/closeAccount", accountController.CloseAccount)

	app.Listen(":8000")
}
