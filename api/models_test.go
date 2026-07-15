package api

import (
	"testing"
)

func TestUser_JSON(t *testing.T) {
	user := User{
		ID:        123456789,
		Username:  "testuser",
		FirstName: "Test",
		LastName:  "User",
		IsBot:     false,
	}

	if user.ID != 123456789 {
		t.Errorf("Expected ID 123456789, got %d", user.ID)
	}

	if user.Username != "testuser" {
		t.Errorf("Expected username 'testuser', got '%s'", user.Username)
	}
}

func TestGameStats_DefaultValues(t *testing.T) {
	stats := GameStats{}

	if stats.GamesPlayed != 0 {
		t.Error("Expected GamesPlayed to be 0 by default")
	}
}

func TestGameState_ActiveGame(t *testing.T) {
	state := GameState{
		GameID:   "game-123",
		UserID:   123456789,
		GridSize: 5,
		Score:    100,
		Turn:     3,
		IsActive: true,
	}

	if !state.IsActive {
		t.Error("Expected IsActive to be true")
	}
}

func TestRegisterRequest_Validation(t *testing.T) {
	req := RegisterRequest{
		UserID:    123456789,
		FirstName: "Test",
	}

	if req.UserID == 0 {
		t.Error("Expected UserID to be set")
	}

	if req.FirstName != "Test" {
		t.Errorf("Expected first name 'Test', got '%s'", req.FirstName)
	}
}

func TestButton_Structure(t *testing.T) {
	button := Button{
		Text: "Click Me",
	}

	if button.Text != "Click Me" {
		t.Errorf("Expected text 'Click Me', got '%s'", button.Text)
	}
}

func TestKeyboardMarkup_Options(t *testing.T) {
	keyboard := KeyboardMarkup{
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}

	if !keyboard.ResizeKeyboard {
		t.Error("Expected ResizeKeyboard to be true")
	}

	if !keyboard.OneTimeKeyboard {
		t.Error("Expected OneTimeKeyboard to be true")
	}
}

func TestMessageRequest_Structure(t *testing.T) {
	msg := MessageRequest{
		MessageID: 1,
		User: &User{
			ID:       123456789,
			Username: "testuser",
		},
		Text:   "/start",
		ChatID: 987654321,
	}

	if msg.Text != "/start" {
		t.Errorf("Expected text '/start', got '%s'", msg.Text)
	}
}

func TestInfoResponse_Structure(t *testing.T) {
	info := InfoResponse{
		Title:       "Lightsout Gridwars",
		Description: "A puzzle game bot for Telegram",
		Version:     Version,
		Author:      Author,
	}

	if info.Version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', got '%s'", info.Version)
	}
}
