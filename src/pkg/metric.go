package root

type Metric struct {
	Id		string	`json:"id,omitempty"`
	Key		string	`json:"key,omitempty"`
	Value	int		`json:"value,omitempty"`
}

type MetricService interface {
	
}