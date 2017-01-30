package gowrapper

type Mapperfunc func(Wrap) Wrap

func Map(fn Mapperfunc, in chan Wrap) chan Wrap {
	out := make(chan Wrap)
	go func() {
		defer close(out)
		for ele := range in {
			out <- fn(ele)
		}
	}()

	return out
}
