package domain

type YAML struct {
	Name         string   `json:"name" yaml:"name"`
	Backend      string              		`json:"backend" yaml:"backend"`
	ResourceType string              		`json:"resource_type" yaml:"resource_type"`
	Actions      map[string][]string 		`json:"actions" yaml:"actions"`
}