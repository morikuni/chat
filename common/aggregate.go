package common

type Aggregate struct {
	Version uint64
}

func (a *Aggregate) Updated() {
	a.Version++
}
