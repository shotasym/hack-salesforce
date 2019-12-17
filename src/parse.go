package main

import (
	"encoding/json"
)

func (d *DailyWorks) ParseJson(b []byte) error {
	if err := json.Unmarshal(b, &d.DailyWorks); err != nil {
		return err
	}
	return nil
}
