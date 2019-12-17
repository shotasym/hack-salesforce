package main

import (
	"github.com/sclevine/agouti"
)

type Driver struct {
	*agouti.WebDriver
}

func NewChromeDriver(c agouti.Option) *Driver {
	return &Driver{agouti.ChromeDriver(c)}
}
