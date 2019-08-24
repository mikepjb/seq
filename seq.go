package seq

type Seq struct{ xs []string }

func New(e string) Seq {
	return Seq{xs: []string{e}}
}

func (s Seq) Map(fi interface{}, args ...interface{}) Seq {
	switch f := fi.(type) {
	case func(string) string:
		for i, x := range s.xs {
			s.xs[i] = f(x)
		}
	case func(string, string, string, int) string:
		for i, x := range s.xs {
			s.xs[i] = f(x, args[0].(string), args[1].(string), args[2].(int))
		}
	default:
	}

	return s
}

func (s Seq) ToSlice() []string {
	return s.xs
}
