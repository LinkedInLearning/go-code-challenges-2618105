package gol

import (
	"reflect"
	"testing"
)

func TestSingleCell(t *testing.T) {
	board := [][]bool{
		{false, false, false},
		{false, true, false},
		{false, false, false},
	}

	expected := [][]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}

	board = Evolve(board)

	if !reflect.DeepEqual(board, expected) {
		t.Errorf("Expected %v, got %v", expected, board)
	}

}

func TestBlinkerOscilates(t *testing.T) {
	board := [][]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}

	form1 := [][]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}

	form2 := [][]bool{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, false, false, false},
	}

	board = Evolve(board)

	if !reflect.DeepEqual(board, form2) {
		t.Errorf("Expected %v, got %v", form2, board)
	}

	board = Evolve(board)

	if !reflect.DeepEqual(board, form1) {
		t.Errorf("Expected %v, got %v", form1, board)
	}

}
