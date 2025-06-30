package env

const (
	Development ENVIROMENT = iota
	Production
	Test
)

var Enviroment = map[ENVIROMENT]string{
	Development: "development",
	Production:  "production",
	Test:        "test",
}

type ENVIROMENT int