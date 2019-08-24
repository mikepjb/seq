package seq

type Seq struct{ xs []string }

func New(e string) Seq {
	return Seq{xs: []string{e}}
}

func (s Seq) Map(f func(string) string) Seq {
	for i, x := range s.xs {
		s.xs[i] = f(x)
	}
	return s
}

func (s Seq) ToSlice() []string {
	return s.xs
}
