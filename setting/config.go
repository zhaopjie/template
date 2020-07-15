package setting

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

//Config is config structure.
type Config struct {
	Mysql SectionMysql `yaml:"mysql" json:"mysql"`
	Redis SectionRedis `yaml:"redis" json:"redis"`
	Kafka SectionKafka `yaml:"kafka" json:"kafka"`
	Log   SectionLog   `yaml:"log" json:"log"`
	Core  SectionCore  `yaml:"core" json:"core"`
}

// SectionRedis is sub section of config.
type SectionRedis struct {
	Addr     string `yaml:"addr" json:"addr"`
	Password string `yaml:"password" json:"password"`
	DB       int    `yaml:"db" json:"db"`
}

// SectionMysql is sub section of config.
type SectionMysql struct {
	Account   string `yaml:"account" json:"account"`
	Password  string `yaml:"password" json:"password"`
	Addr      string `yaml:"addr" json:"addr"`
	Databases string `yaml:"databases" json:"databases"`
	Prefix    string `yaml:"prefix" json:"prefix"`
}

// SectionLog is sub section of config.
type SectionLog struct {
	Level   string `yaml:"level" json:"level"`
	OutPath string `yaml:"out_path" json:"out_path"`
}

// SectionKafka is sub section of config.
type SectionKafka struct {
	Addrs []string `yaml:"addrs" json:"addrs"`
}

type SectionCore struct {
	Port string `yaml:"port" json:"port"`
	//Mode string `yaml:"mode" json:"mode"`
}

// LoadConf 读取匹配的环境变量,只会读取第一个传入的config path
func (cfg *Config) LoadConf(confPath ...string) {
	vpr := viper.New()

	vpr.SetDefault("core.port", "8080")
	vpr.SetDefault("log.level", "debug")
	vpr.SetDefault("log.out_path", "./out.log")

	if len(confPath) != 0 {
		vpr.SetConfigFile(confPath[0])
		if err := vpr.ReadInConfig(); err != nil {
			log.Fatalf("read config err: %v", err)
		}
	} else {
		vpr.AutomaticEnv()
		vpr.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	}

	cfg.Mysql.Account = vpr.GetString("mysql.account")
	cfg.Mysql.Databases = vpr.GetString("mysql.databases")
	cfg.Mysql.Addr = vpr.GetString("mysql.addr")
	cfg.Mysql.Password = vpr.GetString("mysql.password")
	cfg.Mysql.Prefix = vpr.GetString("mysql.prefix")

	cfg.Redis.Addr = vpr.GetString("redis.addr")
	cfg.Redis.Password = vpr.GetString("redis.password")
	cfg.Redis.DB = vpr.GetInt("redis.db")

	cfg.Kafka.Addrs = vpr.GetStringSlice("kafka.addrs")

	cfg.Core.Port = vpr.GetString("core.port")

	cfg.Log.Level = vpr.GetString("log.level")
	cfg.Log.OutPath = vpr.GetString("log.out_path")

	//配置文件热更新
	//vpr.WatchConfig()
	//vpr.OnConfigChange(func(event fsnotify.Event) {
	//	fmt.Printf("Detect config change: %s \n", event.String())
	//})
	return
}

// Source mysql连接字符串
func (mysql SectionMysql) Source() (dbname string, Source string) {
	return "mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysql.Account, mysql.Password, mysql.Addr, mysql.Databases)
}
