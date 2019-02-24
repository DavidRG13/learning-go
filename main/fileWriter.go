package main

import (
	"encoding/json"
	"os"
)

type FileSystemPlayerStore struct {
	db     *json.Encoder
	league League
}

func NewFileSystemPlayerStore(database *os.File) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		db:     json.NewEncoder(&tape{database}),
		league: league,
	}
}

func (s *FileSystemPlayerStore) GetPlayerScore(playerName string) int {
	player := s.league.Find(playerName)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (s *FileSystemPlayerStore) GetLeague() League {
	return s.league
}

func (s *FileSystemPlayerStore) RecordWin(winner string) {
	player := s.league.Find(winner)

	if player != nil {
		player.Wins++
	} else {
		s.league = append(s.league, Player{winner, 1})
	}

	s.db.Encode(s.league)
}
