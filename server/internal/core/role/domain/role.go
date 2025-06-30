package domain

type RoleName string

// converts Relation back to its underlying type
func (s RoleName)String() string {
	return string(s)
}
