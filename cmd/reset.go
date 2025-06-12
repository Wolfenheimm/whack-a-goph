package cmd

import (
	"fmt"

	"github.com/raithbheart/whack-a-goph/internal"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset [exercise_name]",
	Short: "Reset progress for a specific exercise",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		exerciseName := args[0]
		exercise, err := internal.FindExercise(exerciseName)
		if err != nil {
			fmt.Printf("%s %v\n", red("✗"), err)
			return
		}

		err = internal.ResetExercise(exercise.Name)
		if err != nil {
			fmt.Printf("%s Error resetting exercise: %v\n", red("✗"), err)
			return
		}

		fmt.Printf("%s Exercise %s has been reset to pending status\n",
			green("✓"), bold(exercise.Name))
		fmt.Printf("The gopher is back in its hole! 🕳️\n")
	},
}

func init() {
	RootCmd.AddCommand(resetCmd)
	RootCmd.AddCommand(watchCmd)
}
