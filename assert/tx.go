package assert

import "time"

type Assertion struct {
	conditions []func() error
}

func (a *Assertion) add(cond func() error) *Assertion {
	a.conditions = append(a.conditions, cond)
	return a
}

func (a *Assertion) True(cond bool, errMsg ...string) *Assertion {
	return a.add(func() error { return True(cond, errMsg...) })
}

func (a *Assertion) False(cond bool, errMsg ...string) *Assertion {
	return a.add(func() error { return False(cond, errMsg...) })
}

func (a *Assertion) TimeValid(t time.Time, errMsg ...string) *Assertion {
	return a.add(func() error { return TimeValid(t, errMsg...) })
}

func (a *Assertion) IntEQ(expect, actual int64, errMsg ...string) *Assertion {
	return a.add(func() error { return Int64EQ(expect, actual, errMsg...) })
}

func (a *Assertion) IntNE(expect, actual int64, errMsg ...string) *Assertion {
	return a.add(func() error { return Int64NE(expect, actual, errMsg...) })
}

func (a *Assertion) IntGTE(expect, actual int64, errMsg ...string) *Assertion {
	return a.add(func() error { return Int64GTE(expect, actual, errMsg...) })
}

func (a *Assertion) IntGT(expect, actual int64, errMsg ...string) *Assertion {
	return a.add(func() error { return Int64GT(expect, actual, errMsg...) })
}

func (a *Assertion) IntLTE(expect, actual int64, errMsg ...string) *Assertion {
	return a.add(func() error { return Int64LTE(expect, actual, errMsg...) })
}

func (a *Assertion) IntIn(expect []int64, actual int64, errMsg ...string) *Assertion {
	return a.add(func() error { return Int64In(expect, actual, errMsg...) })
}

func (a *Assertion) IntLT(expect, actual int64, errMsg ...string) *Assertion {
	return a.add(func() error { return Int64LT(expect, actual, errMsg...) })
}

func (a *Assertion) StringNotEmpty(s string, errMsg ...string) *Assertion {
	return a.add(func() error { return StringNotEmpty(s, errMsg...) })
}

func (a *Assertion) StringNoLonger(maxLen int, s string, errMsg ...string) *Assertion {
	return a.add(func() error { return StringNoLonger(maxLen, s, errMsg...) })
}

func (a *Assertion) StringEQ(expect, actual string, errMsg ...string) *Assertion {
	return a.add(func() error { return StringEqual(expect, actual, errMsg...) })
}

func (a *Assertion) StringIn(expect []string, actual string, errMsg ...string) *Assertion {
	return a.add(func() error { return StringIn(expect, actual, errMsg...) })
}

func (a *Assertion) Check() error {
	if len(a.conditions) == 0 {
		return nil
	}
	for _, f := range a.conditions {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

func All() *Assertion {
	return &Assertion{}
}
