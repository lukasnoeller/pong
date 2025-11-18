package option

type Button struct {
	Name       string
	Color      string
	CrossedOut bool
}

func (b Button) GetName() string {
	return b.Name
}
func (b Button) GetCrossedOut() bool {
	return b.CrossedOut
}
func (b *Button) ToggleCrossedOut() {
	b.CrossedOut = !b.CrossedOut
}
