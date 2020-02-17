package tinychart

import (
	"fmt"
	"io"
	"unicode/utf8"
)

const singleSpace = ' '

var charts = map[string]Chart{
	"bar":      BlockChart,
	"horizbar": HorizontalBlockChart,
	"integral": IntegralChart,
	"equal":    EqualSignChart,
}

func ByName(name string) (Chart, error) {
	chart, ok := charts[name]
	if !ok {
		return nil, fmt.Errorf("chart not found: %q", name)
	}
	return chart, nil
}

func Custom(chars []string) Chart {
	var runes []rune
	for _, s := range chars {
		r, _ := utf8.DecodeRuneInString(s)
		runes = append(runes, r)
	}
	return Chart(runes)
}

type Chart []rune

func (c Chart) Render(w io.Writer, n, total int) error {
	if n == 0 || total == 0 {
		return write(w, c[0])
	}
	final := (len(c) * n) / total
	return write(w, c[final])
}

var BlockChart = Chart{
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

var HorizontalBlockChart = Chart{
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

var IntegralChart = Chart{
	0x222B,
	0x222E,
	0x222C,
	0x222F,
	0x222D,
	0x2230,
}

var EqualSignChart = Chart{
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

func write(w io.Writer, r rune) error {
	b := make([]byte, utf8.UTFMax)
	utf8.EncodeRune(b, r)
	_, err := w.Write(b)
	if err != nil {
		return err
	}
	return nil
}