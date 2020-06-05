package common

type (
	Daemon interface {
		Start() error
		Stop() error
	}
)
