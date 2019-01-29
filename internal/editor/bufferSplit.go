package editor

import termbox "github.com/nsf/termbox-go"

// Draw ...
func (s *BufferSplit) Draw() {
	i := 1
	if s.Buffer == nil {
		padding := (s.Width - len(welcome)) / 2
		if padding <= 0 {
			padding = 1
		}
		row := &Row{Render: []rune(welcome)}
		row.Draw(padding, s.Height/3, 0, s.Width)
	} else {
		i = s.Buffer.Draw(s.X, s.Y, s.XOff, s.YOff, s.Width, s.Height)
	}
	for i < s.Height {
		termbox.SetCell(0, i, '~', termbox.ColorWhite, termbox.ColorBlack)
		i++
	}
}

func split(m int) (int, int) {
	m1 := m/2 + (m % 2)
	m2 := m / 2
	return m1, m2
}

// Split ...
func (s *BufferSplit) Split(dir SplitDirection) Split {
	js := JoinSplit{}
	s2 := BufferSplit{
		XOff:   s.XOff,
		YOff:   s.YOff,
		Cx:     s.Cx,
		Cy:     s.Cy,
		Buffer: s.Buffer,
		Parent: &js,
	}

	js.Parent = s.Parent
	s.Parent = &js
	if dir == HorizontalSplit {
		s.Height, s2.Height = split(s.Height)
		s2.Width = s.Width
		s2.X = s.X
		s2.Y = s.Y + s.Height
	} else {
		s.Width, s2.Width = split(s.Width)
		s2.Height = s.Height
		s2.X = s.X + s.Width
		s2.Y = s.Y
	}

	js.Child1 = s
	js.Child2 = &s2

	return &js
}
