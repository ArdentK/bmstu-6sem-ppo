package model

import "time"

type Competition struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Date        time.Time `json:"dt"`
	AgeCategory string    `json:"age_category"`
	WeaponType  string    `json:"weapon_type"`
	IsTeam      int       `json:"is_team"`
	Status      string    `json:"status"`
	Sex         string    `json:"sex"`
	Type        string    `json:"type"`
}
