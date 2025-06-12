package cmd

import (
	"fmt"

	"github.com/raithbheart/whack-a-goph/internal"
	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Run the next pending exercise",
	Run: func(cmd *cobra.Command, args []string) {
		exercise, err := internal.GetNextExercise()
		if err != nil {
			fmt.Printf("%s %v\n", green("🎉"), err)
			return
		}

		fmt.Printf("%s Next exercise: %s\n", yellow("🐹"), bold(exercise.Name))
		fmt.Printf("Description: %s\n", exercise.Description)
		fmt.Printf("File: %s\n\n", cyan(exercise.Path))

		passed, output, err := internal.RunExercise(*exercise)
		if err != nil {
			fmt.Printf("%s Error running exercise: %v\n", red("✗"), err)
			return
		}

		if passed {
			fmt.Printf("%s\n", green(gopherSuccess))
			fmt.Printf("%s Exercise passed! 🎉\n", green("✓"))

			// Mark as completed
			if err := internal.MarkCompleted(exercise.Name); err != nil {
				fmt.Printf("%s Warning: Could not save progress: %v\n", yellow("⚠"), err)
			}

			// Check if there's another exercise
			if nextExercise, err := internal.GetNextExercise(); err == nil {
				fmt.Printf("\nNext up: %s - %s\n", bold(nextExercise.Name), nextExercise.Description)
				fmt.Printf("Run %s to continue!\n", green("whack-a-goph next"))
			} else {
				fmt.Printf("\n%s All exercises completed! 🎉\n", green("🎯"))
			}
		} else {
			fmt.Printf("%s\n", red(gopherFail))
			fmt.Printf("%s Exercise failed!\n", red("✗"))
			fmt.Printf("\nTest output:\n%s\n", output)

			if exercise.Hint != "" {
				fmt.Printf("\n%s Hint: %s\n", yellow("💡"), yellow(exercise.Hint))
			}

			fmt.Printf("\nFix the code and run %s again!\n", green("whack-a-goph next"))
		}
	},
}

func init() {
	RootCmd.AddCommand(nextCmd)
}
