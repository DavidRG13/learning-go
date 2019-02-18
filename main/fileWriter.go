package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	db io.ReadWriteSeeker
}

func (s *FileSystemPlayerStore) GetPlayerScore(playerName string) int {

	var wins int

	for _, player := range s.GetLeague() {
		if player.Name == playerName {
			wins = player.Wins
			break
		}
	}

	return wins
}

func (s *FileSystemPlayerStore) GetLeague() []Player {
	s.db.Seek(0, 0)
	league, _ := NewLeague(s.db)
	return league
}

func (s *FileSystemPlayerStore) RecordWin(winner string) {
	league := s.GetLeague()

	for i, player := range league {
		if player.Name == winner {
			league[i].Wins++
		}
	}

	s.db.Seek(0, 0)
	json.NewEncoder(s.db).Encode(league)
}
