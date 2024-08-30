package types

type Player struct {
	Season           int    `json:"season"`
	SeasonType       string `json:"seasonType"`
	Week             int    `json:"week"`
	JerseyNumber     int    `json:"jerseyNumber"`
	LastName         string `json:"lastName"`
	FootballName     string `json:"footballName"`
	FirstName        string `json:"firstName"`
	Position         string `json:"position"`
	GsisItID         int    `json:"gsisItId"`
	GsisID           string `json:"gsisId"`
	EsbID            string `json:"esbId"`
	DisplayName      string `json:"displayName"`
	ShortName        string `json:"shortName"`
	UniformNumber    string `json:"uniformNumber"`
	Status           string `json:"status"`
	CurrentTeamID    string `json:"currentTeamId"`
	SmartID          string `json:"smartId"`
	Headshot         string `json:"headshot"`
	PositionGroup    string `json:"positionGroup"`
	NgsPosition      string `json:"ngsPosition"`
	NgsPositionGroup string `json:"ngsPositionGroup"`
}
