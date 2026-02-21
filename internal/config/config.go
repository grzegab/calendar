package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ConfigFile string
	Debug      bool
	Origins    []string
	Services   map[string]string `mapstructure:"services"`
	Addr       string            `mapstructure:"addr"`
	PprofAddr  string            `mapstructure:"pprof_addr"`
}

var AppConfig Config

func LoadConfig() error {
	v := viper.New()

	// Initial env check from OS directly
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "prod"
	}

	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	v.AddConfigPath("../../")

	// Load base config
	v.SetConfigName("config")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("error reading base config file: %w", err)
		}
	}

	// Load environment specific config
	v.SetConfigName("config." + env)
	if err := v.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("error reading env config file: %w", err)
		}
	}

	v.SetDefault("addr", ":80")
	v.SetDefault("pprof_addr", ":6060")

	// Workaround for dots in keys during Unmarshal
	// Instead of Unmarshal, we can manually get the values or use a simpler structure

	AppConfig.Addr = v.GetString("addr")
	AppConfig.PprofAddr = v.GetString("pprof_addr")

	// Manually load services to avoid mapstructure issues with dots in keys
	AppConfig.Services = make(map[string]string)
	services := v.GetStringMap("services")
	for k, v := range services {
		if s, ok := v.(string); ok {
			AppConfig.Services[k] = s
		}
	}

	origins := v.GetStringSlice("origins")
	for _, o := range origins {
		AppConfig.Origins = append(AppConfig.Origins, o)
	}

	// Override from env vars if present (simple fields)
	if addr := os.Getenv("ADDR"); addr != "" {
		AppConfig.Addr = addr
	}
	if pprofAddr := os.Getenv("PPROF_ADDR"); pprofAddr != "" {
		AppConfig.PprofAddr = pprofAddr
	}

	return nil
}
