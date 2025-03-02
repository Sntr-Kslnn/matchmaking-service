package matchmaking

import (
	"time"

	"github.com/SntrKslnn/matchmaking-service/internal/competition"
	"github.com/SntrKslnn/matchmaking-service/internal/model"
)

type MatchmakingService interface {
	// HandlePlayerJoin handles a player's request to join matchmaking and returns a notification channel
	// that will receive updates about competition matching
	HandlePlayerJoin(playerData model.PlayerData) <-chan MatchMakingNotification
}

// MatchmakingConfig is the configuration for the matchmaking service
type MatchmakingConfig struct {
	LevelMatchingTolerance int
	MatchmakingTimeout     time.Duration
	CompetitionConfig      competition.CompetitionConfig
}

// MatchMakingNotification is a notification that is sent to the player to keep them updated about the matchmaking process
type MatchMakingNotification struct {
	CompetitionID int
	State         MatchmakingState
}

type MatchmakingState string

type matchmakingStateChangeOrigin string

const (
	matchmakingStateChangeOrigin_PlayerAdd matchmakingStateChangeOrigin = "player_added"
	matchmakingStateChangeOrigin_Timeout   matchmakingStateChangeOrigin = "matchmaking_timeout"
)

const (
	// CompetitionState_WaitingForPlayers indicates that the competition is open and accepting new players
	State_WaitingForPlayers MatchmakingState = "waiting_for_players"

	// CompetitionState_Started indicates that the competition has started
	State_Started MatchmakingState = "started"

	// CompetitionState_Aborted indicates that the competition has been aborted
	State_Aborted MatchmakingState = "aborted"
)

// NewMatchmakingService creates a new matchmaking service
// @param config the configuration of the matchmaking service
// @return a new matchmaking service
func NewMatchmakingService(config MatchmakingConfig) MatchmakingService {
	return newMatchmakingService(config)
}

func newMatchmakingService(config MatchmakingConfig) *matchmakingService {
	matchmakingService := &matchmakingService{
		competitionsInMatchmaking: make(map[int]competitionData),
		playersInMatchmaking:      make(map[string]playerInMatchmaking),
		competitionJoinRequests:   make(chan playerInMatchmaking),
		nextCompetitionID:         1,
		config:                    config,
		stateMutationChan:         make(chan stateChangeNotification),
	}
	matchmakingService.start()
	return matchmakingService
}

func (m *matchmakingService) HandlePlayerJoin(playerData model.PlayerData) <-chan MatchMakingNotification {
	return m.handlePlayerJoin(playerData)
}
