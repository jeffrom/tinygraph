package tinygraph

import (
	"errors"
	"regexp"
	"strconv"
)

var thresholdRe = regexp.MustCompile(`^([0-9]+):([^:]+):?(.+)?$`)

type Thresholds []*Threshold

type Threshold struct {
	N      int
	Prefix string
	Suffix string
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
