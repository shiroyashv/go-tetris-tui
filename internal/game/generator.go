package game

import (
	"math/rand"
	"time"
)

type Generator struct {
	bag []int     
	rnd *rand.Rand 
}

func NewGenerator() *Generator {
	src := rand.NewSource(time.Now().UnixNano())
	return &Generator{
		bag: make([]int, 0),
		rnd: rand.New(src),
	}
}

func (gen *Generator) GetNewPiece() PiecePreset {
	if len(gen.bag) == 0 {
		gen.fillBag()
	}

	nextIdx := gen.bag[0]
	
	gen.bag = gen.bag[1:]

	return AllPieces[nextIdx]
}

func (gen *Generator) fillBag() {
	perm := gen.rnd.Perm(len(AllPieces)) 
	gen.bag = append(gen.bag, perm...)
}