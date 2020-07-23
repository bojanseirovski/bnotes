package Config

//Config/Database.go

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {

	f, err := os.Open("config.yml")
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("Status:", err)
	}

	var dbPort, errP = strconv.Atoi(cfg.Server.Port)
	if errP != nil {
		fmt.Println("Status:", errP)
	}
	dbConfig := DBConfig{
		Host:     cfg.Server.Host,
		Port:     dbPort,
		User:     cfg.Database.Username,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.Name,
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
