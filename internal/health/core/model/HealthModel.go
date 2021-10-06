package healthcoremodel

type HealthResponse struct {
	HealthStatuses []HealthStatus `json:"healthStatuses"`
}

type HealthStatus struct {
	System string `json:"system"`
	Status string `json:"status"`
}
