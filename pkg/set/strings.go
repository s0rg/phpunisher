package set

type stub struct{}

// Strings represents unordered set of string.
type Strings map[string]stub

// Add string to set.
func (s Strings) Add(v string) {
	s[v] = stub{}
}

// Has checks string in set.
func (s Strings) Has(v string) (ok bool) {
	_, ok = (s)[v]

	return
}

// FromList populates from list of strings.
func (s Strings) FromList(l []string) {
	for i := 0; i < len(l); i++ {
		s.Add(l[i])
	}
}
