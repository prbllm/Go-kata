package task

import "fmt"

var registry = map[string]Runner{}

func Register(r Runner) {
	if _, exists := registry[r.Name()]; exists {
		panic(fmt.Sprintf("duplicate task name %q", r.Name()))
	}
	registry[r.Name()] = r
}

func Get(name string) (Runner, bool) {
	r, ok := registry[name]
	return r, ok
}

func All() []string {
	names := make([]string, 0, len(registry))
	for n := range registry {
		names = append(names, n)
	}
	return names
}
