/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"

	"github.com/jtgs/aoc24-go/days/day01"
	"github.com/jtgs/aoc24-go/days/day02"
	"github.com/jtgs/aoc24-go/days/day03"
	"github.com/jtgs/aoc24-go/days/day04"
	"github.com/jtgs/aoc24-go/days/day05"
	"github.com/jtgs/aoc24-go/days/day06"
	"github.com/jtgs/aoc24-go/days/day07"
	"github.com/spf13/cobra"
)

type runfunc func() (int, int)

var Day int

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve an AoC puzzle",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		days := []runfunc{
			nil,
			day01.Run,
			day02.Run,
			day03.Run,
			day04.Run,
			day05.Run,
			day06.Run,
			day07.Run,
		}

		if Day > 0 {
			if Day >= len(days) {
				slog.Error("Error: not a valid day")
				return
			}

			chosenDay := days[Day]

			if chosenDay == nil {
				slog.Error("Error: no solution yet")
				return
			}

			part1, part2 := chosenDay()
			fmt.Printf("==== Day %d ====\n", Day)
			fmt.Printf("Part 1: %d\n", part1)
			fmt.Printf("Part 2: %d\n", part2)
		} else {
			// run all
			for ix, day := range days {
				if day != nil {
					part1, part2 := day()
					fmt.Printf("==== Day %d ====\n", ix)
					fmt.Printf("Part 1: %d\n", part1)
					fmt.Printf("Part 2: %d\n\n", part2)
				}
			}
		}

	},
}

func init() {
	solveCmd.Flags().IntVarP(&Day, "day", "d", 0, "Day to run")
	rootCmd.AddCommand(solveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// solveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// solveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
