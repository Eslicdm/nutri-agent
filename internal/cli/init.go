package cli

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Eslicdm/nutri-agent/internal/agent"
	"github.com/Eslicdm/nutri-agent/internal/nutrition"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inicia o shell interativo do NutriAgent",
	Long:  `Abre o shell do NutriAgent onde você pode usar comandos como 'compare' continuamente.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		fmt.Println("⏳ Carregando catálogo local...")
		catalog, err := nutrition.LoadCatalog(os.Getenv("DATABASE_PATH"))
		if err != nil {
			log.Fatalf("Erro ao carregar catálogo: %v", err)
		}

		fmt.Println("🧠 Iniciando Agente ADK...")
		nutriAgent, err := agent.GetAgent(catalog)
		if err != nil {
			log.Fatalf("Erro ao iniciar agente: %v", err)
		}

		runShell(ctx, nutriAgent)
	},
}

func runShell(ctx context.Context, nutriAgent *agent.NutriAgent) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n✅ NutriAgent shell iniciado!")
	fmt.Println("💡 Comandos disponíveis:")
	fmt.Println("   compare <alimento>  — analisa e sugere alternativas")
	fmt.Println("   help                — mostra esta ajuda")
	fmt.Println("   exit / sair         — encerra o shell")

	for {
		fmt.Print("\nnutriagent> ")

		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, " ", 2)
		command := strings.ToLower(parts[0])
		argument := ""
		if len(parts) > 1 {
			argument = strings.TrimSpace(parts[1])
		}

		switch command {
		case "exit", "sair":
			fmt.Println("Até logo! 👋")
			return

		case "help":
			fmt.Println("💡 Comandos disponíveis:")
			fmt.Println("   compare <alimento>  — analisa e sugere alternativas")
			fmt.Println("   help                — mostra esta ajuda")
			fmt.Println("   exit / sair         — encerra o shell")

		case "compare":
			if argument == "" {
				fmt.Println("⚠️  Uso: compare <descrição do alimento>")
				fmt.Println("   Exemplo: compare peito de frango 100g R$12,00")
				continue
			}
			fmt.Println("⏳ Analisando...")
			response, err := nutriAgent.Run(ctx, argument)
			if err != nil {
				fmt.Printf("❌ Erro ao processar: %v\n", err)
				continue
			}
			agent.RenderResponse(response.Response)

		default:
			fmt.Printf("❓ Comando desconhecido: '%s'. Digite 'help' para ver os comandos.\n", command)
		}
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
