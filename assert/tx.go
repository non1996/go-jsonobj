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

func (a *Asserter) WithIf(cond bool, action func(a *Asserter)) *Asserter {
	if cond {
		action(a)
	}
	return a
}

func (a *Asserter) WithIfElse(cond bool, action1 func(a *Asserter), action2 func(a *Asserter)) *Asserter {
	if cond {
		action1(a)
	} else {
		action2(a)
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
