package config

import (
	"log"

	esv7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/spf13/viper"
)

type Application struct {
	Env    *Env
	ESClient *esv7.Client
}

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	// DBURL                  string `mapstructure:"DB_URL"`
	ElasticURL             string `mapstructure:"ELASTICSEARCH_URL"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	ElasticUsername        string `mapstructure:"ELASTICSEARCH_USERNAME"`
	ElasticPassword        string `mapstructure:"ELASTICSEARCH_PASSWORD"`
	SecretKey              string `mapstructure:"SECRET_KEY"`
}

var cfg *Env
func New() *Env {
	cfg = &Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if cfg.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return cfg
}

func Get()*Env{
	if cfg==nil{
		cfg=New()
	}
	return cfg
}

func App() Application {
	var err error
	app := &Application{}
	app.Env = New()
	// app.ESClient, err = ConnectElasticsearch()
	if err != nil {
		log.Fatal(err)
	}
	return *app
}