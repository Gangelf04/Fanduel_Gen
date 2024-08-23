package database

import (
	"Fanduel_Gen/internal/types"
	"log"
)

func (s *service) InsertPassingStats(player types.PassingStats) error {
	_, err := s.db.Exec(`INSERT INTO passing (season, week, seasontype, completions, attempts, interceptions, passtouchdowns, playerid)
	                    VALUES($1, $2, $3, $4, $5, $6, $7, $8);`, player.Season, player.Week, player.SeasonType, player.Completions,
		player.Attempts, player.Interceptions, player.PassTouchdowns, player.Player.SmartID)
	if err != nil {
		log.Print("An error occured while executing query: %v", err)
		return err
	}
	return nil
}
