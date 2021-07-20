package entity

import "time"

//Athlete represents users table in database
type Athlete struct {
	ID                uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name              string    `gorm:"type:varchar(255)" json:"name"`
	Country_of_origin string    `gorm:"type:varchar(255)" json:"country_of_origin"`
	Eta               time.Time `json:"eta"`
	Arrived           bool      `gorm:"type:boolean" json:"arrived"`
	Photo             string    `gorm:"type:varchar(255)" json:"photo"`
	Attributes        string    `json:"attributes"`
}

type Attribute struct {
	ID   uint64 `gorm:"index" json:"-"`
	Name string `gorm:"type:varchar(255)" json:"name"`
}

//AthleteJSON represents users table in database
type AthleteJSON struct {
	ID                uint64    `json:"id"`
	Name              string    `json:"name"`
	Country_of_origin string    `json:"country_of_origin"`
	Eta               time.Time `json:"eta"`
	Arrived           bool      `json:"arrived"`
	Photo             string    `json:"photo"`
	Attributes        []string  `json:"attributes"`
}

//UpdateArrived represents users table in database
type UpdateArrived struct {
	Arrived bool `json:"arrived"`
}
