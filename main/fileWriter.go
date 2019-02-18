package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	db io.ReadSeeker
}

func (s *FileSystemPlayerStore) GetLeague() []Player {
	s.db.Seek(0, 0)
	league, _ := NewLeague(s.db)
	return league
}
