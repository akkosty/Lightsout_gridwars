package api

// Bot constants
const (
	BotName = "Lightsout Gridwars Bot"
	Version = "1.0.0"
	Author  = "Gridwars Team"

	// Menu options
	MenuStart    = "/start"
	MenuInfo     = "Инфо"
	MenuRegister = "Зарегистрироваться"
	MenuBack     = "Назад"

	// Command constants
	CommandStart    = "/start"
	CommandHelp     = "/help"
	CommandRegister = "/register"

	// Error codes
	ErrorCodeBadRequest   = 400
	ErrorCodeUnauthorized = 401
	ErrorCodeNotFound     = 404
	ErrorCodeConflict     = 409
	ErrorCodeInternal     = 500

	// Default values
	DefaultGridSize = 5
	MaxGridSize     = 10
	MinGridSize     = 3
)
