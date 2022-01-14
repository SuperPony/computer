package options

import (
	"encoding/json"
	"errors"

	"github.com/spf13/pflag"
)

type MysqlOption struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// AddFlags add redis flags
func (m *MysqlOption) AddFlags(fs *pflag.FlagSet) error {
	if fs == nil {
		return errors.New("fs is nil")
	}

	fs.StringVar(&m.Host, "host", m.Host, "mysql host")
	fs.StringVar(&m.Port, "port", m.Port, "mysql port")
	fs.StringVar(&m.User, "user", m.User, "mysql user")
	fs.StringVar(&m.Password, "password", m.Password, "mysql password")

	return nil
}

func (m *MysqlOption) String() string {
	res, _ := json.Marshal(m)
	return string(res)
}

// NewMysqlOptions MysqlOptions factory
func NewMysqlOptions() *MysqlOption {
	return &MysqlOption{
		Host:     "127.0.0.1",
		Port:     "6379",
		User:     "root",
		Password: "123456",
	}
}