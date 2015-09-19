package model

import "time"

type Deployment struct {
	Owner       string
	Repo        string
	Sha         string
	Environment string
	Author      string
	Gravatar    string
	Timestamp   time.Time
	Message     string
	Model
}
