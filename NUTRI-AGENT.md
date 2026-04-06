# NutriAgent: Project Specification & Context

## 1. Project Overview
**NutriAgent** is a Go-based Command Line Interface (CLI) application powered by an AI Agent. Its primary purpose is to receive a natural language description of a food item (including its nutritional value and price) and recommend a more cost-effective or practical alternative from a local knowledge base. It features an interactive shell for continuous analysis.

The project focuses on high nutritional density, practical preparation, and cost-efficiency (protein based food versus another protein based food, same with carbs, fats and micronutrients).

Prompt example: `compare peito de frango 1kg 20 reais` (inside the shell)
Output: Ovo é uma melhor substituição (ovo x reais por proteina e peito de frango x reais por proteina)


## 2. Tech Stack & Standards
* **Language:** Go
* **CLI Framework:** `github.com/spf13/cobra`
* **Terminal UI:** `github.com/charmbracelet/lipgloss` and `github.com/charmbracelet/glamour` (Markdown rendering)
* **AI Agent Framework:** Agent Development Kit (ADK) from `adk.dev` (Agent-First architecture).
* **LLM:** Google Gemini via `google.golang.org/genai`.
* **Data Storage:** Local JSON file (`data/catalog.json`).
* **Architecture Pattern:** Standard Go Project Layout (strict separation of `cmd`, `internal/cli`, `internal/agent`, and `internal/nutrition`).

## 3. Current Project State
**Current Progress Marker:** The interactive shell is implemented via the `init` command. The AI Agent is fully integrated with the local catalog and supports rich terminal output.

## 4. Directory Structure
```text
nutri-agent/
├── cmd/
│   └── nutriagent/
│       └── main.go                 # Entry point, loads .env and calls cli.Execute()
├── internal/
│   ├── agent/
│   │   ├── agent.go                # ADK setup, LLM configuration, and UI rendering
│   │   └── tools.go                # Functions to be registered as ADK Tools (Pending)
│   ├── cli/
│   │   ├── root.go                 # Cobra root command
│   │   └── init.go                 # 'init' command: manages the interactive shell loop
│   └── nutrition/
│       └── repository.go           # Data models and JSON parsing
├── data/
│   └── catalog.json               # Local database of alternative foods
├── .env                           # Environment variables (API Keys)
├── go.mod
└── go.sum