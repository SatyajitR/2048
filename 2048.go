package main

import (
	"math/rand"

	"github.com/nsf/termbox-go"
)

const (
	BoardSize = 4
)

type Game struct {
	board [BoardSize][BoardSize]int
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	game := &Game{}
	game.Start()
}

func (g *Game) Start() {
	g.board = [BoardSize][BoardSize]int{}
	g.spawn()
	g.spawn()
	g.draw()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				g.up()
			case termbox.KeyArrowDown:
				g.down()
			case termbox.KeyArrowLeft:
				g.left()
			case termbox.KeyArrowRight:
				g.right()
			case termbox.KeyEsc:
				break mainloop
			}
		}
		g.draw()
	}
}

func (g *Game) draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			termbox.SetCell(x*2, y, rune('0'+g.board[y][x]), termbox.ColorWhite, termbox.ColorBlack)
		}
	}
	termbox.Flush()
}

func (g *Game) spawn() {
	for {
		x := rand.Intn(BoardSize)
		y := rand.Intn(BoardSize)
		if g.board[y][x] == 0 {
			g.board[y][x] = 2
			break
		}
	}
}

func (g *Game) up() {
	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			if g.board[y][x] != 0 {
				for y1 := y + 1; y1 < BoardSize; y1++ {
					if g.board[y1][x] == 0 {
						continue
					}
					if g.board[y1][x] != g.board[y][x] {
						break
					}
					g.board[y][x] *= 2
					g.board[y1][x] = 0
					break
				}
			}
		}
	}
	g.spawn()
}

func (g *Game) down() {
	for x := 0; x < BoardSize; x++ {
		for y := BoardSize - 1; y >= 0; y-- {
			if g.board[y][x] != 0 {
				for y1 := y - 1; y1 >= 0; y1-- {
					if g.board[y1][x] == 0 {
						continue
					}
					if g.board[y1][x] != g.board[y][x] {
						break
					}
					g.board[y][x] *= 2
					g.board[y1][x] = 0
					break
				}
			}
		}
	}
	g.spawn()
}

func (g *Game) left() {
	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			if g.board[y][x] != 0 {
				for x1 := x + 1; x1 < BoardSize; x1++ {
					if g.board[y][x1] == 0 {
						continue
					}
					if g.board[y][x1] != g.board[y][x] {
						break
					}
					g.board[y][x] *= 2
					g.board[y][x1] = 0
					break
				}
			}
		}
	}
	g.spawn()
}

func (g *Game) right() {
	for y := 0; y < BoardSize; y++ {
		for x := BoardSize - 1; x >= 0; x-- {
			if g.board[y][x] != 0 {
				for x1 := x - 1; x1 >= 0; x1-- {
					if g.board[y][x1] == 0 {
						continue
					}
					if g.board[y][x1] != g.board[y][x] {
						break
					}
					g.board[y][x] *= 2
					g.board[y][x1] = 0
					break
				}
			}
		}
	}
	g.spawn()
}
