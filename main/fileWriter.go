package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	db io.Reader
}

func (s *FileSystemPlayerStore) GetLeague() []Player {
	league, _ := NewLeague(s.db)
	return league
}
