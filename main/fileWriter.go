package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	db     *json.Encoder
	league League
}

func NewFileSystemPlayerStore(database *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(database)

	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
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

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

func (s *FileSystemPlayerStore) GetPlayerScore(playerName string) int {
	player := s.league.Find(playerName)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (s *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(s.league, func(i, j int) bool {
		return s.league[i].Wins > s.league[j].Wins
	})
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
