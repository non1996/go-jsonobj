package assert

func Should(conds ...func() error) error {
	for _, c := range conds {
		if err := c(); err != nil {
			return err
		}
	}
	return nil
}

type Asserter struct {
	conds []func() error
}

func NewAsserter() *Asserter {
	return &Asserter{}
}

func (a *Asserter) With(f func() error) *Asserter {
	a.conds = append(a.conds, f)
	return a
}

func (a *Asserter) WithIf(cond bool, action func() func() error) *Asserter {
	if cond {
		a.With(action())
	}
	return a
}

func (a *Asserter) WithIfElse(cond bool, action1, action2 func() func() error) *Asserter {
	if cond {
		a.With(action1())
	} else {
		a.With(action2())
	}
	return a
}

func (a *Asserter) Check() error {
	for _, c := range a.conds {
		if err := c(); err != nil {
			return err
		}
	}
	return nil
}
