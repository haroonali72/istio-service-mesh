package utils
import (
	"github.com/go-redis/redis"
	"github.com/astaxie/beego"
)

var (
	redisHost    = beego.AppConfig.String("redis_url")
)

type Notifier struct {
	Client *redis.Client
}
func (notifier *Notifier)  notify(channel, status string){

	cmd :=notifier.Client.Publish(channel,status)
	beego.Info(*cmd)
}

func (notifier *Notifier) init_notifier() error {
	if notifier.Client != nil {
		return nil
	}

	options := redis.Options{}
	options.Addr = redisHost
	notifier.Client  = redis.NewClient(&options)

	return nil
}
func (notifier *Notifier)  receiver(channel, status string){

	cmd :=notifier.Client.Publish(channel,status)
	beego.Info(*cmd)
}
