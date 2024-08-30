package types

type RotoWire []struct {
	ID                  string `json:"id"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	Player              string `json:"player"`
	URL                 string `json:"URL"`
	Injury              string `json:"injury"`
	Actions             string `json:"actions"`
	ActionsTrigger      string `json:"actions_trigger"`
	Lock                string `json:"lock"`
	Exclude             string `json:"exclude"`
	Like                string `json:"like"`
	Position            string `json:"position"`
	RealPosition        string `json:"real_position"`
	ActualPosition      string `json:"actual_position"`
	Team                string `json:"team"`
	Opponent            string `json:"opponent"`
	MoneyLine           string `json:"money_line"`
	OverUnder           string `json:"over_under"`
	PointSpread         string `json:"point_spread"`
	TeamPoints          string `json:"team_points"`
	Ownership           string `json:"ownership"`
	Multiplier          int    `json:"multiplier"`
	Salary              int    `json:"salary"`
	SalaryOriginal      int    `json:"salary_original"`
	SalaryCustom        int    `json:"salary_custom"`
	ProjPoints          string `json:"proj_points"`
	ProjRotowire        string `json:"proj_rotowire"`
	ProjOriginal        string `json:"proj_original"`
	ProjAvg             string `json:"proj_avg"`
	ProjCeiling         string `json:"proj_ceiling"`
	ProjFloor           string `json:"proj_floor"`
	ProjThirdPartyOne   int    `json:"proj_third_party_one"`
	ProjThirdPartyTwo   int    `json:"proj_third_party_two"`
	ProjThirdPartyThree int    `json:"proj_third_party_three"`
	ProjCustom          int    `json:"proj_custom"`
	Value               string `json:"value"`
	Dfsopp              string `json:"dfsopp"`
	Pa                  string `json:"pa"`
	Ra                  string `json:"ra"`
	Targets             string `json:"targets"`
	Touches             string `json:"touches"`
}
