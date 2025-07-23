package shm

type Namer func(name string) string

type Options struct {
	name  string
	namer Namer
}

type Option func(*Options)

func WithName(name string) Option {
	return func(options *Options) {
		options.name = name
	}
}

func WithNamer(namer Namer) Option {
	return func(options *Options) {
		options.namer = namer
	}
}

func getOptions(opts ...Option) *Options {
	options := &Options{
		name:  "",
		namer: defaultNamer,
	}
	for _, fn := range opts {
		fn(options)
	}
	return options
}
