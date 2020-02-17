package tinygraph

import (
	"fmt"
	"io"
	"unicode/utf8"
)

const singleSpace = ' '

var graphs = map[string]Graph{
	"bar":      BlockGraph,
	"horizbar": HorizontalBlockGraph,
	"integral": IntegralGraph,
	"equal":    EqualSignGraph,
}

func ByName(name string) (Graph, error) {
	graph, ok := graphs[name]
	if !ok {
		return nil, fmt.Errorf("graph not found: %q", name)
	}
	return graph, nil
}

func Custom(chars []string) Graph {
	var runes []rune
	for _, s := range chars {
		r, _ := utf8.DecodeRuneInString(s)
		runes = append(runes, r)
	}
	return Graph(runes)
}

type Graph []rune

func (c Graph) Render(w io.Writer, n, total int, prefix string, thresholds Thresholds) error {
	t := getThreshold(n, total, thresholds)
	if n == 0 || total == 0 {
		return write(w, c[0], prefix, t)
	}
	if n == total {
		return write(w, c[len(c)-1], prefix, t)
	}
	final := (len(c) * n) / total
	return write(w, c[final], prefix, t)
}

var BlockGraph = Graph{
	singleSpace,
	0x2581,
	0x2582,
	0x2583,
	0x2584,
	0x2585,
	0x2586,
	0x2587,
	0x2588,
}

var HorizontalBlockGraph = Graph{
	singleSpace,
	0x258F,
	0x258E,
	0x258D,
	0x258C,
	0x258B,
	0x258A,
	0x2589,
	0x2588,
}

var IntegralGraph = Graph{
	0x222B,
	0x222E,
	0x222C,
	0x222F,
	0x222D,
	0x2230,
}

var EqualSignGraph = Graph{
	' ',
	0x22C5,
	'-',
	0x223C,
	'=',
	0x2243,
	0x2248,
	0x2261,
	0x224C,
	0x224A,
	0x224B,
	0x2263,
}

func write(w io.Writer, r rune, prefix string, t *Threshold) error {
	b := make([]byte, utf8.UTFMax)
	utf8.EncodeRune(b, r)

	if t != nil {
		_, err := w.Write([]byte(t.Prefix))
		if err != nil {
			return err
		}
	}

	if prefix != "" {
		_, err := w.Write([]byte(prefix))
		if err != nil {
			return err
		}
	}

	_, err := w.Write(b)
	if err != nil {
		return err
	}

	if t != nil && t.Suffix != "" {
		_, err := w.Write([]byte(t.Suffix))
		if err != nil {
			return err
		}
	}
	return nil
}
