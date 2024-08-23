package types

type NextGenStatsResponseRoot struct {
	Season     int            `json:"season"`
	SeasonType string         `json:"seasonType"`
	Week       int            `json:"week"`
	Filter     string         `json:"filter"`
	Threshold  int            `json:"threshold"`
	Stats      []PassingStats `json:"stats"`
}
