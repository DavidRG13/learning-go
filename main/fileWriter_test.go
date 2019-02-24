package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSystemStore(t *testing.T) {
	t.Run("/league from reader", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},		
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store := NewFileSystemPlayerStore(db)

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)
	})

	t.Run("/league from reader twice", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},		
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store := NewFileSystemPlayerStore(db)

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store := NewFileSystemPlayerStore(db)

		got := store.GetPlayerScore("Chris")

		assertScoreEquals(t, got, 33)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store := NewFileSystemPlayerStore(db)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
