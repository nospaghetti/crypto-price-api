package v1

type HealthRequest struct {
}

type HealthResponse struct {
	Status string `json:"status"`
}
