package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/raithbheart/whack-a-goph/internal"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch [exercise_name]",
	Short: "Watch mode - automatically runs tests when files change",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var exercise *internal.Exercise
		var err error

		if len(args) == 1 {
			// Watch specific exercise
			exercise, err = internal.FindExercise(args[0])
			if err != nil {
				fmt.Printf("%s %v\n", red("✗"), err)
				return
			}
		} else {
			// Watch next exercise
			exercise, err = internal.GetNextExercise()
			if err != nil {
				fmt.Printf("%s %v\n", green("🎉"), err)
				return
			}
		}

		fmt.Printf("%s\n", cyan(gopher))
		fmt.Printf("%s Watching exercise: %s\n", yellow("👁️"), bold(exercise.Name))
		fmt.Printf("Description: %s\n", exercise.Description)
		fmt.Printf("File: %s\n", cyan(exercise.Path))
		fmt.Printf("\nPress Ctrl+C to stop watching...\n")
		fmt.Println("─────────────────────────────────────────")

		// Set up signal handling
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		// Watch for file changes
		lastModTime := getFileModTime(exercise.Path)

		// Run initial test
		runExerciseWatch(*exercise)

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-c:
				fmt.Printf("\n%s Stopping watch mode...\n", yellow("👋"))
				return
			case <-ticker.C:
				currentModTime := getFileModTime(exercise.Path)
				if !currentModTime.Equal(lastModTime) {
					lastModTime = currentModTime
					fmt.Printf("\n%s File changed, running tests...\n", yellow("🔄"))
					fmt.Println("─────────────────────────────────────────")
					runExerciseWatch(*exercise)
				}
			}
		}
	},
}

func runExerciseWatch(exercise internal.Exercise) {
	passed, output, err := internal.RunExercise(exercise)
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

		fmt.Printf("\nGreat job! The gopher has been whacked! 🎯\n")
		fmt.Printf("Run %s to move to the next exercise.\n", green("whack-a-goph next"))
	} else {
		fmt.Printf("%s\n", red(gopherFail))
		fmt.Printf("%s Exercise failed!\n", red("✗"))
		fmt.Printf("\nTest output:\n%s\n", output)

		if exercise.Hint != "" {
			fmt.Printf("%s Hint: %s\n", yellow("💡"), yellow(exercise.Hint))
		}

		fmt.Printf("\nKeep working on it! The file is being watched for changes...\n")
	}
	fmt.Println("─────────────────────────────────────────")
}

func getFileModTime(filepath string) time.Time {
	info, err := os.Stat(filepath)
	if err != nil {
		return time.Time{}
	}
	return info.ModTime()
}
