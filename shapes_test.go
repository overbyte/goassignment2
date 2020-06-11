package main

import "testing"

func TestSquare(t *testing.T) {
	sq := square{
		sideLength: 10.0,
	}

	area := 100.0
	name := "square"

	if sq.getName() != name {
		t.Errorf("square name should be %v but was %v", name, sq.getName())
	}

	if sq.getArea() != area {
		t.Errorf("square area should be %v but was %v", area, sq.getArea())
	}
}

func TestTriange(t *testing.T) {
	tr := triangle{
		height: 10.0,
		base:   10.0,
	}

	area := 50.0
	name := "triangle"

	if tr.getName() != name {
		t.Errorf("triangle name should be %v but was %v", name, tr.getName())
	}

	if tr.getArea() != area {
		t.Errorf("triangle area should be %v but was %v", area, tr.getArea())
	}
}
