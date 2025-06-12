package cmd

import (
	"fmt"
	"strings"

	"github.com/raithbheart/whack-a-goph/internal"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all exercises and their status",
	Run: func(cmd *cobra.Command, args []string) {
		exercises, err := internal.GetExercises()
		if err != nil {
			fmt.Printf("%s Error loading exercises: %v\n", red("✗"), err)
			return
		}

		fmt.Printf("%s\n", cyan(gopher))
		fmt.Printf("%s Exercise List\n", yellow("🐹"))
		fmt.Println(strings.Repeat("─", 60))

		completedCount := 0
		for _, exercise := range exercises {
			status := red("🕳️  HOLE")
			if exercise.Completed {
				status = green("🎯 WHACKED")
				completedCount++
			}

			fmt.Printf("%-15s %s %s\n",
				bold(exercise.Name),
				status,
				exercise.Description)
		}

		fmt.Println(strings.Repeat("─", 60))
		fmt.Printf("Progress: %s/%d exercises completed\n",
			green(fmt.Sprintf("%d", completedCount)),
			len(exercises))

		if completedCount < len(exercises) {
			fmt.Printf("\nRun %s to continue!\n", green("whack-a-goph next"))
		} else {
			fmt.Printf("\n%s All gophers whacked! Great job! 🎉\n", green("🎯"))
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
