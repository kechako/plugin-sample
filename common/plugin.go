package common

type Plugin interface {
	Name() string
	Do() error
}
