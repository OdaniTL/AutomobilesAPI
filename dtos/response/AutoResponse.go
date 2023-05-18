package response

type AutosResponse struct {
	AutoID       int     `json:"auto_id,omitempty"`
	Brand        string  `json:"brand,omitempty"`
	Model        string  `json:"model,omitempty"`
	Year         int     `json:"year,omitempty"`
	BodyType     string  `json:"body_type,omitempty"`
	EngineType   string  `json:"engine_type,omitempty"`
	Displacement int     `json:"displacement,omitempty"`
	Price        float32 `json:"price,omitempty"`
	State        bool    `json:"state,omitempty"`
}
