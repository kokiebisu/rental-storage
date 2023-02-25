package coordinate

type ValueObject struct {
	Latitude  float64
	Longitude float64
}

type DTO struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Raw struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (d DTO) ToValueObject() ValueObject {
	return ValueObject(d)
}

func (o ValueObject) ToDTO() DTO {
	return DTO(o)
}

func (r Raw) ToValueObject() ValueObject {
	return ValueObject(r)
}
