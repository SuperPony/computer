package options

import (
	"encoding/json"
	"errors"

	"github.com/spf13/pflag"
)

type RedisOptions struct {
	Host  string `json:"host"`
	Port  string `json:"port"`
	Index uint   `json:"index"`
}

// AddFlags add redis flags
func (r *RedisOptions) AddFlags(fs *pflag.FlagSet) (err error) {
	if fs == nil {
		return errors.New("fs is nil")
	}

	fs.StringVar(&r.Host, "host", r.Host, "redis host")
	fs.StringVar(&r.Port, "port", r.Port, "redis port")
	fs.UintVar(&r.Index, "index", r.Index, "redis index")

	return nil
}

func (r *RedisOptions) String() string {
	res, _ := json.Marshal(r)
	return string(res)
}

// NewRedisOptions factory
func NewRedisOptions() *RedisOptions {
	return &RedisOptions{
		Host:  "127.0.0.1",
		Port:  "6379",
		Index: 0,
	}
}