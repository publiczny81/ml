package model

type Network struct {
	Features int       `json:"features"`
	Metrics  string    `json:"metrics"`
	Shape    []int     `json:"shape"`
	Topology int       `json:"topology"`
	Weights  []float64 `json:"weights"`
}
