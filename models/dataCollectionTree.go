package models

//api_repsonse
type ApiResponse struct {
	Msg  string `json:"msg,omitempty"`
	Data *Data  `json:"data,omitempty"`
}

// Dimension struct
type Dimension struct {
	Key   string `json:"key"`
	Value string `json:"val"`
}

// Metric struct
type Metric struct {
	Key   string `json:"key"`
	Value int    `json:"val"`
}

// Data struct
type Data struct {
	Dimensions []Dimension `json:"dim"`
	Metrics    []Metric    `json:"metrics"`
}
