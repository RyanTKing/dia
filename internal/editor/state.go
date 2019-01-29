package editor

import (
	termbox "github.com/nsf/termbox-go"
)

const (
	// Normal mode is used for controlling navigation
	Normal Mode = iota + 1
	// Insert mode is used for modifying the text content
	Insert

	welcome = "dia, the tiny, extensible editor"
)

// NewState creates a new state
func NewState() (*State, error) {
	if err := termbox.Init(); err != nil {
		return nil, err
	}
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCursor(0, 0)

	w, h := termbox.Size()
	split := BufferSplit{
		Width:  w,
		Height: h,
		Active: true,
	}
	state := State{
		Width:       w,
		Height:      h,
		Mode:        Normal,
		TopSplit:    &split,
		ActiveSplit: &split,
	}

	state.redraw()

	return &state, nil
}

// StartLoop starts the main event loop
func (s *State) StartLoop() error {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == 'q' {
				return nil
			}

		case termbox.EventError:
			return ev.Err
		}

		s.redraw()
	}
}

// Destroy closes termbox and destroys the remaining state
func (s *State) Destroy() {
	termbox.Close()
}

func (s *State) redraw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	s.ActiveSplit.Draw()
	termbox.Flush()
}
