package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	Auth struct {
		AccessSecret         string // 生成jwt token的密钥，最简单的方式可以使用一个uuid值
		AccessExpire         int64  // jwt token有效期，单位：秒  1天=86400秒
		RememberAccessExpire int64  // jwt  记住登录时token有效期，单位：秒  3天=259200秒,  7天=604800秒
	}
}
