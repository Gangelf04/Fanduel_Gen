package types

type NextGenRushingStats struct {
	Season     int          `json:"season"`
	SeasonType string       `json:"seasonType"`
	Week       int          `json:"week"`
	Filter     string       `json:"filter"`
	Threshold  int          `json:"threshold"`
	Stats      RushingStats `json:"stats"`
}

type RushingStats struct {
	AvgTimeToLos                     float64 `json:"avgTimeToLos"`
	ExpectedRushYards                float64 `json:"expectedRushYards"`
	RushAttempts                     int     `json:"rushAttempts"`
	RushPctOverExpected              float64 `json:"rushPctOverExpected"`
	RushTouchdowns                   int     `json:"rushTouchdowns"`
	RushYards                        int     `json:"rushYards"`
	RushYardsOverExpected            float64 `json:"rushYardsOverExpected"`
	RushYardsOverExpectedPerAtt      float64 `json:"rushYardsOverExpectedPerAtt"`
	Player                           Player  `json:"player"`
	TeamID                           string  `json:"teamId"`
	Efficiency                       float64 `json:"efficiency"`
	PercentAttemptsGteEightDefenders int     `json:"percentAttemptsGteEightDefenders"`
	AvgRushYards                     float64 `json:"avgRushYards"`
}
