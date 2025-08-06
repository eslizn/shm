package shm

import (
	"os"
	"path/filepath"
)

type Finder func(name string) string

var defaultFinder = func(name string) string {
	return filepath.Join(os.TempDir(), name)
}

type Options struct {
	name   string
	finder Finder
	force  bool
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

func WithForce(force bool) Option {
	return func(options *Options) {
		options.force = force
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
