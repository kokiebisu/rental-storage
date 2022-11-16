package port

import domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"

type ItemRepository interface {
	Save(domain.Item) error
	Delete(uid string) error
	FindOneById(uid string) (domain.Item, error)
}

type UserRepository interface {
	Setup() error
	Save(domain.User) (string, error)
	Delete(uid string) error
	FindOneById(uid string) (domain.User, error)
	FindOneByEmail(emailAddress string) (domain.User, error)
}
