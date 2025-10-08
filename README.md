<p align="center">
  <h1 align="center">OPTS</h1>
  <p align="center">Too simple, zero-dependencies library to flexible passing dynamic options</p>
</p>

<div align="center">

  [![Test & Build](https://github.com/rosemound/opts/actions/workflows/build.yml/badge.svg)](https://github.com/rosemound/opts/actions/workflows/build.yml)
  [![Go Reference](https://pkg.go.dev/badge/github.com/rosemound/opts.svg)](https://pkg.go.dev/github.com/rosemound/opts)
  [![Go Report Card](https://goreportcard.com/badge/github.com/rosemound/opts)](https://goreportcard.com/report/github.com/rosemound/opts)
  
</div>

### About
Too simple, zero-dependencies library to flexible passing of dynamic params into variadic functions.

### Installation
```bash
go get github.com/rosemound/opts
```

### Usage mapped opts
```go
import (
	"context"
	"errors"

	"github.com/rosemound/opts/mopts"
)

// Mapped option key type
type MoptionKey string

// Owner key
const OwnerKey MoptionKey = "owner"

// Company key (like ctx keys)
const CompanyKey MoptionKey = "company"

// simple service
type service struct {
	owner   string
	company any
}

// main
func main() {

	// service instantination with dynamic params
	s, err := NewService(context.Background(), WithCompany("test"), WithOwner("test"))

	// err handling
	if err != nil {
		panic(err)
	}

	// do staff
}

func NewService(ctx context.Context, opts ...mopts.Moption[MoptionKey]) (*service, error){

	// init moption container
	conf := mopts.MoptionContainer[MoptionKey]{}
	
	// apply opts
	if err := conf.ApplyA(opts); err != nil {
		return nil, err
	}

	// or apply silently
	conf.ApplySilentA(opts)

	// allocate your instance
	return &service{
		owner: conf.Get(OwnerKey).(string),
		company: conf.Get(CompanyKey),
	}, nil
}

// With company
func WithCompany(company any) mopts.Moption[MoptionKey] {
	return func(o mopts.MoptionContainer[MoptionKey]) error {
		o.Set(CompanyKey, company)
		return nil
	}
}

// With owner & err
func WithOwner(val string) mopts.Moption[MoptionKey] {
	return func(o mopts.MoptionContainer[MoptionKey]) error {
		if val == "" {
			return errors.New("owner must be present")
		}

		o.Set(OwnerKey, val)
		return nil
	}
}
```

### Usage opts
```go
package main

import (
	"context"
	"errors"

	"github.com/rosemound/opts/opts"
)

// Option key type
type OptionKey uint8

// Owner key
const OwnerKey OptionKey = 0

// Company key
const CompanyKey OptionKey = 1

// simple service
type service struct {
	owner   string
	company any
}

// main
func main() {

	// service instantination with dynamic params
	s, err := NewService(context.Background(), WithCompany("test"), WithOwner("test"))

	// err handling
	if err != nil {
		panic(err)
	}

	// do staff
}

func NewService(ctx context.Context, options ...opts.Option) (*service, error) {

	// init moption container
	conf := opts.OptionContainer{}

	// apply opts
	if err := conf.ApplyA(options); err != nil {
		return nil, err
	}

	// or apply silently
	conf.ApplySilentA(options)

	// allocate your instance
	return &service{
		owner:   conf.Get(OwnerKey).(string),
		company: conf.Get(CompanyKey),
	}, nil
}

// With company
func WithCompany(company any) opts.Option {
	return func(o opts.OptionContainer) error {
		o.Set(CompanyKey, company)
		return nil
	}
}

// With owner & err
func WithOwner(val string) opts.Option {
	return func(o opts.OptionContainer) error {
		if val == "" {
			return errors.New("owner must be present")
		}

		o.Set(OwnerKey, val)
		return nil
	}
}

```