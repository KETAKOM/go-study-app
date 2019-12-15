package user

type Industry struct {
	ID   int
	Name string
	Code string
}

func (i Industry) GetIndustry() (in Industry, err error) {
	return i, nil
}

func (i Industry) Validate() bool {
	// TODO: impl validate
	return true
}
