package domain

import "errors"

type ImageUrl struct {
	Value string
}

func NewImageUrl(url string) (ImageUrl, error) {
	err := isValidImageUrl(url)
	if err != nil {
		return ImageUrl{}, err
	}
	return ImageUrl{
		Value: url,
	}, nil
}

func isValidImageUrl(value string) error {
	if value == "" {
		return errors.New("image url cannot be empty")
	}
	return nil
}
