package user

type DTO struct {
	Id           int64  `json:"id"`
	UId          string `json:"uid"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	CreatedAt    string `json:"createdAt"`
}
