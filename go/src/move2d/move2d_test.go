package move2d


import (
	"testing"
)

func TestNewPoint(t *testing.T) {
	p := NewPoint(0,0)
	if p.x != 0 || p.y != 0 {
		t.Errorf("Error constructing new point")
	}
}

func TestPoint_MoveUp(t *testing.T) {
	p := NewPoint(0,0)
	p.MoveUp(7)
	if p.y != 7 {
		t.Errorf("Could not move up")
	}
	p.MoveUp(2)
	if p.y != 9 {
		t.Errorf("Could not move up")
	}
}

func TestPoint_MoveDown(t *testing.T) {
	p := NewPoint(0,0)
	p.MoveDown(3)
	if p.y != -3 {
		t.Errorf("Could not move up")
	}
	p.MoveDown(9)
	if p.y != -12 {
		t.Errorf("Could not move up")
	}
}

func TestPoint_MoveLeft(t *testing.T) {
	p := NewPoint(0,0)
	p.MoveLeft(1)
	if p.x != -1 {
		t.Errorf("Could not move up")
	}
	p.MoveLeft(9)
	if p.x != -10 {
		t.Errorf("Could not move up")
	}
}

func TestPoint_MoveRight(t *testing.T) {
	p := NewPoint(0,0)
	p.MoveRight(2)
	if p.x != 2 {
		t.Errorf("Could not move up")
	}
	p.MoveRight(3)
	if p.x != 5 {
		t.Errorf("Could not move up")
	}
}

