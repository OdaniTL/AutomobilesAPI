package models

type Models struct {
	ModelId     int    `json:"category_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Decriptions string `json:"decriptions,omitempty"`
	State       bool   `json:"state,omitempty"`
}

type Autos struct {
	AutoID       int     `json:"auto_id,omitempty"`
	Brand        string  `json:"name,omitempty"`
	Model        string  `json:"models,omitempty"`
	Year         int     `json:"year,omitempty"`
	BodyType     string  `json:"body_type,omitempty"`
	EngineType   string  `json:"engine_type,omitempty"`
	Displacement int     `json:"displacement,omitempty"`
	Price        float32 `json:"price,omitempty"`
	State        bool    `json:"state,omitempty"`
}
