package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileSystemPlayerStore struct {
	db     *json.Encoder
	league League
}

func NewFileSystemPlayerStore(database *os.File) (*FileSystemPlayerStore, error) {
	database.Seek(0, 0)

	info, err := database.Stat()

	if err != nil {
		return nil, fmt.Errorf("problem getting file info from file %s, %v", database.Name(), err)
	}

	if info.Size() == 0 {
		database.Write([]byte("[]"))
		database.Seek(0, 0)
	}

	league, err := NewLeague(database)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", database.Name(), err)
	}

	return &FileSystemPlayerStore{
		db:     json.NewEncoder(&tape{database}),
		league: league,
	}, nil
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
