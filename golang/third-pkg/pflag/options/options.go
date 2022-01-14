package options

import (
	"encoding/json"
)

type Options struct {
	Mysql *MysqlOption  `json:"mysql"`
	Redis *RedisOptions `json:"redis"`
}

var NamedOptions *Options

func NewOptions() *Options {
	return &Options{
		Mysql: NewMysqlOptions(),
		Redis: NewRedisOptions(),
	}
}

func (o *Options) AddFlags() {
	o.Mysql.AddFlags(FlagGroup.FlagSet("mysql"))
	o.Redis.AddFlags(FlagGroup.FlagSet("redis"))
}

func (o *Options) String() string {
	res, _ := json.Marshal(o)
	return string(res)
}