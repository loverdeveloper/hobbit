package Models

import "encoding/json"

type HealthCheckModel struct {
	IsOk bool `json:"health"`
}

func (h *HealthCheckModel) CheckHealth() ([]byte, error) {
	// TODO
	h.IsOk = true
	return json.Marshal(h)
}
