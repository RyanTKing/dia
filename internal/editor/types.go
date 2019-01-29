package editor

// Mode represents the mode of the editor
type Mode int

// State represents the global editor state
type State struct {
	Width       int
	Height      int
	Mode        Mode
	TopSplit    Split
	ActiveSplit *BufferSplit
}

// SplitDirection represents the direction of a split
type SplitDirection int

// Split represents a drawable split
type Split interface {
	Draw()
	Split(dir SplitDirection) Split
}

// JoinSplit represents a joining of two splits
type JoinSplit struct {
	Child1 Split
	Child2 Split

	Parent Split
}

// BufferSplit represents a split that holds a single buffer
type BufferSplit struct {
	X      int
	Y      int
	Width  int
	Height int
	XOff   int
	YOff   int
	Cx     int
	Cy     int
	Active bool

	Buffer *Buffer

	Parent Split
}

// Buffer holds a file buffer
type Buffer struct {
	Rows  []Row
	Dirty bool
}

// Row holds a single editor row
type Row struct {
	Raw    []rune
	Render []rune
}
