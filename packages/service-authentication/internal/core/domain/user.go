package domain

type User struct {
	Id           int8   `json:"id"`
	Uid          string `json:"uid"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	CreatedAt    string `json:"createdAt"`
}
