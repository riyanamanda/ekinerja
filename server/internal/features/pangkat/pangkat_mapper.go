package pangkat

func MapToListResponse(list []Pangkat) []PangkatResponse {
	responses := make([]PangkatResponse, len(list))
	for i, pangkat := range list {
		responses[i] = MapToPangkatResponse(pangkat)
	}
	return responses
}

func MapToPangkatResponse(pangkat Pangkat) PangkatResponse {
	return PangkatResponse(pangkat)
}
