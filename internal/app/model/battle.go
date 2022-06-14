package model

type Battle struct {
	ID            int `json:"id"`
	IDWinner      int `json:"id_winner"`
	IDLooser      int `json:"id_looser"`
	IDCompetition int `json:"id_competition"`
	WinnerScore   int `json:"winner_score"`
	LooserScore   int `json:"looser_score"`
}
