package option

type Option interface {
	GetName() string
	GetCrossedOut() bool
	ToggleCrossedOut()
}
