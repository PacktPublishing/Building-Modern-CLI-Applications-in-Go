package interfaces

type Command interface {
	ParseFlags([]string) error
	Run() error
	Name() string
}
