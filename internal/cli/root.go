package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without subcommands
var rootCmd = &cobra.Command{
	Use:   "nutriagent",
	Short: "CLI para otimizar custo-benefício nutricional com Agentes IA",
	Long: `NutriAgent é uma ferramenta de linha de comando que analisa 
informações nutricionais e recomenda opções mais eficientes e baratas.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
