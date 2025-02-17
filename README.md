# Matchmaking Service

- A simple matchmaking service that allows players to join a competition and wait for other players to join. 
- Service works as a TCP server that listens for connections from players.
- Service groups users into competitions matching their level
    - Level tolerance is configurable. 
    - When a player joins, the service will check if there is a competition available close to the player's level
    - If there is no competition in matchmaking, a new competition is created.
        - Level range for new competition is calculated based on the player's level and the tolerance defined by `-level-matching-tolerance`
        - For example: If the player's level is 5 and the tolerance is 3, the new competition will be created with a level range of 2-8

- Service will start a competition in following cases:
    - Maximum number of players is reached
    - Timeout for matchmaking occurs and minimum number of players is reached
- If not enough players are found for a competition, the competition will be aborted

- Upon competition start, the service will notify all players in the competition that the competition has started
- If competition is aborted, the service will notify all players in the competition that the competition has been aborted

## Running the service

`go run cmd/matchmaking-server/main.go`

### Flags

- `-port`: The port to listen on.
- `-max-players`: The maximum number of players that can join the competition.
- `-min-players`: The minimum number of players that must join the competition before it starts.
- `-timeout`: The timeout for the matchmaking in seconds.
- `-level-matching-tolerance`: The tolerance for the level matching in the competition.

### Example 
`go run main.go -port=8080 -max-players=3 -min-players=2 -timeout=15s -level-matching-tolerance=3`

### Joining to the matchmaking service as player
`
echo '{"Id" : "4", "Level": 4}' | nc localhost 8080
` 

## Tools used in the project
- IDE: [Cursor](https://www.cursor.com/) Claude 3.5 Sonnet set up as LLM


## Running tests
`go test ./...`
