package config

import (
	"github.com/go-ini/ini"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

type IniConf struct {
	confpath string
	cfg      ini.File
}

func (c *IniConf) CheckAndLoadConf(path string) {
	c.confpath = path
	cfg, err := ini.LooseLoad(path)
	c.cfg = *cfg

	if err != nil {
		log.Panic("Configuration file '"+path+"' not found ", err)
	}
}

func (c *IniConf) GetStringKey(section, key string) string {
	getsection := c.CheckSection(section)
	if !getsection.HasKey(key) {
		log.Panic("'" + key + "' key not found")
	}

	stringkey, err := c.cfg.Section(section).GetKey(key)
	if err != nil {
		log.Panic("Key '" + key + "' not found ")
	}

	return stringkey.String()
}

func (c *IniConf) GetBoolKey(section, key string) bool {
	c.CheckSection(section)
	return c.cfg.Section(section).Key(key).MustBool(false)
}

func (c *IniConf) CheckSection(section string) *ini.Section {
	getsection, err := c.cfg.GetSection(section)
	if err != nil {
		log.Panic("Section '" + section + "' of configuration file not found ")
	}
	return getsection
}

type Config struct {
	Debug      bool
	AuthSecret string
	Version    string `mapstructure:"version"`
	Port       string `mapstructure:"port"`
	Messages   Messages
}

type Messages struct {
	Responses
	Errors
}

type Responses struct {
	InlineContentTitle       string `mapstructure:"inline_content_title"`
	InlineContentDescription string `mapstructure:"inline_content_description"`
	NoStats                  string `mapstructure:"no_stats"`
}

type Errors struct {
	UnknownCommand string `mapstructure:"unknown_command"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	if err := parseEnv(&cfg); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &cfg, nil
}

func parseEnv(cfg *Config) error {

	if err := godotenv.Load(); err != nil {
		return err
	}

	viper.AutomaticEnv()

	cfg.Debug = viper.GetBool("DEBUG")
	cfg.AuthSecret = viper.GetString("AUTH_SECRET")

	return nil
}
