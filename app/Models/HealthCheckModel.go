package Models

type HealthCheckModel struct {
	IsOk bool `json:"health"`
}

func (h *HealthCheckModel) CheckHealth() {
	// TODO
	h.IsOk = true
}
