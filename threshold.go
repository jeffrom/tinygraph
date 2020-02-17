package tinygraph

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var isTmuxStatus = os.Getenv("TMUX_STATUS") != ""

var thresholdRe = regexp.MustCompile(`^([0-9]+):([^:]+):?(.+)?$`)

type Thresholds []*Threshold

type Threshold struct {
	N      int
	Prefix string
	Suffix string
	parsed bool
	isANSI bool
}

func (t *Threshold) For(s, prefix string) string {
	if t == nil {
		return prefix + s
	}
	if isTmuxStatus {
		return fmt.Sprintf("#[%s]%s%s", t.Prefix, prefix, s)
	}

	t.parse()
	if t.isANSI {
		return fmt.Sprintf("\033[38;5;%sm%s%s\033[39;49m", t.Prefix, prefix, s)
	}

	return t.Prefix + s + t.Suffix
}

func (t *Threshold) parse() {
	if t.parsed {
		return
	}
	defer func() {
		t.parsed = true
	}()

	parts := strings.SplitN(t.Prefix, ",", 2)
	_, err := strconv.ParseInt(parts[0], 10, 16)
	if err != nil {
		// still want to pass through the string in this case
		return
	}
	t.isANSI = true
}

func NewThresholds(raw ...string) (Thresholds, error) {
	var ts Thresholds
	for _, s := range raw {
		t, err := newThreshold(s)
		if err != nil {
			return Thresholds{}, err
		}
		ts = append(ts, t)
	}
	return ts, nil
}

func newThreshold(s string) (*Threshold, error) {
	if !thresholdRe.MatchString(s) {
		return nil, errors.New(`threshold must match format: N:prefix[:suffix]`)
	}
	parts := thresholdRe.FindStringSubmatch(s)

	n, err := strconv.ParseInt(parts[1], 10, 16)
	if err != nil {
		return nil, err
	}

	t := &Threshold{
		N:      int(n),
		Prefix: parts[2],
	}
	if len(parts) > 3 {
		t.Suffix = parts[3]
	}
	return t, nil
}

func getThreshold(n, total int, thresholds Thresholds) *Threshold {
	outOf100 := (n * 100) / total
	var t *Threshold
	for _, possible := range thresholds {
		if outOf100 >= possible.N {
			t = possible
		}
	}
	return t
}
