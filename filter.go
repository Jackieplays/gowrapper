package gowrapper

type FilterFunc func(Wrap) bool

func Filter(fn FilterFunc, in chan Wrap) chan Wrap {
	defer close(in)
	out := make(chan Wrap, len(in))
	for ele := range in {
		if fn(ele) {
			out <- ele
		}
	}
	return out
}
