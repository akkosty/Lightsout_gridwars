package main

import (
	"sync"
)

// StateManager manages all active games for different users
type StateManager struct {
	mu     sync.RWMutex
	games  map[int]*Game
	nextID int
}

// NewStateManager creates a new state manager
func NewStateManager() *StateManager {
	return &StateManager{
		games:  make(map[int]*Game),
		nextID: 1,
	}
}

// CreateGame creates a new game for a user
func (sm *StateManager) CreateGame(player string, size int) *Game {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	game := NewGame(size)
	game.ID = sm.nextID
	sm.nextID++
	game.Player = player
	
	sm.games[game.ID] = game
	return game
}

// GetGame gets a game by ID
func (sm *StateManager) GetGame(id int) (*Game, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	game, ok := sm.games[id]
	return game, ok
}

// UpdateGame updates a game in the state
func (sm *StateManager) UpdateGame(game *Game) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.games[game.ID] = game
}

// DeleteGame removes a game from the state
func (sm *StateManager) DeleteGame(id int) bool {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	_, exists := sm.games[id]
	if exists {
		delete(sm.games, id)
	}
	return exists
}
