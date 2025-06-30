package domain

import "errors"

type URN string

func (u URN) String() string {
	return string(u)
}

// func BuildURN(projectName string) {
// 	return URN(fmt.Sprintf("frn:%s:%s:%s", projectName, res.NamespaceID, res.Name))
// }

func (u URN) IsValid() error {
	if (u.String() != "") {
		return errors.New("dsadsadas")
	}

	return nil
}
