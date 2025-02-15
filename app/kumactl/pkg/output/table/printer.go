package table

import (
	"io"
)

type Printer interface {
	Print(Table, io.Writer) error
}

type Table struct {
	Headers []string
	NextRow func() []string
}

func NewPrinter() Printer {
	return &printer{}
}

var _ Printer = &printer{}

type printer struct{}

func (p *printer) Print(data Table, out io.Writer) error {
	table := NewWriter(out)
	defer table.Flush()

	if 0 < len(data.Headers) {
		if err := table.Headers(data.Headers...); err != nil {
			return err
		}
	}

	if data.NextRow == nil {
		return nil
	}

	for {
		row := data.NextRow()
		if row == nil {
			return nil
		}
		if err := table.Row(row...); err != nil {
			return err
		}
	}
}
