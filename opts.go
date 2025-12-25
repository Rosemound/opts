package opts

type OptionContainer[T comparable] map[T]any

type Option[T comparable] func(o OptionContainer[T]) error

func CreateContainer[T comparable]() OptionContainer[T] {
	return OptionContainer[T]{}
}

func CreateContainerWithOptions[T comparable](o []Option[T]) (OptionContainer[T], error) {
	c := CreateContainer[T]()
	if err := c.ApplyA(o); err != nil {
		return nil, err
	}

	return c, nil
}

func CreateContainerWithOptionsS[T comparable](o []Option[T]) OptionContainer[T] {
	c := CreateContainer[T]()
	_ = c.ApplySilentA(o)
	return c
}

func (c OptionContainer[T]) Set(k T, v any) OptionContainer[T] {
	c[k] = v
	return c
}

func (c OptionContainer[T]) Exist(k T) bool {
	_, ok := c[k]
	return ok
}

func (c OptionContainer[T]) Get(k T) any {
	return c[k]
}

func (c OptionContainer[T]) Apply(opts ...Option[T]) error {
	return c.ApplyA(opts)
}

func (c OptionContainer[T]) ApplySilent(opts ...Option[T]) error {
	return c.ApplySilentA(opts)
}

func (c OptionContainer[T]) ApplyA(opts []Option[T]) error {
	var err error

	for _, opt := range opts {
		if err = opt(c); err != nil {
			return err
		}
	}

	return err
}

func (c OptionContainer[T]) ApplySilentA(opts []Option[T]) error {
	for _, opt := range opts {
		opt(c)
	}

	return nil
}
