# Elo Now

This is a simple implementation of the Elo rating system in Go. One api endpoint is exposed to calculate the new ratings of two players after a match.

## Usage

To run the server, execute the following command:

```bash
go run ./cmd/main.go 
```

The server will be running on port 8080.

To calculate the new ratings of two players after a match, make a GET request to the `/elo/rating` endpoint.

```bash
GET /elo/rating?rating1=1200&rating2=1200&outcome=1
```

The query parameters are as follows:

- `rating1`: The rating of player 1
- `rating2`: The rating of player 2
- `outcome`: The outcome of the match. 1 if player 1 won, 0 if player 2 won, 0.5 if it was a draw.

The response will be a JSON object with the new ratings of the two players.

```json
{
    "data": [
        {
            "prevRating": 1200,
            "diff": 16,
            "newRating": 1216
        },
        {
            "prevRating": 1200,
            "diff": -16,
            "newRating": 1184
        }
    ]
}
```
