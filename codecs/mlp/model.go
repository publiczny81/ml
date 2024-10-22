package mlp

type LayerSpec struct {
	Activation string `json:"activation"`
	Neurons    int    `json:"neurons"`
}

type Network struct {
	Input   int         `json:"input"`
	Layers  []LayerSpec `json:"layers,omitempty"`
	Weights []float64   `json:"weights,omitempty"`
}
