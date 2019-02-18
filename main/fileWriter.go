package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	db io.ReadSeeker
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
