package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	db io.Reader
}

func (s *FileSystemPlayerStore) GetLeague() []Player {
	var league []Player
	json.NewDecoder(s.db).Decode(&league)
	return league
}
