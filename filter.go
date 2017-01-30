package gowrapper

type FilterFunc func(Wrap) Wrap

func Filter(fn FilterFunc, in chan Wrap) chan Wrap {
	defer close(in)
	out := make(chan Wrap)
	for ele := range in {
		if fn(ele) {
			out <- ele
		}
	}
	return out
}
