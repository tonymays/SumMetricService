package root

// ---- Metric ----
type Metric struct {
	Id			string	`json:"id,omitempty"`
	Key			string	`json:"key,omitempty"`
	Value		int		`json:"value,omitempty"`
	EntryTime	string	`json:"entry_time,omitempty"`
}

// ---- MetricSum ----
type MetricSum struct {
	Key			string	`json:"key"`
	Sum			int		`json:"sum"`
}

// ---- Metric Interface ----
type MetricService interface {
	AddMetric(m Metric) (Metric, error)
	GetMetrics() ([]Metric, error)
	ShowActiveMetrics() ([]Metric, error)
	SumMetrics(key string) (MetricSum, error)
	ClearOutdatedMetrics(key string) error
	CanCount(m Metric) bool
}