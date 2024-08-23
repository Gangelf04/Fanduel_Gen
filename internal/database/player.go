package database

import (
	"Fanduel_Gen/internal/types"
	"log"
)

func (s *service) InsertPlayer(player types.Player) error {
	_, err := s.db.Exec(`INSERT INTO public.player (smartid, firstname, lastname, position, currentteamid)
                       VALUES ($1, $2, $3, $4, $5)
                       ON CONFLICT (smartid)
                       DO UPDATE SET 
                      firstname = EXCLUDED.firstname,
                      lastname = EXCLUDED.lastname,
                      position = EXCLUDED.position,
                      currentteamid = EXCLUDED.currentteamid;`, player.SmartID, player.FirstName, player.LastName, player.Position, player.CurrentTeamID)
	if err != nil {
		log.Print("An error occured while executing query: %v", err)
		return err
	}
	return nil
}
