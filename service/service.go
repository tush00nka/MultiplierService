package service

import "math/rand"

type GeneratorService struct {
	rtp float64
}

func NewGenerator(rtp float64) *GeneratorService {
	return &GeneratorService{rtp: rtp}
}

func (g *GeneratorService) Get() float64 {
	// adjust random float according to rtp value
	m := g.rtp / (1 - rand.Float64())

	// cap in range [1;10_000]
	m = max(min(m, 10_000), 1)

	return m
}
