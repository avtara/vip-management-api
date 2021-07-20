package dto

import "time"

//UpdateArrivedDTO is a model that used by client when POST from /login url
type UpdateArrivedDTO struct {
	Arrived bool `json:"arrived" form:"arrived" binding:"required"`
}

type InsertDTO struct {
	Name              string     `json:"name" form:"name" binding:"required"`
	Country_of_origin string     `json:"country_of_origin" form:"country_of_origin" binding:"required"`
	Eta               *time.Time `json:"eta" form:"eta" binding:"required,bookabledate" time_format:"2006-01-02 00:00:00" time_utc:"7"`
	Photo             string     `json:"photo" form:"photo" binding:"required"`
	Attributes        []string   `json:"attributes" form:"attributes" binding:"required"`
}
