package move2d

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

// Point represents a physical point on the grid
type Point struct {
	x int
	y int
}

// NewPoint is a constructor to create a new point
func NewPoint(x, y int) *Point {
	return &Point{x: x, y: y}
}

func init() {
	fmt.Println("Enter Steps: ")
}

func main() {
	coor := NewPoint(0, 0)
	count := 0

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		printResults(*coor, count)
		os.Exit(1)
	}()

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		line := s.Text()
		dir, n := parseInstructions(line)
		switch dir {
		case "left", "l":
			coor.MoveLeft(n)
		case "right", "r":
			coor.MoveRight(n)
		case "down", "d":
			coor.MoveDown(n)
		case "up", "u":
			coor.MoveUp(n)
		default:
			fmt.Printf("Could not interpret directions: [%s]\n",line)
		}
		count++

		if count % 5 == 0 {
			fmt.Printf("Moved %d times | Position:%s", count, coor.getPosition())
		}
	}
}


func parseInstructions(s string) (string, int) {
	inst := strings.Split(s, " ")
	if len(inst) < 2 {
		log.Printf("Invalid instructions: [%s]\n", s)
		return "", 0
	}
	n, err  := strconv.Atoi(inst[1])
	if err != nil {
		log.Printf("Could not move: [%s] %v\n", s, err)
		return "", 0
	}
	return strings.ToLower(inst[0]), n
}

func printResults(p Point, count int) {
	fmt.Printf("\nFinal Position: %s\nNumber of Moves: %d\n", p.getPosition(), count)
}

func (p Point) getPosition() string {
	return fmt.Sprintf("[%d,%d]", p.x, p.y)
}

// MoveLeft moves Point n positions left on the x axis
func (p *Point) MoveLeft(n int) {
	p.x -= n
}

// MoveRight moves Point n positions right on the x axis
func (p *Point) MoveRight(n int) {
	p.x += n
}

// MoveDown moves Point n positions down on the y axis
func (p *Point) MoveDown(n int) {
	p.y -= n
}

// MoveUp moves Point n positions up on the y axis
func (p *Point) MoveUp(n int) {
	p.y += n
}
