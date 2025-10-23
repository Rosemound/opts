package opts

type OptionContainer map[any]any

type Option func(c OptionContainer) error

func (c OptionContainer) Set(k any, v any) {
	c[k] = v
}

func (c OptionContainer) Exist(k any) bool {
	_, ok := c[k]
	return ok
}

func (c OptionContainer) Get(k any) any {
	return c[k]
}

func (c OptionContainer) Apply(opts ...Option) error {
	return c.ApplyA(opts)
}

func (c OptionContainer) ApplySilent(opts ...Option) error {
	return c.ApplySilentA(opts)
}

func (c OptionContainer) ApplyA(opts []Option) error {
	var err error

	for _, opt := range opts {
		if err = opt(c); err != nil {
			return err
		}
	}

	return err
}

func (c OptionContainer) ApplySilentA(opts []Option) error {
	for _, opt := range opts {
		opt(c)
	}

	return nil
}
