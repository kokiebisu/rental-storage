package location

import "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/coordinate"

type ValueObject struct {
	Address      string
	City         string
	Country      string
	CountryCode  string
	Phone        string
	Province     string
	ProvinceCode string
	Zip          string
	Coordinate   coordinate.ValueObject
}

type DTO struct {
	Address      string         `json:"address"`
	City         string         `json:"city"`
	Country      string         `json:"country"`
	CountryCode  string         `json:"countryCode"`
	Phone        string         `json:"phone"`
	Province     string         `json:"province"`
	ProvinceCode string         `json:"provinceCode"`
	Zip          string         `json:"zip"`
	Coordinate   coordinate.DTO `json:"coordinate"`
}

type Raw struct {
	Address      string         `json:"address"`
	City         string         `json:"city"`
	Country      string         `json:"country"`
	CountryCode  string         `json:"country_code"`
	Phone        string         `json:"phone"`
	Province     string         `json:"province"`
	ProvinceCode string         `json:"province_code"`
	Zip          string         `json:"zip"`
	Coordinate   coordinate.Raw `json:"coordinate"`
}

func (d DTO) ToValueObject() ValueObject {
	return ValueObject{
		Address:      d.Address,
		City:         d.City,
		Country:      d.Country,
		CountryCode:  d.CountryCode,
		Phone:        d.Phone,
		Province:     d.Province,
		ProvinceCode: d.ProvinceCode,
		Zip:          d.Zip,
		Coordinate:   d.Coordinate.ToValueObject(),
	}
}

func (o ValueObject) ToDTO() DTO {
	return DTO{
		Address:      o.Address,
		City:         o.City,
		Country:      o.Country,
		CountryCode:  o.CountryCode,
		Phone:        o.Phone,
		Province:     o.Province,
		ProvinceCode: o.ProvinceCode,
		Zip:          o.Zip,
		Coordinate:   o.Coordinate.ToDTO(),
	}
}

func (r Raw) ToValueObject() ValueObject {
	return ValueObject{
		Address:      r.Address,
		City:         r.City,
		Country:      r.Country,
		CountryCode:  r.CountryCode,
		Phone:        r.Phone,
		Province:     r.Province,
		ProvinceCode: r.ProvinceCode,
		Zip:          r.Zip,
		Coordinate:   r.Coordinate.ToValueObject(),
	}
}
