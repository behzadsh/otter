package scenarios

import "github.com/behzadsh/otter/internal/assertions"

type Step struct {
	Name    string               `yaml:"name"`
	URL     string               `yaml:"url"`
	Path    string               `yaml:"path"`
	Method  Method               `yaml:"method"`
	Headers map[string][]string  `yaml:"headers"`
	Query   any                  `yaml:"query"`
	Body    any                  `yaml:"body"`
	Assert  assertions.Assertion `yaml:"assert"`
}
