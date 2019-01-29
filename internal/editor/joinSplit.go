package editor

const (
	// HorizontalSplit represents a split made in the horizontal direction
	HorizontalSplit SplitDirection = iota + 1
	// VerticalSplit represents a split made in the vertical direction
	VerticalSplit
)

// Draw ...
func (s *JoinSplit) Draw() {
	s.Child1.Draw()
	s.Child2.Draw()
}

// Split ...
func (s *JoinSplit) Split(dir SplitDirection) Split {
	return s
}
