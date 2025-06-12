package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	green  = color.New(color.FgGreen).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
	bold   = color.New(color.Bold).SprintFunc()
)

// ASCII Art
const gopher = `
    🐹  WHACK-A-GOPH!  🐹
`

const gopherSuccess = `
        🎯
        🐹 *WHACKED!*
`

const gopherFail = `
        🐹 *WHOOPS!*
`

// RootCmd represents the base command
var RootCmd = &cobra.Command{
	Use:   "whack-a-goph",
	Short: "🐹 Small exercises to get you used to reading and writing Go code!",
	Long: fmt.Sprintf(`%s
%s

Whack-a-Goph is a collection of small Go exercises designed to get you used to 
reading and writing Go code. Each exercise is a "gopher" popping up from a hole - 
fix the code to make the tests pass and "whack" the gopher!

Based on the fantastic Rustlings project but tailored for Go programmers.

Available commands:
  list    - Show all exercises and their status
  run     - Run a specific exercise  
  watch   - Watch mode for an exercise
  next    - Run the next pending exercise
  reset   - Reset progress for an exercise
  help    - Show help information

Get started with: %s`,
		cyan(gopher),
		yellow("Welcome to Whack-a-Goph!"),
		green("whack-a-goph list")),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", cyan(gopher))
		fmt.Printf("%s\n\n", yellow("Welcome to Whack-a-Goph!"))
		fmt.Printf("Run %s to see all available exercises.\n", green("whack-a-goph list"))
		fmt.Printf("Run %s to start with the first exercise.\n", green("whack-a-goph next"))
		fmt.Printf("Run %s for more information.\n", green("whack-a-goph help"))
	},
}

func init() {
	cobra.OnInitialize()
}
