package main

import (
	"github.com/sclevine/agouti"

	"flag"
	"io/ioutil"
	"time"
)

type DailyWorks struct {
	DailyWorks []DailyWork
}

type DailyWork struct {
	Date       string    `json:"date"`
	Start      string    `json:"start"`
	End        string    `json:"end"`
	Projects   []Project `json:"projects"`
	TypeChange TypeChange
}

type Project struct {
	Name     string `json:"name"`
	Duration string `json:"duration"`
}

type TypeChange struct {
	Date time.Time
}

func (d *DailyWork) ChangeType() error {

	// 日付をtime型に
	parseDate, err := time.Parse(WorkTdDateFormat, d.Date)
	if err != nil {
		return err
	}
	d.TypeChange.Date = parseDate
	return nil
}

var (
	configPath string
	jsonFile   string
)

func main() {
	flag.StringVar(&configPath, "config_path", "", "ini config path")
	flag.StringVar(&jsonFile, "jsonfile", "", "json file for dailywork")
	flag.Parse()

	logger := NewLogger()
	logger.Info("start.")

	// Setting
	if err := flagValidate(configPath, jsonFile); err != nil {
		logger.Errorf("Invalid Flag:%v", err)
		return
	}
	config, err := NewConfig(configPath)
	if err != nil {
		logger.Errorf("NewConfig Error:%v", err)
		return
	}
	if err = config.validate(); err != nil {
		logger.Errorf("config validate Error:%v", err)
		return
	}
	dailyWorks := &DailyWorks{}
	bytes, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		logger.Errorf("Failed to ReadFile:%v", err)
		return
	}
	if err := dailyWorks.ParseJson(bytes); err != nil {
		logger.Errorf("Failed to ParseJson:%v", err)
		return
	}
	if err := dailyWorks.validate(); err != nil {
		logger.Errorf("dailyWorks Validate Error:%v", err)
		return
	}

	// Driver Start
	driver := NewChromeDriver(agouti.Desired(agouti.Capabilities{}))
	if err := driver.Start(); err != nil {
		logger.Errorf("Failed to start:%v", err)
		return
	}
	defer driver.Stop()
	sf, err := driver.NewSalesForce(config.User, config.Password)
	if err != nil {
		logger.Errorf("Failed to create instance:%v", err)
		return
	}
	err = sf.Login()
	if err != nil {
		logger.Errorf("Failed to login:%v", err)
		return
	}

	// Setting Daily Works
	for _, dailyWork := range dailyWorks.DailyWorks {
		logger.Infof("start to register %s work.", dailyWork.Date)
		if err := dailyWork.ChangeType(); err != nil {
			logger.Errorf("Failed to ChangeType:%v", err)
			return
		}
		if err := sf.RegisterWork(dailyWork); err != nil {
			logger.Errorf("Failed to RegisterWork:%v", err)
			return
		}
		logger.Infof("finish to register %s work.", dailyWork.Date)
	}

	logger.Info("finish.")
}
