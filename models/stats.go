package models

type StatsResponse struct {
	GamesPlayed   int     `json:"games_played"`
	Wins          int     `json:"wins"`
	Losses        int     `json:"losses"`
	Draws         int     `json:"draws"`
	WinPercentage float64 `json:"win_percentage"`
}
