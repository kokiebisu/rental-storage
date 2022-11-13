package port

type ResponsePayload struct {
	ProviderId string `json:"providerId"`
	ProviderType string `json:"providerType"`
}

type Error struct {
	Message string `json:"message"`
}