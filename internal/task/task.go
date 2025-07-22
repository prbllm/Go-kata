package task

type Runner interface {
	Name() string
	Run() error
}
