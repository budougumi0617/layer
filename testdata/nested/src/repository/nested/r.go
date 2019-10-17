package repository

import (
	dn "domain/nested"
)

func UserRepository() *Repo {
	return &Repo{}
}

type Repo struct {
}

func (u *Repo) Get(id int) *dn.User {
	return nil
}
