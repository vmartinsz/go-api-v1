package entity

import "fmt"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      int8   `json:"age"`
}

func (u *User) Validate() error {
	if u.Age <= 0 {
		return fmt.Errorf("idade deve ser maior do que 0")
	}

	if u.LastName == "" {
		return fmt.Errorf("sobrenome vazio")
	}

	if u.Name == "" {
		return fmt.Errorf("nome vazio")
	}

	return nil
}
