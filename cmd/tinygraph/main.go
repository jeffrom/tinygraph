package main

import (
	"fmt"
	"os"

	"github.com/jeffrom/tinygraph"
	"github.com/spf13/cobra"
)

var defaultThresholds []string

func init() {
	if os.Getenv("TMUX_STATUS") != "" {
		defaultThresholds = []string{"60:#[fg=yellow,bold]", "83:#[fg=red,bold]"}
	} else {
		defaultThresholds = []string{"60:\033[1;33m:\033[0m", "83:\033[1;31m:\033[0m"}
	}
}

func workout(cfg *config) {
	fmt.Printf("thresholds: %q\n", cfg.Thresholds)
	ts, err := tinygraph.NewThresholds(cfg.Thresholds...)
	if err != nil {
		panic(err)
	}

	graph := tinygraph.BlockGraph
	fmt.Println()
	fmt.Println()
	for i := 0; i < 10; i++ {
		if err := graph.Render(os.Stdout, i, 10, "", ts); err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println()
	for i := 0; i < 100; i++ {
		if err := graph.Render(os.Stdout, i, 100, "", ts); err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println()
	for i := 0; i < 63; i++ {
		if err := graph.Render(os.Stdout, i, 100, "", ts); err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println()
	for i := 0; i < len(tinygraph.IntegralGraph); i++ {
		if err := tinygraph.IntegralGraph.Render(os.Stdout, i, len(tinygraph.IntegralGraph), "", ts); err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println()
	for i := 0; i < len(tinygraph.EqualSignGraph); i++ {
		if err := tinygraph.EqualSignGraph.Render(os.Stdout, i, len(tinygraph.EqualSignGraph), "", ts); err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println()
	graph = tinygraph.HorizontalBlockGraph
	for i := 0; i < 10; i++ {
		if err := graph.Render(os.Stdout, i, 10, "", ts); err != nil {
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
	Thresholds  []string
	Prefix      string
}

func newRootCmd() *cobra.Command {
	cfg := &config{}
	rootCmd := &cobra.Command{
		Use:   "tinygraph",
		Short: "print a graph",
		Long: `tinygraph prints a graph using a single character.

Use one of the provided charts, or bring your own:

tinygraph --custom "1,2,3" -n 4 -t 10
# 2
`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return start(cfg, args)
		},
	}

	flags := rootCmd.Flags()
	flags.StringVarP(&cfg.GraphName, "graph", "g", "bar", "NAME of the graph (bar, horizbar, integral, equal)")
	flags.StringVarP(&cfg.Prefix, "prefix", "p", "", "Prefix string for the graph")
	flags.StringSliceVar(&cfg.CustomGraph, "custom", nil, "use your own custom graph")
	flags.StringArrayVar(&cfg.Thresholds, "threshold", defaultThresholds, "set a threshold with color code")
	flags.IntVarP(&cfg.N, "amount", "n", 0, "amount to graph")
	flags.IntVarP(&cfg.Total, "total", "t", 0, "total to graph against amount")
	flags.BoolVar(&cfg.Workout, "workout", false, "show me the graphs")

	return rootCmd
}

func start(cfg *config, args []string) error {
	if cfg.Workout {
		workout(cfg)
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
	thresholds, err := tinygraph.NewThresholds(cfg.Thresholds...)
	if err != nil {
		return err
	}
	return graph.Render(os.Stdout, cfg.N, cfg.Total, cfg.Prefix, thresholds)
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
