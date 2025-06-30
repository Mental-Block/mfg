package domain

type State bool

func (s State) Bool() bool {
	return bool(s)
}

const (
	Enabled  State = true
	Disabled State = false
)
