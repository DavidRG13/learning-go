package main

import "testing"

func TestFindNextCellToFill(t *testing.T) {

	input := [][]int{
		{5, 1, 7, 6, 0, 0, 0, 3, 4},
		{2, 8, 9, 0, 0, 4, 0, 0, 0},
		{3, 4, 6, 2, 0, 5, 0, 9, 0},
		{6, 0, 2, 0, 0, 0, 0, 1, 0},
		{0, 3, 8, 0, 0, 6, 0, 4, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 9, 0, 0, 0, 0, 0, 7, 8},
		{7, 0, 3, 4, 0, 0, 5, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	t.Run("first 0 in row", func(t *testing.T) {
		wantedX, wantedY := 0, 4

		x, y := findNextCellToFill(input)

		if wantedX != x || wantedY != y {
			t.Errorf("got X %v wanted %v, got Y %v wanted %v", x, wantedX, y, wantedY)
		}
	})

	completedGrid := [][]int{
		{5, 1, 7, 6, 2, 2, 2, 3, 4},
		{2, 8, 9, 2, 2, 4, 2, 2, 2},
		{3, 4, 6, 2, 2, 5, 2, 9, 2},
		{6, 2, 2, 2, 2, 2, 2, 1, 2},
		{2, 3, 8, 2, 2, 6, 2, 4, 7},
		{2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 9, 2, 2, 2, 2, 2, 7, 8},
		{7, 2, 3, 4, 2, 2, 5, 6, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2},
	}

	t.Run("grid completed", func(t *testing.T) {
		wantedX, wantedY := -1, -1

		x, y := findNextCellToFill(completedGrid)

		if wantedX != x || wantedY != y {
			t.Errorf("got X %v wanted %v, got Y %v wanted %v", x, wantedX, y, wantedY)
		}
	})
}

func TestContainsInRow(t *testing.T) {

	input := [][]int{
		{5, 1, 7, 6, 0, 0, 0, 3, 4},
		{2, 8, 9, 0, 0, 4, 0, 0, 0},
		{3, 4, 6, 2, 0, 5, 0, 9, 0},
		{6, 0, 2, 0, 0, 0, 0, 1, 0},
		{0, 3, 8, 0, 0, 6, 0, 4, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 9, 0, 0, 0, 0, 0, 7, 8},
		{7, 0, 3, 4, 0, 0, 5, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	t.Run("row 0", func(t *testing.T) {
		isInRow := containsInRow(input[0], 5)

		if !isInRow {
			t.Errorf("Not found in row")
		}
	})

	t.Run("row 1", func(t *testing.T) {
		isInRow := containsInRow(input[1], 4)

		if !isInRow {
			t.Errorf("Not found in row")
		}
	})

	t.Run("row 7", func(t *testing.T) {
		isInRow := containsInRow(input[7], 6)

		if !isInRow {
			t.Errorf("Not found in row")
		}
	})

	t.Run("not in row 7", func(t *testing.T) {
		isInRow := containsInRow(input[7], 1)

		if isInRow {
			t.Errorf("Not expected in row")
		}
	})
}

func TestContainsInColumn(t *testing.T) {

	input := [][]int{
		{5, 1, 7, 6, 0, 0, 0, 3, 4},
		{2, 8, 9, 0, 0, 4, 0, 0, 0},
		{3, 4, 6, 2, 0, 5, 0, 9, 0},
		{6, 0, 2, 0, 0, 0, 0, 1, 0},
		{0, 3, 8, 0, 0, 6, 0, 4, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 9, 0, 0, 0, 0, 0, 7, 8},
		{7, 0, 3, 4, 0, 0, 5, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	t.Run("column 0", func(t *testing.T) {
		isInColumn := containsInColumn(input, 0, 5)

		if !isInColumn {
			t.Errorf("Not found in column")
		}
	})

	t.Run("column 1", func(t *testing.T) {
		isInColumn := containsInColumn(input, 1, 9)

		if !isInColumn {
			t.Errorf("Not found in column")
		}
	})

	t.Run("column 7", func(t *testing.T) {
		isInColumn := containsInColumn(input, 7, 6)

		if !isInColumn {
			t.Errorf("Not found in column")
		}
	})

	t.Run("not in column 7", func(t *testing.T) {
		isInColumn := containsInColumn(input, 7, 5)

		if isInColumn {
			t.Errorf("Not expected in column")
		}
	})
}

func TestContainsInSector(t *testing.T) {

	input := [][]int{
		{5, 1, 7, 6, 0, 0, 0, 3, 4},
		{2, 8, 9, 0, 0, 4, 0, 0, 0},
		{3, 4, 6, 2, 0, 5, 0, 9, 0},
		{6, 0, 2, 0, 0, 0, 0, 1, 0},
		{0, 3, 8, 0, 0, 6, 0, 4, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 9, 0, 0, 0, 0, 0, 7, 8},
		{7, 0, 3, 4, 0, 0, 5, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	t.Run("sector 0,0", func(t *testing.T) {
		isInSector := containsInSector(input, 0, 0, 6)

		if !isInSector {
			t.Errorf("Not found in sector")
		}
	})

	t.Run("sector 1,1", func(t *testing.T) {
		isInSector := containsInSector(input, 4, 5, 6)

		if !isInSector {
			t.Errorf("Not found in sector")
		}
	})

	t.Run("sector 2,2", func(t *testing.T) {
		isInSector := containsInSector(input, 8, 8, 6)

		if !isInSector {
			t.Errorf("Not found in sector")
		}
	})
}

func TestGetTopLeftCornerOfSector(t *testing.T) {

	t.Run("0,0 -> 0,0", func(t *testing.T) {
		wantedX, wantedY := 0, 0
		gotX, gotY := getTopLeftCellOfSector(0, 0)

		if wantedX != gotX || wantedY != gotY {
			t.Errorf("got X %v wanted %v, got Y %v wanted %v", gotX, wantedX, gotY, wantedY)
		}
	})

	t.Run("3,4 -> 3,3", func(t *testing.T) {
		wantedX, wantedY := 3, 3
		gotX, gotY := getTopLeftCellOfSector(3, 3)

		if wantedX != gotX || wantedY != gotY {
			t.Errorf("got X %v wanted %v, got Y %v wanted %v", gotX, wantedX, gotY, wantedY)
		}
	})

	t.Run("8,8 -> 6,6", func(t *testing.T) {
		wantedX, wantedY := 6, 6
		gotX, gotY := getTopLeftCellOfSector(8, 8)

		if wantedX != gotX || wantedY != gotY {
			t.Errorf("got X %v wanted %v, got Y %v wanted %v", gotX, wantedX, gotY, wantedY)
		}
	})
}
