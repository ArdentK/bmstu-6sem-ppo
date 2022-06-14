package model

import (
	"testing"
	"time"
)

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}

func TestCompetition(t *testing.T) *Competition {
	return &Competition{
		Name:        "example",
		Date:        time.Now(),
		AgeCategory: "juniors",
		WeaponType:  "epee",
		IsTeam:      0,
		Status:      "Russian",
		Sex:         "female",
		Type:        "Tournament of Russia",
	}
}

func TestAthlet(t *testing.T) *Athlet {
	return &Athlet{
		Name:       "Ololo Kekovich",
		Birthday:   time.Now(),
		Role:       "athlet",
		WeaponType: "epee",
		Sex:        "female",
		RFSubject:  "Moscow",
		Rank:       "Master of Sport",
		SportOrg:   "DInamo",
	}
}

func TestBattle(t *testing.T) *Battle {
	return &Battle{
		IDWinner:      1,
		IDLooser:      2,
		IDCompetition: 1,
		WinnerScore:   15,
		LooserScore:   11,
	}
}
