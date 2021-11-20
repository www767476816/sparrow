package frame

import (
	"github.com/go-redis/redis"
)

func (this* Frame) InitRedis() {
	this.redisList=make([]*redis.Client, 16)
}

func (this* Frame) StartRedis() {
	for i:=0;i<16;i++{
		redisAddr:=this.normalConfig.Redis.Ip+":"+this.normalConfig.Redis.Port
		client := redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: this.normalConfig.Redis.Password,
			DB:       i,
		})

		if _, err := client.Ping().Result();err!=nil{
			this.log.Panicln(err)
		}
		this.redisList[i]=client
	}
}
func (this* Frame) CloseRedis() {
	for _,cli:=range this.redisList{
		cli.Close()
	}
	this.redisList=this.redisList[0:0]
}