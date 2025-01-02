package config

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
)

// ServerConfig 服务配置
type ServerConfig struct {
	JwtKey      string `yaml:"JwtKey"` //JWT密钥
	N           int    `yaml:"N"`      //一次feed推流可提供的视频数量
	MysqlConfig `yaml:"MysqlConfig"`
	RedisConfig `yaml:"RedisConfig"`
}

// MysqlConfig mysql connection's config
type MysqlConfig struct {
	Name         string `yaml:"Name"`         //MySQL登录名
	Password     string `yaml:"Password"`     //MySQL登录密码
	Ip           string `yaml:"Ip"`           //MySQL连接ip
	Port         string `yaml:"Port"`         //MySQL连接端口
	DatabaseName string `yaml:"DatabaseName"` //MySQL数据库名
}

// RedisConfig config
type RedisConfig struct {
	Network     string `yaml:"Network"`     //redis网络连接类型 tcp/udp
	Address     string `yaml:"Address"`     //redis地址
	Num         int    `yaml:"Num"`         //redis数据库（0~11）
	RdbPassword string `yaml:"RdbPassword"` //redis密码
}

func GetConfigMessageFromYaml(yamlFile *string, ServerConfig *ServerConfig, c *gin.Context) {
	content, err := os.ReadFile(*yamlFile)
	if err != nil && err != io.EOF {
		log.Fatal("yaml file read fail", err)
	}

	if err := yaml.Unmarshal(content, ServerConfig); err != nil {
		log.Fatal("yaml file parse fail\n", err)
	}
	c.Set("ServerConfig", ServerConfig)
}
