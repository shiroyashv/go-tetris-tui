package game

// Tetromino — матрица фигуры
type Tetromino [][]int

// PiecePreset — шаблон фигуры
type PiecePreset struct {
	Shape Tetromino
	Color int
}

const (
	ColorCyan   = 1 // I
	ColorBlue   = 2 // J
	ColorOrange = 3 // L
	ColorYellow = 4 // O
	ColorGreen  = 5 // S
	ColorPurple = 6 // T
	ColorRed    = 7 // Z
)

// Текущая фигура в игре
type CurrentPiece struct {
	Shape    Tetromino
	X, Y     int
	Color    int
	Rotation int // 0=0°, 1=90°, 2=180°, 3=270°
}

var (
	// I-Piece (Cyan)
	PieceI = Tetromino{
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	// J-Piece (Blue)
	PieceJ = Tetromino{
		{1, 0, 0},
		{1, 1, 1},
		{0, 0, 0},
	}

	// L-Piece (Orange)
	PieceL = Tetromino{
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	// O-Piece (Yellow)
	PieceO = Tetromino{
		{1, 1},
		{1, 1},
	}

	// S-Piece (Green)
	PieceS = Tetromino{
		{0, 1, 1},
		{1, 1, 0},
		{0, 0, 0},
	}

	// T-Piece (Purple)
	PieceT = Tetromino{
		{0, 1, 0},
		{1, 1, 1},
		{0, 0, 0},
	}

	// Z-Piece (Red)
	PieceZ = Tetromino{
		{1, 1, 0},
		{0, 1, 1},
		{0, 0, 0},
	}

	AllPieces = []PiecePreset{
		{PieceI, ColorCyan},
		{PieceJ, ColorBlue},
		{PieceL, ColorOrange},
		{PieceO, ColorYellow},
		{PieceS, ColorGreen},
		{PieceT, ColorPurple},
		{PieceZ, ColorRed},
	}
)