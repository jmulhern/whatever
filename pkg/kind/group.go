package kind

type Group struct {
	Unique string  `json:"unique" yaml:"unique"`
	Name   string  `json:"name" yaml:"name"`
	Movies []Movie `json:"movies,omitempty" yaml:"movies,omitempty"`
}
