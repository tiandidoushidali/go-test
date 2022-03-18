package trace

var (
	// global tracer
	_tracer Tracer = nooptracer{}
)

type Tracer interface {
	New(operationName string, opts ...Option) Trace

	Inject(t Trace, format interface{}, carrier interface{}) error

	Extract(format interface{}, carrier interface{}) (Trace, error)
}

// New trace instance with given operationName
func New(operationName string, opts ...Option) Trace {
	return _tracer.New(operationName, opts...)
}

type Trace interface {
	TraceID() string
	// Fork fork a trace with client trace
	Fork(serviceName, operationName string) Trace

	Follow(serviceName, operationName string) Trace

	// Finish when trace finish call it.
	Finish(err *error)

	SetTag(tags ...Tag) Trace

	SetLog(logs ...LogFiled) Trace
}