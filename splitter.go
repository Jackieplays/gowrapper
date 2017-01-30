package gowrapper

func Split(filter FilterFunc, in chan Wrap) (chan Wrap, chan Wrap) {
	fail := make(chan Wrap)
	pass := make(chan Wrap)
	defer close(in)
	for ele := range in {
		if filter(ele) {
			pass <- ele
		} else {
			fail <- ele
		}
	}
	return fail, pass
}
