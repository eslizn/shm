package shm

type Options struct {
	name string
	dir  string
}

type Option func(*Options)

func WithName(name string) Option {
	return func(options *Options) {
		options.name = name
	}
}

func WithDir(dir string) Option {
	return func(options *Options) {
		options.dir = dir
	}
}

func getOptions(opts ...Option) *Options {
	options := &Options{
		name: "",
		dir:  DefaultDir,
	}
	for _, fn := range opts {
		fn(options)
	}
	return options
}
