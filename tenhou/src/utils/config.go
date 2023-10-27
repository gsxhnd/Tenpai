package utils

import (
	"time"

	"github.com/google/wire"
)

type tenhouLogConfig struct {
	LogUrl    string    `json:"log_url"`
	StartDate time.Time `json:"start_date`
	EndDate   time.Time `json:"end_date`
}

func NewTenhouLogConfig() *tenhouLogConfig {
	return &tenhouLogConfig{
		LogUrl:    "",
		StartDate: time.Now(),
		EndDate:   time.Now(),
	}
}

var ConfigSet = wire.NewSet(NewTenhouLogConfig())
