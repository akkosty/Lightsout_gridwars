# Lightsout Grid Wars

An innovative, strategy-driven variant of the classic Lights Out game. Players must use logic and foresight to clear the grid by toggling lights in a specific pattern.

## Overview

**Lightsout_gridwars** is a modern take on the beloved puzzle game "Lights Out." While traditional versions focus on turning all lights off, this twist introduces strategic layers, competitive multiplayer elements, and dynamic grid configurations. The goal remains simple: manipulate the grid to achieve a target state (e.g., all lights off or on), but now with added complexity such as limited moves, power-ups, and time constraints.

This repository contains both the client-side game logic and a RESTful backend API built in Go, making it easy to deploy and integrate into web or mobile applications.

## Features

- **Dynamic Grid**: Choose from various grid sizes (5x5, 6x6, 8x8).
- **Game Modes**:
  - *Classic*: Turn all lights off.
  - *Pattern Match*: Replicate a given pattern.
  - *Multiplayer Arena*: Compete in real-time to solve the puzzle first.
- **Smart Hints**: AI-powered suggestions for next moves.
- **Progress Tracking**: Persistent storage of high scores and game history via API.

## Technology Stack

- **Frontend**: React (for web), Flutter (for mobile), or custom UIs.
- **Backend**:
  - Language: Go
  - Framework: Standard library `net/http` (lightweight, performant)
  - Database: SQLite (local) / PostgreSQL (production-ready via config)
- **Deployment**:
  - Docker for containerization
  - Render for cloud hosting

## Getting Started

### Prerequisites

- Go 1.20 or higher
- Docker and Docker Compose (optional, for containerized deployment)

### Local Development

1. Clone the repository:

   ```bash
   git clone https://github.com/akkosty/Lightsout_gridwars.git
   cd Lightsout_gridwars
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Run the server (default port: 8080):

   ```bash
   go run main.go
   ```

4. Access the game API documentation at `http://localhost:8080/docs`.

### Docker Deployment

1. Build and run with Docker Compose:

   ```bash
   docker-compose up --build
   ```

2. The server will be available at `http://localhost:8080`.

## API Endpoints

| Method | Endpoint             | Description                      |
|--------|----------------------|----------------------------------|
| POST   | `/api/v1/game/new`   | Start a new game with specified grid size. |
| GET    | `/api/v1/game/{id}`  | Get current state of a game by ID. |
| PUT    | `/api/v1/game/{id}/move` | Apply a move (toggle cell) to the game. |
| GET    | `/api/v1/leaderboard`| Retrieve top players and scores. |

*Full OpenAPI 3.0 documentation is available in `docs/openapi.yaml`.*

## Project Structure

```
.
├── docs/                  # API documentation
│   └── openapi.yaml
├── frontend/              # Optional: Web UI (React)
├── internal/
│   ├── game/             # Core game logic
│   ├── api/              # HTTP handlers and middleware
│   └── database/         # Data persistence layer
├── main.go               # Application entry point
├── go.mod / go.sum       # Go module files
├── Dockerfile            # Container definition
└── docker-compose.yml    # Local multi-container setup
```

## Contributing

We welcome contributions! To contribute:

1. Fork the repository.
2. Create a feature branch: `git checkout -b feature/your-feature`.
3. Commit your changes: `git commit -m 'Add some feature'`.
4. Push to the branch: `git push origin feature/your-feature`.
5. Open a Pull Request.

Please read [CONTRIBUTING.md](https://github.com/akkosty/Lightsout_gridwars/blob/main/CONTRIBUTING.md) for details on our code of conduct and development process.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/akkosty/Lightsout_gridwars/blob/main/LICENSE) file for details.

## Acknowledgments

- Inspired by the classic "Lights Out" puzzle game.
- Special thanks to all contributors and users of this repository.

---

**Enjoy playing Lightsout Grid Wars!**