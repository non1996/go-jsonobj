package assert

func Should(conds ...func() error) error {
	for _, c := range conds {
		if err := c(); err != nil {
			return err
		}
	}
	return nil
}
