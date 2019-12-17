package main

import "gopkg.in/go-ini/ini.v1"

type Config struct {
	User     string
	Password string
}

const (
	salesforceSection = "salesforce"
	userKey           = "user"
	passwordKey       = "password"
)

func NewConfig(configPath string) (*Config, error) {
	c, err := ini.Load(configPath)
	if err != nil {
		return nil, err
	}
	return &Config{
		User:     c.Section(salesforceSection).Key(userKey).String(),
		Password: c.Section(salesforceSection).Key(passwordKey).String(),
	}, nil
}
