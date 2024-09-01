package main

import (
	"consumer/repositories"
	"consumer/services"
	"context"
	"events"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM/sarama"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Version = sarama.V2_6_0_0

	consumerGroup, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)
	if err != nil {
		panic(err)
	}

	defer consumerGroup.Close()

	db := initDatabase()
	accountRepo := repositories.NewAccountRepo(db)
	accountEventHandler := services.NewAccountEventHandler(accountRepo)
	accountConsumerHandler := services.NewConsumerHandler(accountEventHandler)

	//log.Println("account consumer start")
	// for {
	// 	consumerGroup.Consume(context.Background(), events.Topics, accountConsumerHandler)
	// }

	log.Println("account consumer start")
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	for {
		err := consumerGroup.Consume(context.Background(), events.Topics, accountConsumerHandler)
		if err != nil {
			log.Fatalf("Failed to consume messages: %s", err)
		}
		log.Println("messages!!!")
	}
	// go func() {
	// 	for {
	// 		err := consumerGroup.Consume(ctx, events.Topics, accountConsumerHandler)
	// 		if err != nil {
	// 			log.Fatalf("Failed to consume messages: %s", err)
	// 		}
	// 		// Check if context was canceled, signaling to stop
	// 		if ctx.Err() != nil {
	// 			return
	// 		}
	// 	}
	// }()

	// // Wait for signal to exit
	// sigterm := make(chan os.Signal, 1)
	// signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	// <-sigterm

}
func initDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.user"),
		viper.GetString("db.pass"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.db_name"))

	fmt.Println(dsn)
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold: time.Second, // Slow SQL threshold
	// 		LogLevel:      logger.Info, // Log ALL level
	// 		Colorful:      true,        // Disable color
	// 	},
	// )
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ /*Logger: newLogger*/ })
	if err != nil {
		panic("failed to connect to database")
	}
	postgres, err := db.DB()
	if err != nil {
		panic("failed to connect to database")
	}
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	postgres.SetMaxOpenConns(10)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	postgres.SetConnMaxLifetime(3 * time.Minute)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	postgres.SetMaxIdleConns(10)
	fmt.Println("Connected !!")
	return db
}
