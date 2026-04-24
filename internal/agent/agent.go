package agent

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Eslicdm/nutri-agent/internal/nutrition"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/runner"
	"google.golang.org/adk/session"
	"google.golang.org/genai"
)

var (
	headerStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00FF00")).
		PaddingTop(1).
		PaddingBottom(1).
		Underline(true)
)

// NutriAgent wraps the ADK runner and session for direct CLI use
const (
	appName   = "nutri_agent_app"
	userID    = "cli-user"
	sessionID = "cli-session"
)

// NutriAgent wraps the ADK runner for direct CLI use
type NutriAgent struct {
	r *runner.Runner
}

// RunResult holds the agent's text response
type RunResult struct {
	Response string
}

// GetAgent builds the LLM agent and returns a NutriAgent ready for CLI use
func GetAgent(catalog []nutrition.Product) (*NutriAgent, error) {
	ctx := context.Background()

	model, err := gemini.NewModel(ctx, os.Getenv("LLM_MODEL"), &genai.ClientConfig{
		APIKey: os.Getenv("LLM_API_KEY"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create model: %w", err)
	}

	catalogStr := "Local Catalog:\n"
	for _, p := range catalog {
		catalogStr += fmt.Sprintf(
			"- %s (%s): %.0fkcal, P: %.1fg, C: %.1fg, F: %.1fg. Price: R$%.2f per %.1fg. Notes: %s\n",
			p.Name, p.Category, p.Calories, p.ProteinG, p.CarbG, p.FatG, p.AveragePrice, p.PortionG, p.Strength,
		)
	}

	nutriAgent, err := llmagent.New(llmagent.Config{
		Name:        "nutri_agent",
		Model:       model,
		Description: "Expert nutritionist that suggests cost-effective food alternatives.",
		Instruction: "You are a cost-benefit focused nutritionist. Your task is to compare user-provided food with the catalog.\n" +
			"CRITICAL LOGIC:\n" +
			"1. Identify the primary macro based on the catalog category (protein, carbs, or fat).\n" +
			"2. Calculate the 'Cost per Gram' of that primary macro (Price / Grams of Macro per portion) or 'Cost per 100kcal'.\n" +
			"3. Factor in the 'Strength' notes (micros, digestion) for the final recommendation.\n" +
			"4. Compare the user's input price/macros against the catalog's efficiency.\n" +
			"Be concise and professional. Catalog:\n" + catalogStr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create agent: %w", err)
	}

	sessionSvc := session.InMemoryService()

	_, err = sessionSvc.Create(ctx, &session.CreateRequest{
		AppName:   appName,
		UserID:    userID,
		SessionID: sessionID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	r, err := runner.New(runner.Config{
		AppName:        appName,
		Agent:          nutriAgent,
		SessionService: sessionSvc,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create runner: %w", err)
	}

	return &NutriAgent{r: r}, nil
}

// Run sends a single user message and returns the agent's text response
func (n *NutriAgent) Run(ctx context.Context, input string) (*RunResult, error) {
	msg := genai.NewContentFromText(input, "user")

	var sb strings.Builder

	for event, err := range n.r.Run(ctx, userID, sessionID, msg, agent.RunConfig{}) {
		if err != nil {
			return nil, fmt.Errorf("agent error: %w", err)
		}

		if event.LLMResponse.Content != nil && !event.LLMResponse.Partial {
			for _, part := range event.LLMResponse.Content.Parts {
				if part.Text != "" {
					sb.WriteString(part.Text)
				}
			}
		}
	}

	response := sb.String()
	if response == "" {
		response = "(no response)"
	}

	return &RunResult{Response: response}, nil
}

// RenderResponse formats the agent's response using glamour and lipgloss for a dark terminal
func RenderResponse(response string) {
	fmt.Println(headerStyle.Render("--- ANÁLISE NUTRICIONAL ---"))

	r, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(100),
	)
	if err != nil {
		fmt.Println(response)
		return
	}

	out, err := r.Render(response)
	if err != nil {
		fmt.Println(response)
		return
	}
	fmt.Print(out)
}
