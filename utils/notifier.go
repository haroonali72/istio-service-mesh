package utils

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"istio-service-mesh/constants"
)

var (
	redisHost = constants.NotificationURL
)

type Notifier struct {
	Client *redis.Client
}

func (notifier *Notifier) Notify(channel, status string) {

	cmd := notifier.Client.Publish(channel, status)
	beego.Info(*cmd)
}

func (notifier *Notifier) Init_notifier() error {
	if notifier.Client != nil {
		return nil
	}
	redisHost = constants.NotificationURL

	options := redis.Options{}
	options.Addr = redisHost
	notifier.Client = redis.NewClient(&options)

	return nil
}
func (notifier *Notifier) receiver(channel, status string) {

	cmd := notifier.Client.Publish(channel, status)
	beego.Info(*cmd)
}