package server

import (
	"Fanduel_Gen/internal/types"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

const URL = "https://www.rotowire.com/daily/tables/value-report-nfl.php"

func (s *Server) RotoWireHandler(c echo.Context) error {
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, URL, nil)

	q := req.URL.Query()
	q.Add("siteID", "2")
	q.Add("slateID", "6778")
	q.Add("projSource", "RotoWire")
	q.Add("oshipSource", "RotoWire")
	req.URL.RawQuery = q.Encode()

	if err != nil {
		log.Print(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Print(err)
	}

	var data types.RotoWire

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		log.Print(err)
	}

	return c.JSON(http.StatusOK, data)
}
