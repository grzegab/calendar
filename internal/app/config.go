package app

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type HttpConfig struct {
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

type JWTConfig struct {
	Secret string
}

type Config struct {
	ConfigFile *string
	Debug      bool
	Origins    []string
	DB         DatabaseConfig
	HTTP       HttpConfig
	JWT        JWTConfig
	Modules    []string `mapstructure:"modules"`
	Addr       string   `mapstructure:"addr"`
	PprofAddr  string   `mapstructure:"pprof_addr"`
}

var AppConfig Config

func LoadConfig() error {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	v.AddConfigPath("../../")

	// Load default config file
	v.SetConfigName("config.default")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := errors.AsType[viper.ConfigFileNotFoundError](err); !ok {
			return fmt.Errorf("error reading base config file: %w", err)
		}
	}

	// Load user config file
	if AppConfig.ConfigFile != nil {
		fmt.Println("Loading user config file: ", *AppConfig.ConfigFile)
		v.SetConfigName(*AppConfig.ConfigFile)
		if err := v.MergeInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return fmt.Errorf("error reading env config file: %w", err)
			}
		}
	}

	// Set listening port
	AppConfig.Addr = v.GetString("addr")

	// Set profiler port
	AppConfig.PprofAddr = v.GetString("pprof_addr")

	// Set database values
	AppConfig.DB.Host = v.GetString("db.host")
	AppConfig.DB.Port = v.GetInt("db.port")
	AppConfig.DB.User = v.GetString("db.user")
	AppConfig.DB.Password = v.GetString("db.password")
	AppConfig.DB.DBName = v.GetString("db.name")
	AppConfig.DB.SSLMode = v.GetString("db.sslMode")

	// HTTP settings
	AppConfig.HTTP.ReadTimeout = v.GetInt("http.readTimeout")
	AppConfig.HTTP.WriteTimeout = v.GetInt("http.writeTimeout")
	AppConfig.HTTP.IdleTimeout = v.GetInt("http.idleTimeout")

	// JWT settings
	AppConfig.JWT.Secret = v.GetString("jwt_secret")

	// Enable modules
	modules := v.GetStringSlice("modules")
	for _, m := range modules {
		AppConfig.Modules = append(AppConfig.Modules, m)
	}

	// Set CORS origins
	origins := v.GetStringSlice("origins")
	for _, o := range origins {
		AppConfig.Origins = append(AppConfig.Origins, o)
	}

	return nil
}
