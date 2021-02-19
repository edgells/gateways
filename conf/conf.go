package conf

import (
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int

	RedisHost     string
	RedisUser     string
	RedisPassword string

	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/settings.conf")
	if err != nil {
		log.Fatalf("Fail to parse conf/settings.conf, :%v", err)
	}

	LoadBase()
	LoadApp()
	LoadServer()
	LoadDb()
}

func LoadBase() {
	// 载入app配置
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalln(err)
	}

	JwtSecret = sec.Key("JWT").Value()
	if JwtSecret == "" {
		log.Fatalln("jwt is not none")
	}
}

func LoadServer() {
	// 载入服务配置
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalln(err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8888)
	ReadTimeout = time.Duration(sec.Key("ReadTimeout").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WriteTimeout").MustInt(60)) * time.Second

}

func LoadDb() {
	sec, err := Cfg.GetSection("redis")
	if err != nil {
		log.Fatalln(err)
	}

	// load redis config
	RedisHost = sec.Key("HOST").MustString("localhost:6379")
	RedisUser = sec.Key("USERNAME").MustString("")
	RedisPassword = sec.Key("PASSWORD").MustString("")
}
