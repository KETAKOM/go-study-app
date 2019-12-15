package user

type User struct {
	ID        int
	Name      string
	Address   string
	Office    Office
	TaxOffice TaxOffice
}

func (u User) GetUser(id int) (user User, err error) {
	return u, nil
}

func (u User) Validate() bool {
	// TODO: impl validate
	if !u.Office.Validate() || !u.TaxOffice.Validate() {
		return false
	}
	return true
}
