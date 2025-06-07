# Go LeetCode API

[![Live Demo](https://img.shields.io/badge/Live%20Demo-View%20Live-brightgreen)](https://go-leetcode-api.onrender.com/ping)

This project is a Go-based API that interacts with the LeetCode GraphQL API to fetch user profiles, submission statistics, and other related data.

## Project Structure

```
go-leetcode-api/
├── handler/
│ ├── leetcode.go
│ ├── query.go
│ ├── response.go
├── loggger/
│ ├── loggger.go
├── router/
│ ├── router.go
├── tmp/
├── .air.toml
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Installation

1. Clone the repository:

```sh
git clone https://github.com/etharrra/go-leetcode-api.git
cd go-leetcode-api
```

2. Install dependencies:

```sh
go mod tidy
```

## Usage

1. Run the application:

```sh
go run main.go
```

2. The API will be available at `http://localhost:3000`.

## Endpoints

-   `GET /ping`: Health check endpoint.
-   `GET /api/:username`: Fetches the user profile and submission statistics for the specified LeetCode username.

## Example

To fetch the profile and submission statistics for a user with the username `etharrra`, make a GET request to:

```
http://localhost:3000/api/etharrra
```

## Logging

Logs are written to log.log. The loggger package handles the logging functionality.

## Configuration

The project uses [Air](https://github.com/cosmtrek/air) for live reloading during development. Configuration is provided in the .air.toml file.

## Contributing

If you'd like to contribute to this project, feel free to fork the repository and submit a pull request. Please ensure that your code follows the existing style and includes appropriate tests.

---

Happy coding! 🚀
