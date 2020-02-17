# tinygraph

Me: I really want a single character graph.

Me: Someone else has made this already?

...

Me: It'll take me less time to just make it myself.

## install

```sh
GO111MODULE=off go get github.com/jeffrom/tinygraph
```

## use

```
jeff@x1cats:~/.../github.com/jeffrom/tinygraph (master) $ tinygraph -h
tinygraph prints a graph using a single character.

Use one of the provided charts, or bring your own:

tinygraph --custom "1,2,3" -n 4 -t 10
# 2

You can also set custom thresholds to add 256 color codes, tmux status, or
arbitrary strings:

tinygraph --threshold '50:24' --threshold '60:35'

It's helpful to use --workout when trying thresholds.

Usage:
  tinygraph [flags]

Flags:
  -n, --amount int              amount to graph
      --custom strings          use your own custom graph
  -g, --graph string            NAME of the graph (bar, horizbar, integral, equal) (default "bar")
  -h, --help                    help for tinygraph
  -p, --prefix string           Prefix string for the graph
      --threshold stringArray   set a threshold with color code (default [60::,83::])
  -t, --total int               total to graph against amount
      --workout                 show me the graphs
```

Here are some example graphs:

```
jeff@x1cats:~/.../github.com/jeffrom/tinygraph (master) $ tinygraph --workout
thresholds: ["60:\x1b[1;33m:\x1b[0m" "83:\x1b[1;31m:\x1b[0m"]


  ▁▂▃▄▅▆▇█

            ▁▁▁▁▁▁▁▁▁▁▁▂▂▂▂▂▂▂▂▂▂▂▃▃▃▃▃▃▃▃▃▃▃▄▄▄▄▄▄▄▄▄▄▄▅▅▅▅▅▅▅▅▅▅▅▆▆▆▆▆▆▆▆▆▆▆▇▇▇▇▇▇▇▇▇▇▇███████████

            ▁▁▁▁▁▁▁▁▁▁▁▂▂▂▂▂▂▂▂▂▂▂▃▃▃▃▃▃▃▃▃▃▃▄▄▄▄▄▄▄▄▄▄▄▅▅▅▅▅▅▅

∫∮∬∯∭∰

 ⋅-∼=≃≈≡≌≊≋≣



▏
▎
▍
▌
▋
▊
▉
█
```
