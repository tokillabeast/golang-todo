package models

import "time"

type Base struct {
	Created  time.Time `gorethink:"created,omitempty" json:"created,omitempty"`
	Modified time.Time `gorethink:"modified,omitempty" json:"modified,omitempty"`
}
