package model

import "time"

type Athlet struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Birthday   time.Time `json:"birthday"`
	Role       string    `json:"role_part"`
	WeaponType string    `json:"weapon_type"`
	Sex        string    `json:"sex"`
	RFSubject  string    `json:"rf_subject"`
	Rank       string    `json:"sport_rank"`
	SportOrg   string    `json:"sport_org"`
}
