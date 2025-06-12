package cmd

import (
	"fmt"

	"github.com/raithbheart/whack-a-goph/internal"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [exercise_name]",
	Short: "Run a specific exercise",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		exerciseName := args[0]
		exercise, err := internal.FindExercise(exerciseName)
		if err != nil {
			fmt.Printf("%s %v\n", red("✗"), err)
			return
		}

		fmt.Printf("%s Running exercise: %s\n", yellow("🐹"), bold(exercise.Name))
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
		} else {
			fmt.Printf("%s\n", red(gopherFail))
			fmt.Printf("%s Exercise failed!\n", red("✗"))
			fmt.Printf("\nTest output:\n%s\n", output)

			if exercise.Hint != "" {
				fmt.Printf("\n%s Hint: %s\n", yellow("💡"), yellow(exercise.Hint))
			}
		}
	},
}

var hintCmd = &cobra.Command{
	Use:   "hint [exercise_name]",
	Short: "Get a hint for a specific exercise",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		exerciseName := args[0]
		exercise, err := internal.FindExercise(exerciseName)
		if err != nil {
			fmt.Printf("%s %v\n", red("✗"), err)
			return
		}

		fmt.Printf("%s Hint for %s:\n", yellow("💡"), bold(exercise.Name))
		if exercise.Hint != "" {
			fmt.Printf("%s\n", yellow(exercise.Hint))
		} else {
			fmt.Printf("No hint available for this exercise.\n")
		}
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(hintCmd)
}
