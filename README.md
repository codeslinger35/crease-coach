# Crease Coach
This repo will store all related code for the Crease Coach helper tool. This tool is meant to help ice hockey goalie coaches keep track of the goalies they work with. This tool is a work in progress and is not ready for any real production environment, but I am test driving the tool with my teams this season (23-24).

## Overview
This tool tracks seasons, games, periods per game, and coaching notes per period for each goalie in the system. Currently this is a simple JSON file stored at the project root, but I hope to expand this application to be a full-stack solution.

## Loading Data
the application expects a file called `goalies.json` in the src/public folder. A sample of that file looks like this:
```json
[
  {
    "id": 1,
    "firstName": "John",
    "lastName": "Doe",
    "age": "U19",
    "team": "My Jr. Team",
    "seasons": [
      {
        "id": 1,
        "title": "U19 Jr Canes",
        "description": "The 2023-24 season with the Jr Canes",
        "team": "My Jr. Team",
        "year": "2023-24",
        "games": [
          {
            "id": "1",
            "date": "08/25/23",
            "opponent": "Team To Beat",
            "started": true,
            "pulled": false,
            "periods": [
              {
                "periodNumber": 1,
                "shotsAgainst": 32,
                "saves": 28,
                "notes": "Great start to the game!",
                "coachingPoints": [
                  {
                    "id": 1,
                    "time": "P2: 2:48",
                    "event": "Shot",
                    "notes": "Good tracking and rebound to the corner",
                    "url": "https://google.com",
                    "tags": ["rebound-control"]
                  }
                ]
              }
            ]
          }
        ]
      }
    ]
  },
]
```
