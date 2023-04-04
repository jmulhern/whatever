package kind

type Movie struct {
	Unique   string `json:"unique" yaml:"unique"`
	Title    string `json:"title" yaml:"title"`
	Released string `json:"released" yaml:"released"`
}
