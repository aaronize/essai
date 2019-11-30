package service

import "github.com/aaronize/essai/internal/models"

type Register struct {
	Account 	*models.Accounts
}

func (r *Register) Register() error {

	manager := models.NewManager()

	if _, err := manager.Add(r.Account); err != nil {

	}

	return nil
}
