package main

import (
	"fmt"
	"os"

	"github.com/jeffrom/tinychart"
	"github.com/spf13/cobra"
)

func workout() {
	chart := tinychart.BlockChart
	fmt.Println()
	fmt.Println()
	for i := 0; i < 10; i++ {
		if err := chart.Render(os.Stdout, i, 10); err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println()
	for i := 0; i < 100; i++ {
		if err := chart.Render(os.Stdout, i, 100); err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println()
	for i := 0; i < 63; i++ {
		if err := chart.Render(os.Stdout, i, 100); err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println()
	for i := 0; i < len(tinychart.IntegralChart); i++ {
		if err := tinychart.IntegralChart.Render(os.Stdout, i, len(tinychart.IntegralChart)); err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println()
	for i := 0; i < len(tinychart.EqualSignChart); i++ {
		if err := tinychart.EqualSignChart.Render(os.Stdout, i, len(tinychart.EqualSignChart)); err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println()
	chart = tinychart.HorizontalBlockChart
	for i := 0; i < 10; i++ {
		if err := chart.Render(os.Stdout, i, 10); err != nil {
			panic(err)
		}
		fmt.Println()
	}
}

type config struct {
	ChartName   string
	N           int
	Total       int
	Workout     bool
	CustomChart []string
}

func newRootCmd() *cobra.Command {
	cfg := &config{}
	rootCmd := &cobra.Command{
		Use:   "tinychart",
		Short: "print a single character chart",
		Long:  `tinychart prints a chart using a single unicode character.`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return start(cfg, args)
		},
	}

	flags := rootCmd.Flags()
	flags.StringVarP(&cfg.ChartName, "chart", "c", "bar", "NAME of the chart (bar, horizbar, integral, equal)")
	flags.StringArrayVar(&cfg.CustomChart, "custom", nil, "use your own custom chart")
	flags.IntVarP(&cfg.N, "amount", "n", 0, "amount to chart")
	flags.IntVarP(&cfg.Total, "total", "t", 0, "total to chart against amount")
	flags.BoolVar(&cfg.Workout, "workout", false, "show me the graphs")

	return rootCmd
}

func start(cfg *config, args []string) error {
	if cfg.Workout {
		workout()
		return nil
	}
	var chart tinychart.Chart
	var err error
	if len(cfg.CustomChart) > 0 {
		chart = tinychart.Custom(cfg.CustomChart)
	} else {
		chart, err = tinychart.ByName(cfg.ChartName)
		if err != nil {
			return err
		}
	}
	return chart.Render(os.Stdout, cfg.N, cfg.Total)
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
