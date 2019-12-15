package main

import (
	"fmt"

	"github.com/KETAKOM/go-study-app/application/domain/taxr/sof"
	"github.com/KETAKOM/go-study-app/application/domain/taxr/user"
)

func main() {

	i := user.Industry{
		ID:   1,
		Name: "IT・通信",
		Code: "123423431",
	}
	to := user.TaxOffice{
		ID:          5,
		Name:        "AA税理士事務所",
		Address:     "東京都渋谷区",
		PhoneNumber: "11112345",
	}
	o := user.Office{
		ID:       5,
		Name:     "株式会社サンプルAAAA",
		Address:  "東京都新宿区",
		Industry: i,
	}
	user := user.User{
		ID:        1,
		Name:      "サンプル 太郎",
		Address:   "東京都渋谷区",
		Office:    o,
		TaxOffice: to,
	}
	if !user.Validate() {
		fmt.Println("validate failed")
		return
	}

	fmt.Println("ok")

	sof := sof.Sof{
		10000000,
		10000000,
		1000000,
		1000000,
		600000,
		0,
	}
	s, err := sof.CalIncome()
	if err != nil {
		fmt.Println("CalIncome failed")
		return
	}

	fmt.Println(s.Income)
}
