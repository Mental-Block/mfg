package domain

type RelationName string

// converts Relation back to its underlying type
func (s RelationName)String() string {
	return string(s)
}

func (s RelationName) IsValid() error {
	if (s.String() == "") {
		return nil
	}

	return nil 
}
