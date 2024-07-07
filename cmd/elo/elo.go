package elo

import (
	"math"
)

type Elo struct {
	K float64
	D float64
}

func (e *Elo) CalculateProbability(rating1 int64, rating2 int64) float64 {
	denominator := math.Pow(10, (float64(rating2-rating1)/e.D)) + 1
	return 1 / denominator
}

func (e *Elo) CalculateRating(rating int64, outcome float64, expected float64) int64 {
	result := float64(rating) + float64(float64(e.K)*(float64(outcome)-expected))

	rounded := math.Round(result)

	return int64(rounded)
}

func (e *Elo) CalculateRatings(rating1 int64, rating2 int64, outcome float64) (int64, int64) {
	prob1 := e.CalculateProbability(rating1, rating2)
	result1 := e.CalculateRating(rating1, outcome, prob1)
	result2 := rating1 - result1 + rating2

	return result1, int64(result2)
}

func CreateBaseElo() *Elo {
	elo := &Elo{
		K: 32,
		D: 400,
	}

	return elo
}
