package config

import (
	"log"

	"github.com/go-ini/ini"
)

type (
	Server struct {
		RunMode  string
		HttpPort int
	}

	Database struct {
		Type         string
		Username     string
		Password     string
		Host         string
		Database     string
		TablePrefix  string
		Charset      string
		ParseTime    bool
		MaxIdleConns int
		MaxOpenConns int
	}

	Redis struct {
		Host     string
		Port     int
		Password string
		DB       int
		PoolSize int
	}
)

var (
	cfg            *ini.File
	ServerConfig   Server
	DatabaseConfig Database
	RedisConfig    Redis
)

func NewConfig(confPath string) {
	var err error
	cfg, err = ini.Load(confPath)
	if err != nil {
		log.Fatalf("config: fail to parse config: %v", err)
	}

	mapTo("server", &ServerConfig)
	mapTo("database", &DatabaseConfig)
	mapTo("redis", &RedisConfig)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("config: MapTo %s err: %v", section, err)
	}
}
