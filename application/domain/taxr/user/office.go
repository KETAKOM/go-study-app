package user

type Office struct {
	ID       int
	Name     string
	Address  string
	Industry Industry
}

func (o Office) GetOffice() (of Office, err error) {
	return o, nil
}

func (t Office) Validate() bool {
	// TODO: impl validate
	if !t.Industry.Validate() {
		return false
	}
	return true
}
