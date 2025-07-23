package shm

type Finder func(name string) string

type Options struct {
	name   string
	finder Finder
}

type Option func(*Options)

func WithName(name string) Option {
	return func(options *Options) {
		options.name = name
	}
}

func WithFinder(finder Finder) Option {
	return func(options *Options) {
		options.finder = finder
	}
}

func getOptions(opts ...Option) *Options {
	options := &Options{
		name:   "",
		finder: defaultFinder,
	}
	for _, fn := range opts {
		fn(options)
	}
	return options
}
