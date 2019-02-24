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

		store, err := NewFileSystemPlayerStore(db)

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)
		assertNoError(t, err)
	})

	t.Run("/league from reader twice", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},		
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store, err := NewFileSystemPlayerStore(db)

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
		assertNoError(t, err)
	})

	t.Run("get player score", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store, err := NewFileSystemPlayerStore(db)

		got := store.GetPlayerScore("Chris")

		assertScoreEquals(t, got, 33)
		assertNoError(t, err)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store, err := NewFileSystemPlayerStore(db)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
		assertNoError(t, err)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
		assertNoError(t, err)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)
	})
}

func assertScoreEquals(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didnt expect an error but got one, %v", err)
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
