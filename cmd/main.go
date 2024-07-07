package main

import (
	"strconv"

	"github.com/brian-nunez/elo-now/cmd/elo"
	"github.com/gin-gonic/gin"
)

type RatingResponse struct {
	PrevRating int64 `json:"prevRating"`
	Diff       int64 `json:"diff"`
	NewRating  int64 `json:"newRating"`
}

func main() {
	r := gin.Default()

	r.GET("/elo/rating", func(c *gin.Context) {
		rating1Param := c.Query("rating1")
		rating2Param := c.Query("rating2")
		outcomeParam := c.Query("outcome")

		r1Int, err := strconv.ParseInt(rating1Param, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "rating1 is not a valid number",
			})
			return
		}

		r2Int, err := strconv.ParseInt(rating2Param, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "rating2 is not a valid number",
			})
			return
		}

		outcomeInt, err := strconv.ParseFloat(outcomeParam, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "outcome is not a valid number",
			})
			return
		}

		eloItem := elo.CreateBaseElo()

		rating1, rating2 := eloItem.CalculateRatings(
			r1Int,
			r2Int,
			outcomeInt,
		)

		c.JSON(200, gin.H{
			"data": []RatingResponse{
				{
					PrevRating: r1Int,
					Diff:       rating1 - r1Int,
					NewRating:  rating1,
				},
				{
					PrevRating: r2Int,
					Diff:       rating2 - r2Int,
					NewRating:  rating2,
				},
			},
		})
	})

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
