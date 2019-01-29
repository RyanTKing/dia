package screen

const (
	welcomeMsg = "dia - a simple, extensible editor"
)

// func DrawScreen(buf *state.Buffer, x, y, w, h int) {
// 	var i, j int
// 	if buf == nil {
// 		j = h / 3
// 		i = (w - len(welcomeMsg)) / 2
// 		for k := 0; k < len(welcomeMsg); k++ {
// 			termbox.SetCell(i+k, j, rune(welcomeMsg[k]), termbox.ColorWhite, termbox.ColorBlack)
// 		}
// 	}

// 	j++
// 	for j < h {
// 		termbox.SetCell(0, j, '~', termbox.ColorWhite, termbox.ColorBlack)
// 		j++
// 	}
// }
