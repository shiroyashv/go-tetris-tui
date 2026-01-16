package game

type Tetromino [][]int

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

	AllPieces = []Tetromino{PieceI, PieceJ, PieceL, PieceO, PieceS, PieceT, PieceZ}
)

type CurrentPiece struct {
	Shape Tetromino
	X, Y  int
	Color int
}
