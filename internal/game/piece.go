package game

type Tetromino [][]int

type PiecePreset struct {
	Shape Tetromino
	Color int
}

var (
	PieceI = Tetromino{
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	PieceJ = Tetromino{
		{1, 0, 0},
		{1, 1, 1},
		{0, 0, 0},
	}

	PieceL = Tetromino{
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	PieceO = Tetromino{
		{1, 1},
		{1, 1},
	}

	PieceS = Tetromino{
		{0, 1, 1},
		{1, 1, 0},
		{0, 0, 0},
	}

	PieceT = Tetromino{
		{0, 1, 0},
		{1, 1, 1},
		{0, 0, 0},
	}

	PieceZ = Tetromino{
		{1, 1, 0},
		{0, 1, 1},
		{0, 0, 0},
	}

	AllPieces = []PiecePreset{{PieceI, 1}, {PieceJ, 2}, {PieceL, 3}, {PieceO, 4}, {PieceS, 5}, {PieceT, 6}, {PieceZ, 7}}
)

type CurrentPiece struct {
	Shape Tetromino
	X, Y  int
	Color int
}
