package sof

type Sof struct {
	Sales  int // 売上金額
	SalD   int // 売上差引金額(SalesDeduction)
	Exp    int // 経費(Expense)
	ExpD   int // 経費差引金額
	Bdsd   int // 青色深刻特別控除額(Blue deduction special deduction)
	Income int // 所得金額
}

func (s Sof) CalIncome() (sof Sof, err error) {
	s.Income = s.SalD - s.ExpD - s.Bdsd
	return s, nil
}
