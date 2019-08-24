package seq

type Seq struct{ xs []string }

func New(e string) Seq {
	return Seq{xs: []string{}}
}

func (s *Seq) ToSlice() []string {
	return s.xs
}
