package utils

import "time"

type TenhouLogConfig struct {
	LogUrl    string    `json:"log_url`
	StartDate time.Time `json:"start_date`
	EndDate   time.Time `json:"end_date`
}
