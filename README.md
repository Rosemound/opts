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
go get github.com/rosemound/opts/v2
```

### Usage mapped opts
```go
import (
	"context"
	"errors"

	"github.com/rosemound/opts/v2"
)

// Mapped option key type
type OptionKey string

// Owner key
const OwnerKey OptionKey = "owner"

// Company key (like ctx keys)
const CompanyKey OptionKey = "company"

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

func NewService(ctx context.Context, o ...opts.Option[OptionKey]) (*service, error){

	// init option container with err handling
	c, err := opts.CreateContainerWithOptions(o)

	if err != nil {
		return nil, err
	}

	// or silently
	c := opts.CreateContainerWithOptionsS(o)

	// allocate your instance
	return &service{
		owner: conf.Get(OwnerKey).(string),
		company: conf.Get(CompanyKey),
	}, nil
}

// With company
func WithCompany(company any) opts.Option[OptionKey] {
	return func(o opts.OptionContainer[OptionKey]) error {
		o.Set(CompanyKey, company)
		return nil
	}
}

// With owner & err
func WithOwner(val string) opts.Option[OptionKey] {
	var err error

	if val == "" {
		err = errors.New("owner must be present")
	}

	return func(o opts.OptionContainer[OptionKey]) error {
		if err == nil {
			o.Set(OwnerKey, val)
		}

		return err
	}
}
```