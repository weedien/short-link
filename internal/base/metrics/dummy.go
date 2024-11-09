package metrics

type Client interface {
	Inc(key string, value int)
}

type NoOp struct{}

func (d NoOp) Inc(_ string, _ int) {
	// todo - add some implementation!
}
