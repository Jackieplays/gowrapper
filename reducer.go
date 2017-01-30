package gowrapper

type ReducerFunc func(Wrap, Wrap) Wrap

func Reducer(fn ReducerFunc, in chan Wrap) chan Wrap {
	defer close(in)
	out := make(chan Wrap, len(in))
	var accum Wrap
	for ele := range in {
		if accum == nil {
			accum = ele
		} else {
			accum = fn(accum, ele)
		}
	}
	out <- accum
	return out
}
