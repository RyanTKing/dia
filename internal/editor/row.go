package editor

import termbox "github.com/nsf/termbox-go"

// Update updates the render of the row
func (r *Row) Update(ts int) {
	nTabs := 0
	for _, c := range r.Raw {
		if c == '\t' {
			nTabs++
		}
	}

	r.Render = make([]rune, len(r.Raw)+nTabs*(ts-1))
	i := 0
	for _, c := range r.Raw {
		if c == '\t' {
			r.Render[i] = ' '
			i++
			for i%ts != 0 {
				r.Render[i] = ' '
				i++
			}
		} else {
			r.Render[i] = c
			i++
		}
	}
}

// Draw draws the row
func (r *Row) Draw(x, y, xOff, width int) {
	for i := 0; i < len(r.Render) && i+xOff < width; i++ {
		termbox.SetCell(x+i, y, r.Render[i+xOff], termbox.ColorWhite, termbox.ColorBlack)
	}
}
