package Model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Operation struct {
	gorm.Model
	Name    string
	Mobile  int
	Address string
	Arrived string
	File string

	Date time.Duration
}
