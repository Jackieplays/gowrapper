package gowrapper

func Combine(in ...chan Wrap) chan Wrap {
	out := make(chan Wrap)
	for inp := range in {
		defer close(in)
		for el := range inp {
			out <- el
		}
	}
	return out
}
