package repository

import (
	"domain"
)

func UserRepository() *Repo {
	return &Repo{}
}

type Repo struct {
}

func (u *Repo) Get(id int) *domain.User {
	return nil
}
