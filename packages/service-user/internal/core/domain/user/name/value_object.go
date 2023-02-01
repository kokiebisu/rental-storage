package name

type NameType string

const (
	FirstName NameType = "firstName"
	LastName  NameType = "lastName"
)

type ValueObject struct {
	Value    string
	NameType NameType
}
