package trace

type option struct {
	Debug bool
}

type Option func(*option)

