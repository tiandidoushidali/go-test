package trace

var _ Tracer = nooptracer{}

type nooptracer struct {}

func (n nooptracer) New(title string, opts ...Option) Trace {
	return noopspan{}
}

func (n nooptracer) Inject(t Trace, format interface{}, carrier interface{}) error {
	return nil
}

func (n nooptracer) Extract(format interface{}, carrier interface{}) (Trace, error) {
	return noopspan{}, nil
}

type noopspan struct {}

func (n noopspan) TraceID() string { return "" }

func (n noopspan) Fork(string, string) Trace {
	return noopspan{}
}

func (n noopspan) Follow(string, string) Trace {
	return noopspan{}
}

func (n noopspan) Finish(err *error) {

}

func (n noopspan) SetTag(tags ...Tag) Trace {
	return noopspan{}
}

func (n noopspan) SetLog(logs ...LogFiled) Trace {
	return noopspan{}
}