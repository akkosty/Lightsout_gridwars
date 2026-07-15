package api

// Response represents a standard API response structure for Telegram bot
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *Error      `json:"error,omitempty"`
}

// Error represents an error response with code and details
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// User represents a Telegram user model
type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	IsBot     bool   `json:"is_bot"`
}

// GameStats represents user's game statistics
type GameStats struct {
	UserID      int64 `json:"user_id"`
	GamesPlayed int   `json:"games_played"`
	Wins        int   `json:"wins"`
	Losses      int   `json:"losses"`
	TotalScore  int   `json:"total_score"`
}

// GameState represents the current state of a game
type GameState struct {
	GameID     string `json:"game_id"`
	UserID     int64  `json:"user_id"`
	GridSize   int    `json:"grid_size"`
	Score      int    `json:"score"`
	Turn       int    `json:"turn"`
	IsActive   bool   `json:"is_active"`
	CreatedAt  string `json:"created_at,omitempty"`
	FinishedAt string `json:"finished_at,omitempty"`
}

// RegisterRequest represents a user registration request
type RegisterRequest struct {
	UserID    int64  `json:"user_id"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
}

// RegisterResponse represents the response for registration
type RegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	UserID  int64  `json:"user_id,omitempty"`
}

// Button represents a keyboard button in Telegram
type Button struct {
	Text string `json:"text"`
}

// KeyboardMarkup represents inline or reply keyboard markup
type KeyboardMarkup struct {
	Keyboard    [][]Button `json:"keyboard,omitempty"`
	ResizeKeyboard  bool `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`
	InlineKeyboard  [][]Button `json:"inline_keyboard,omitempty"`
}

// MessageRequest represents an incoming message from Telegram
type MessageRequest struct {
	MessageID int64   `json:"message_id"`
	User      *User   `json:"from,omitempty"`
	Text      string  `json:"text,omitempty"`
	ChatID    int64   `json:"chat_id"`
	Timestamp int64   `json:"date"`
}

// InfoResponse represents the info menu response
type InfoResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Author      string `json:"author"`
}
