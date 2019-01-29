package editor

// Draw draws the buffer to the split
func (b *Buffer) Draw(x, y, xOff, yOff, width, height int) int {
	last := 0
	for i, row := range b.Rows {
		if i < yOff {
			continue
		}
		if i-yOff >= height {
			break
		}

		row.Draw(x, y+i-yOff, xOff, width)
		last = i - yOff
	}

	return last
}
