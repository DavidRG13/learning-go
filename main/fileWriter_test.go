package main

import (
	"strings"
	"testing"
)

func TestSystemStore(t *testing.T) {
	t.Run("/league from reader", func(t *testing.T) {
		db := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},		
			{"Name": "Chris", "Wins": 33}]`)

		store := FileSystemPlayerStore{db}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)
	})

	t.Run("/league from reader twice", func(t *testing.T) {
		db := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},		
			{"Name": "Chris", "Wins": 33}]`)

		store := FileSystemPlayerStore{db}

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
		database := strings.NewReader(`[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Chris")

		assertScoreEquals(t, got, 33)
	})
}

func assertScoreEquals(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
