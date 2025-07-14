package model

import "time"

type Player struct {
	Username  string
	Initials  string
	CreatedAt time.Time
}
