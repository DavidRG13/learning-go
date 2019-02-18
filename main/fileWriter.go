package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	db io.ReadWriteSeeker
}

func (s *FileSystemPlayerStore) GetPlayerScore(playerName string) int {

	player := s.GetLeague().Find(playerName)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (s *FileSystemPlayerStore) GetLeague() League {
	s.db.Seek(0, 0)
	league, _ := NewLeague(s.db)
	return league
}

func (s *FileSystemPlayerStore) RecordWin(winner string) {
	league := s.GetLeague()
	player := league.Find(winner)

	if player != nil {
		player.Wins++
	}

	s.db.Seek(0, 0)
	json.NewEncoder(s.db).Encode(league)
}
