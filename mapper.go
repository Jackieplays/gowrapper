package gowrapper

type Mapperfunc func(Wrap) Wrap

func Map(fn Mapperfunc, in chan Wrap) chan Wrap {
	defer close(in)
	out := make(chan Wrap)
	for ele := range in {
		out <- fn(ele)
	}
	return out
}
