package presenter

import (
	"github.com/thtrangphu/bookStoreSvc/entity"
)

// User data
type User struct {
	ID       entity.ID `json:"id"`
	Email    string    `json:"email"`
	FullName string    `json:"full_name"`
}
