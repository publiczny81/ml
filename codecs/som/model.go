package som

type Network struct {
	Features int       `json:"features"`
	Metrics  string    `json:"metrics"`
	Shape    []int     `json:"shape"`
	Topology string    `json:"topology"`
	Weights  []float64 `json:"weights"`
}
