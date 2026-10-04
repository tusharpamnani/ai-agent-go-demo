# goaiagent

A minimal CLI chat agent using OpenRouter with the OpenAI Go SDK.

## Features

- Interactive REPL for chatting with an LLM via OpenRouter
- Colored terminal output using `fatih/color`
- Automatic conversation history (keeps the last 4 messages after reaching 10)
- Built-in commands: `clear` (clear screen) and `exit` (quit)

## Prerequisites

- Go 1.27.1 or higher
- An [OpenRouter](https://openrouter.ai/) API key set in your environment

## Setup

```bash
# Set your API key
export OPENROUTER_API_KEY="sk-or-v1-..."

# Download dependencies
go mod tidy
```

## Usage

### Run directly

```bash
go run .
```

### Build and run

```bash
go build -o goaiagent
./goaiagent
```

Once running, type a message and press Enter. Use `clear` to reset the screen or `exit` to quit.

## Configuration

- **Model**: Default is `google/gemma-4-26b-a4b-it:free`. You can change this in `main.go` (line 35).
- **API Base URL**: Uses `https://openrouter.ai/api/v1` by default.

## Dependencies

- [github.com/fatih/color](https://github.com/fatih/color) - Terminal colors
- [github.com/openai/openai-go/v2](https://github.com/openai/openai-go) - OpenAI-compatible Go SDK