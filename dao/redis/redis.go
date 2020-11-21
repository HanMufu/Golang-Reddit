package redis

import (
	"fmt"
	"go-web-app/settings"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

// 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	//client = redis.NewClient(&redis.Options{
	//	Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
	//	Password: viper.GetString("redis.password"), // no password set
	//	DB:       viper.GetInt("redis.db"),          // use default DB
	//	PoolSize: viper.GetInt("redis.pool_size"),   // 连接池大小
	//})

	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	return
}

func Close() {
	_ = client.Close()
}

//func V8Example() {
//	ctx := context.Background()
//	if err := initClient(); err != nil {
//		return
//	}
//
//	err := client.Set(ctx, "key", "value", 0).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	val, err := client.Get(ctx, "key").Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("key", val)
//
//	val2, err := client.Get(ctx, "key2").Result()
//	if err == redis.Nil {
//		fmt.Println("key2 does not exist")
//	} else if err != nil {
//		panic(err)
//	} else {
//		fmt.Println("key2", val2)
//	}
//	// Output: key value
//	// key2 does not exist
//}
