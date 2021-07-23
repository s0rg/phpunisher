package main

type score struct {
	Scanner string
	Score   float64
}

type scores []*score

func (s scores) Len() int           { return len(s) }
func (s scores) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s scores) Less(i, j int) bool { return s[i].Score < s[j].Score }

func (s scores) Sum() (rv float64) {
	for i := 0; i < len(s); i++ {
		rv += s[i].Score
	}

	return
}
