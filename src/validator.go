package main

import (
	"errors"
	"time"
)

func flagValidate(configPath, jsonfile string) error {
	if configPath == "" {
		return errors.New("empty configPath.")
	}
	if jsonfile == "" {
		return errors.New("empty jsonfile.")
	}
	return nil
}

func (c *Config) validate() error {
	if c.User == "" {
		return errors.New("empty UserName from ini file.")
	}
	if c.Password == "" {
		return errors.New("empty Password from ini file.")
	}
	return nil
}

func (d *DailyWorks) validate() error {
	for _, dailyWork := range d.DailyWorks {
		if err := dailyWork.validate(); err != nil {
			return err
		}
	}
	return nil
}

const (
	DailyWorkDateFormat = "2006-01-02"
	DailyWorkTimeFormat = "15:04"
)

func (d *DailyWork) validate() error {

	_, err := time.Parse(DailyWorkDateFormat, d.Date)
	if err != nil {
		return err
	}

	for _, t := range []string{d.Start, d.End} {
		_, err := time.Parse(DailyWorkTimeFormat, t)
		if err != nil {
			return err
		}
	}
	for _, t := range d.Projects {
		_, err := time.Parse(DailyWorkTimeFormat, t.Duration)
		if err != nil {
			return err
		}
	}
	return nil
}
