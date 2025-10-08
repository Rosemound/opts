package mopts

type MoptionContainer[T comparable] map[T]any

type Moption[T comparable] func(o MoptionContainer[T]) error

func (c MoptionContainer[T]) Set(k T, v any) {
	c[k] = v
}

func (c MoptionContainer[T]) Get(k T) any {
	return c[k]
}

func (c MoptionContainer[T]) Apply(opts ...Moption[T]) error {
	return c.ApplyA(opts)
}

func (c MoptionContainer[T]) ApplySilent(opts ...Moption[T]) error {
	return c.ApplySilentA(opts)
}

func (c MoptionContainer[T]) ApplyA(opts []Moption[T]) error {
	var err error

	for _, opt := range opts {
		if err = opt(c); err != nil {
			return err
		}
	}

	return err
}

func (c MoptionContainer[T]) ApplySilentA(opts []Moption[T]) error {
	for _, opt := range opts {
		opt(c)
	}

	return nil
}
