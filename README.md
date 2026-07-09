# NutriAgent

[![Go](https://img.shields.io/badge/Go-1.26-blue.svg)](https://go.dev/)
[![Google Gemini](https://img.shields.io/badge/Gemini-LLM-red.svg)](https://ai.google.dev/)
[![ADK](https://img.shields.io/badge/Google-ADK-green.svg)](https://adk.dev/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## 📖 Overview

NutriAgent is a Go-based Command Line Interface (CLI) application powered by an AI Agent. It receives natural language descriptions of food items—including nutritional value and price—and recommends more cost-effective or practical alternatives from a local knowledge base. The system focuses on high nutritional density, practical preparation, and cost-efficiency by comparing similar macro categories (protein vs protein, carbs vs carbs, fats vs fats).

## 🚀 Key Features

*   **AI-Powered Recommendations:** Uses **Google Gemini via ADK (Agent Development Kit)** to analyze food items and suggest optimal alternatives.
*   **Cost-Benefit Analysis:** Calculates cost per gram of primary macro and cost per 100kcal for accurate comparisons.
*   **Interactive Shell:** Continuous analysis session via the `init` command for seamless food comparisons.
*   **Local Knowledge Base:** JSON-based catalog of food products with detailed nutritional data.
*   **Rich Terminal Output:** Beautiful Markdown rendering using **Glamour** and **Lipgloss** for enhanced readability.
*   **Natural Language Input:** Supports flexible prompts in both Portuguese and English.

## 🏗 Architecture

The project follows the **Standard Go Project Layout** with strict separation of concerns:

```
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
```

## 🛠 Tech Stack

### Core
*   **Language:** Go 1.26
*   **CLI Framework:** Cobra (`github.com/spf13/cobra`)
*   **Terminal UI:** Lipgloss, Glamour (`github.com/charmbracelet/*`)

### AI & LLM
*   **Agent Framework:** Google ADK (`google.golang.org/adk`)
*   **LLM Provider:** Google Gemini (`google.golang.org/genai`)

### Data & Configuration
*   **Data Storage:** Local JSON file (`data/catalog.json`)
*   **Environment:** Godotenv (`github.com/joho/godotenv`)

## 📋 Prerequisites

*   **Go 1.26** or later
*   **Google Gemini API Key** (Get one at [ai.google.dev](https://ai.google.dev/))

## 🐳 Getting Started

### 1. Environment Variables

Create a `.env` file in the project root:

```bash
cp .env.example .env
```

Edit `.env` and add your API key:

```env
LLM_API_KEY=your_gemini_api_key_here
LLM_MODEL=gemini-3-flash-preview
DATABASE_PATH=data/catalog.json
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the Application

```bash
go run cmd/nutriagent/main.go init
```

### 4. Using NutriAgent

Once inside the `nutriagent>` shell, use the `compare` command:

*   **By Price:** `compare peito de frango 1kg 30 reais`
*   **By Macros:** `compare dose de whey 30g com 24g de proteina custando 5 reais`
*   **Natural Language:** `compare 3 ovos cozidos no café da manhã`

#### Available Commands

| Command | Description |
| :--- | :--- |
| `compare <food>` | Analyzes food and suggests alternatives |
| `help` | Shows available commands |
| `exit` / `sair` | Exits the shell |

## 📊 Example Output

```
nutriagent> compare peito de frango 1kg 30 reais

--- ANÁLISE NUTRICIONAL ---

Ovo em Pó é uma melhor alternativa:
- Peito de frango: R$0.13 por g de proteína
- Ovo em Pó: R$0.21 por g de proteína (rica em vitaminas e minerais)
```

## 📁 Data Catalog

The `data/catalog.json` file contains food items with the following structure:

```json
{
  "name": "Product Name",
  "total_weight_g": 100.0,
  "portion_g": 100.0,
  "carb_g": 0.0,
  "protein_g": 24.0,
  "fat_g": 10.0,
  "fiber_g": 0.0,
  "calories": 186.0,
  "average_price": 5.50,
  "category": "protein",
  "strength": "High Omega-3, calcium and B12"
}
```

### Categories

*   `protein` - High protein foods (meat, eggs, whey, etc.)
*   `carbs` - Carbohydrate-rich foods (oats, rice, etc.)
*   `fat` - Healthy fat sources (nuts, oils, etc.)

## 🔧 Development

### Project Structure

*   `cmd/` - Application entry points
*   `internal/` - Private application code
    *   `agent/` - AI agent configuration and execution
    *   `cli/` - Command-line interface commands
    *   `nutrition/` - Data models and repository
*   `data/` - Static data files

### Adding New Foods

Edit `data/catalog.json` to add new food items to the catalog:

```json
{
  "name": "New Food Item",
  "total_weight_g": 100.0,
  "portion_g": 100.0,
  "carb_g": 0.0,
  "protein_g": 0.0,
  "fat_g": 0.0,
  "fiber_g": 0.0,
  "calories": 0.0,
  "average_price": 0.0,
  "category": "protein",
  "strength": "Description of nutritional strengths"
}
```
