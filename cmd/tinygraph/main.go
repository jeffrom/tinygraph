package main

import (
	"fmt"
	"os"

	"github.com/jeffrom/tinygraph"
	"github.com/spf13/cobra"
)

func workout() {
	graph := tinygraph.BlockGraph
	fmt.Println()
	fmt.Println()
	for i := 0; i < 10; i++ {
		if err := graph.Render(os.Stdout, i, 10); err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println()
	for i := 0; i < 100; i++ {
		if err := graph.Render(os.Stdout, i, 100); err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println()
	for i := 0; i < 63; i++ {
		if err := graph.Render(os.Stdout, i, 100); err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println()
	for i := 0; i < len(tinygraph.IntegralGraph); i++ {
		if err := tinygraph.IntegralGraph.Render(os.Stdout, i, len(tinygraph.IntegralGraph)); err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println()
	for i := 0; i < len(tinygraph.EqualSignGraph); i++ {
		if err := tinygraph.EqualSignGraph.Render(os.Stdout, i, len(tinygraph.EqualSignGraph)); err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println()
	graph = tinygraph.HorizontalBlockGraph
	for i := 0; i < 10; i++ {
		if err := graph.Render(os.Stdout, i, 10); err != nil {
			panic(err)
		}
		fmt.Println()
	}
}

type config struct {
	GraphName   string
	N           int
	Total       int
	Workout     bool
	CustomGraph []string
}

func newRootCmd() *cobra.Command {
	cfg := &config{}
	rootCmd := &cobra.Command{
		Use:   "tinygraph",
		Short: "print a single character graph",
		Long:  `tinygraph prints a graph using a single unicode character.`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return start(cfg, args)
		},
	}

	flags := rootCmd.Flags()
	flags.StringVarP(&cfg.GraphName, "graph", "c", "bar", "NAME of the graph (bar, horizbar, integral, equal)")
	flags.StringArrayVar(&cfg.CustomGraph, "custom", nil, "use your own custom graph")
	flags.IntVarP(&cfg.N, "amount", "n", 0, "amount to graph")
	flags.IntVarP(&cfg.Total, "total", "t", 0, "total to graph against amount")
	flags.BoolVar(&cfg.Workout, "workout", false, "show me the graphs")

	return rootCmd
}

func start(cfg *config, args []string) error {
	if cfg.Workout {
		workout()
		return nil
	}
	var graph tinygraph.Graph
	var err error
	if len(cfg.CustomGraph) > 0 {
		graph = tinygraph.Custom(cfg.CustomGraph)
	} else {
		graph, err = tinygraph.ByName(cfg.GraphName)
		if err != nil {
			return err
		}
	}
	return graph.Render(os.Stdout, cfg.N, cfg.Total)
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}