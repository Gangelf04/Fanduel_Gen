package types

type PassingStats struct {
	Aggressiveness                       float64 `json:"aggressiveness"`
	Attempts                             int     `json:"attempts"`
	AvgAirDistance                       float64 `json:"avgAirDistance"`
	AvgAirYardsDifferential              float64 `json:"avgAirYardsDifferential"`
	AvgAirYardsToSticks                  float64 `json:"avgAirYardsToSticks"`
	AvgCompletedAirYards                 float64 `json:"avgCompletedAirYards"`
	AvgIntendedAirYards                  float64 `json:"avgIntendedAirYards"`
	AvgTimeToThrow                       float64 `json:"avgTimeToThrow"`
	CompletionPercentage                 float64 `json:"completionPercentage"`
	CompletionPercentageAboveExpectation float64 `json:"completionPercentageAboveExpectation"`
	Completions                          int     `json:"completions"`
	ExpectedCompletionPercentage         float64 `json:"expectedCompletionPercentage"`
	GamesPlayed                          int     `json:"gamesPlayed"`
	Interceptions                        int     `json:"interceptions"`
	MaxAirDistance                       float64 `json:"maxAirDistance"`
	MaxCompletedAirDistance              float64 `json:"maxCompletedAirDistance"`
	PassTouchdowns                       int     `json:"passTouchdowns"`
	PassYards                            int     `json:"passYards"`
	PasserRating                         float64 `json:"passerRating"`
	PlayerName                           string  `json:"playerName"`
	Season                               int     `json:"season"`
	SeasonType                           string  `json:"seasonType"`
	Week                                 int     `json:"week"`
	Player                               Player  `json:"player"`
	Position                             string  `json:"position"`
	TeamID                               string  `json:"teamId"`
}
