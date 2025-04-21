package initializer

import (
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

const (
	OTPTopic = "otp-auth-topic"
)

func InitKafka() {
	kafkaConfig := global.Config.Kafka
	global.Kafka = &kafka.Writer{
		Addr:     kafka.TCP(fmt.Sprintf("%s:%v", kafkaConfig.Host, kafkaConfig.Port)),
		Topic:    OTPTopic,
		Balancer: &kafka.LeastBytes{},
	}
}

func CloseKafka() {
	if err := global.Kafka.Close(); err != nil {
		global.Logger.Error("Failed to close kafka producer: %v", zap.Error(err))
	}
	global.Logger.Info("Close kafka producer success")
}
