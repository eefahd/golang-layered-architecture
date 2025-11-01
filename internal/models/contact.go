package models

type Contact struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (c *Contact) FullName() string {
	return c.FirstName + " " + c.LastName
}
