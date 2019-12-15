package user

type TaxOffice struct {
	ID          int
	Name        string
	Address     string
	PhoneNumber string
}

func (t TaxOffice) GetTaxOffice() (to TaxOffice, err error) {
	return t, nil
}

func (t TaxOffice) Validate() bool {
	// TODO: impl validate
	return true
}
