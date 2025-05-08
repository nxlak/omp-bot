package user

import "fmt"

type Profile struct {
	ID    uint64
	Title string
}

func (p *Profile) String() string {
	return fmt.Sprintf("Entity: ID=%d, Title=%s", p.ID, p.Title)
}
