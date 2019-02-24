package main

import (
	"encoding/json"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	db     io.Writer
	league League
}

func NewFileSystemPlayerStore(database *os.File) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		db:     &tape{database},
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

	json.NewEncoder(s.db).Encode(s.league)
}
