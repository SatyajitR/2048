package main

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

const (
	BoardSize = 4
)

type Game struct {
	board [BoardSize][BoardSize]*gtk.Label
}

func main() {
	// Initialize GTK.
	gtk.Init(nil)

	// Create a new window.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("2048")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new game.
	game := &Game{}

	// Create a new grid.
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	grid.SetRowSpacing(10)
	grid.SetColumnSpacing(10)

	// Create the game board.
	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			label, err := gtk.LabelNew("0")
			if err != nil {
				log.Fatal("Unable to create label:", err)
			}
			grid.Attach(label, x, y, 1, 1)
			game.board[y][x] = label
		}
	}

	// Add the grid to the window.
	win.Add(grid)

	// Show all widgets in the window.
	win.ShowAll()

	// Handle key press events.
	win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{Event: ev}
		switch keyEvent.KeyVal() {
		case gdk.KEY_Up:
			// Handle up arrow key press.
			for y := 0; y < BoardSize; y++ {
				for x := 0; x < BoardSize; x++ {
					game.board[y][x].SetText("U")
				}
			}
		case gdk.KEY_Down:
			// Handle down arrow key press.
			for y := 0; y < BoardSize; y++ {
				for x := 0; x < BoardSize; x++ {
					game.board[y][x].SetText("D")
				}
			}
		case gdk.KEY_Left:
			// Handle left arrow key press.
			for y := 0; y < BoardSize; y++ {
				for x := 0; x < BoardSize; x++ {
					game.board[y][x].SetText("L")
				}
			}
		case gdk.KEY_Right:
			// Handle right arrow key press.
			for y := 0; y < BoardSize; y++ {
				for x := 0; x < BoardSize; x++ {
					game.board[y][x].SetText("R")
				}
			}
		}
	})

	// Start the GTK main event loop.
	gtk.Main()
}
