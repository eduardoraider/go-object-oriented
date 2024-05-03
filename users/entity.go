package users

import "errors"

type user struct {
	Name     string
	Username string
	Secret   string
	Type     string
}

func (u user) GetName() string {
	return u.Name
}

func (u user) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Secret == "" {
		return errors.New("secret is required")
	}

	return nil
}
