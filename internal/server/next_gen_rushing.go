package server

import (
	"Fanduel_Gen/internal/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) NextGenRushingHandler(c echo.Context) error {
	client := http.Client{}

	for i := 1; i <= 18; i++ {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://nextgenstats.nfl.com/api/statboard/rushing?season=2023&seasonType=REG&week=%d", i), nil)
		req.Header.Set("Referer", fmt.Sprintf("https://nextgenstats.nfl.com/stats/rushing/2023/REG/%d", i))

		if err != nil {
			log.Print(err)
		}

		res, err := client.Do(req)
		if err != nil {
			log.Print(err)
		}

		var data types.NextGenStatsResponseRoot

		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			log.Print(err)
		}

		for _, p := range data.Stats {
			err = s.db.InsertPlayer(p.Player)
			if err != nil {
				log.Print(err)
			}

			err = s.db.InsertPassingStats(p)
			if err != nil {
				log.Print(err)
			}
		}
	}

	return c.JSON(http.StatusOK, nil)
}

