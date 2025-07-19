package config

import (
	"cmp"
	"fmt"
	"os"
)

type Config struct {
	JwtSecret string
	Server    ServerConfig
	MySQL     MySQLConfig
}

type ServerConfig struct {
	Port        string
	AllowOrigin string
}

type MySQLConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func GetHttpConfig() *Config {
	cfg := &Config{
		JwtSecret: cmp.Or(os.Getenv("JWT_SECRET"), "your-secure-test-secret-key-for-testing-12345"),
		Server: ServerConfig{
			Port:        cmp.Or(os.Getenv("SERVER_PORT"), "8080"),
			AllowOrigin: cmp.Or(os.Getenv("CORS_ALLOW_ORIGIN"), "http://localhost:3000"),
		},
		MySQL: MySQLConfig{
			Host:     cmp.Or(os.Getenv("MYSQL_HOST"), "db"),
			Port:     cmp.Or(os.Getenv("MYSQL_PORT"), "3306"),
			User:     cmp.Or(os.Getenv("MYSQL_USER"), "root"),
			Password: cmp.Or(os.Getenv("MYSQL_PASSWORD"), "password"),
			Database: cmp.Or(os.Getenv("MYSQL_DATABASE"), "ito_denwa"),
		},
	}
	return cfg
}

func (c *Config) GetMySQLDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.MySQL.User,
		c.MySQL.Password,
		c.MySQL.Host,
		c.MySQL.Port,
		c.MySQL.Database,
	)
}
