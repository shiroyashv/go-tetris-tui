package game

import (
	"math/rand"
	"time"

	"github.com/shiroyashv/go-tetris-tui/internal/config"
)

type Game struct {
	Grid [config.BoardHeight][config.BoardWidth]int

	Piece     CurrentPiece
	NextPiece CurrentPiece

	Score    int
	GameOver bool
	Paused   bool
	ConfirmRestart bool
	TickRate time.Duration
}

func NewGame() *Game {
	g := &Game{
		TickRate: time.Millisecond * 800,
	}

	preset := AllPieces[rand.Intn(len(AllPieces))]
	g.NextPiece = CurrentPiece{
		Shape: preset.Shape,
		Color: preset.Color,
	}

	g.SpawnPiece()
	return g
}

func (g *Game) SpawnPiece() {

	g.Piece = g.NextPiece
	g.Piece.X = config.BoardWidth/2 - 2
	g.Piece.Y = -2

	preset := AllPieces[rand.Intn(len(AllPieces))]
	g.NextPiece = CurrentPiece{
		Shape: preset.Shape,
		Color: preset.Color,
	}
}

func (g *Game) Update() {
	if g.GameOver {
		return
	}

	if !g.CheckCollision(g.Piece.X, g.Piece.Y+1, g.Piece.Shape) {
		g.Piece.Y++
	} else {
		g.LockPiece()
		g.ClearLines()
	}
}

func (g *Game) CheckCollision(x, y int, shape Tetromino) bool {
	for row := 0; row < len(shape); row++ {
		for col := 0; col < len(shape[row]); col++ {

			if shape[row][col] == 0 {
				continue
			}

			boardX := x + col
			boardY := y + row

			if boardX < 0 || boardX >= config.BoardWidth || boardY >= config.BoardHeight {
				return true
			}

			if boardY >= 0 && g.Grid[boardY][boardX] != 0 {
				return true
			}
		}
	}
	return false
}

func (g *Game) ClearLines() {
	linesCleared := 0

	for y := config.BoardHeight - 1; y >= 0; y-- {

		full := true
		for x := 0; x < config.BoardWidth; x++ {
			if g.Grid[y][x] == 0 {
				full = false
				break
			}
		}

		if full {
			linesCleared++

			for k := y; k > 0; k-- {
				g.Grid[k] = g.Grid[k-1]
			}

			g.Grid[0] = [config.BoardWidth]int{}

			y++
		}
	}

	if linesCleared > 0 {
		g.Score += linesCleared * 100 * linesCleared

		newRate := g.TickRate - time.Duration(linesCleared*20)*time.Millisecond
		if newRate < 100*time.Millisecond {
			newRate = 100 * time.Millisecond
		}
		g.TickRate = newRate
	}
}

func (g *Game) LockPiece() {
	for row := 0; row < len(g.Piece.Shape); row++ {
		for col := 0; col < len(g.Piece.Shape[row]); col++ {
			if g.Piece.Shape[row][col] == 1 {
				boardX := g.Piece.X + col
				boardY := g.Piece.Y + row

				if boardY < 0 {
					g.GameOver = true
					return
				}

				if boardY >= 0 && boardY < config.BoardHeight &&
					boardX >= 0 && boardX < config.BoardWidth {
					g.Grid[boardY][boardX] = g.Piece.Color
				}
			}
		}
	}
	g.ClearLines()

	g.SpawnPiece()

	if g.CheckCollision(g.Piece.X, g.Piece.Y, g.Piece.Shape) {
		g.GameOver = true
	}
}

func (g *Game) MoveLeft() {
	if g.GameOver {
		return
	}
	if !g.CheckCollision(g.Piece.X-1, g.Piece.Y, g.Piece.Shape) {
		g.Piece.X--
	}
}

func (g *Game) MoveRight() {
	if g.GameOver {
		return
	}
	if !g.CheckCollision(g.Piece.X+1, g.Piece.Y, g.Piece.Shape) {
		g.Piece.X++
	}
}

func (g *Game) Rotate() {
	if g.GameOver {
		return
	}

	originalShape := g.Piece.Shape
	rows := len(originalShape)
	cols := len(originalShape[0])

	newShape := make(Tetromino, rows)
	for i := range newShape {
		newShape[i] = make([]int, cols)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			newShape[c][rows-1-r] = originalShape[r][c]
		}
	}

	kicks := []struct{ x, y int }{
		{0, 0},
		{1, 0},
		{-1, 0},
		{2, 0},
		{-2, 0},
	}

	for _, k := range kicks {

		if !g.CheckCollision(g.Piece.X+k.x, g.Piece.Y+k.y, newShape) {
			g.Piece.Shape = newShape
			g.Piece.X += k.x
			g.Piece.Y += k.y
			return
		}
	}
}

func (g *Game) HardDrop() {
	dropHeight := 0
	
	for {
		if g.CheckCollision(g.Piece.X, g.Piece.Y+1, g.Piece.Shape) {
			break
		}
		g.Piece.Y++
		dropHeight++
	}

	g.Score += dropHeight * 2

	g.LockPiece()
	
	if !g.GameOver {
		g.ClearLines()
	}
}
