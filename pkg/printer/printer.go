package printer

import (
	"strings"

	"github.com/pkg/errors"
)

type Format string

const (
	FormatJson      Format = "json"
	FormatTable     Format = "table"
	FormatYaml      Format = "yaml"
	FormatPlaintext Format = "plaintext"
)

type Formats []Format

func (f Formats) String() string {
	var buf strings.Builder
	for i, format := range f {
		buf.WriteString(string(format))
		if i < len(f)-1 {
			buf.WriteString(", ")
		}
	}
	return buf.String()
}

func (p *PrintOptions) ParseFormat() error {
	format := Format(p.RawFormat)
	switch format {
	case FormatTable, FormatJson, FormatYaml, FormatPlaintext:
		p.Format = format
		return nil
	default:
		return errors.Errorf("invalid format: %s", p.RawFormat)
	}
}

func (p *PrintOptions) Validate(allowedFormats []Format) error {
	// Default unspecified to plaintext
	if p.RawFormat == "" {
		p.RawFormat = string(FormatPlaintext)
	}

	format := Format(p.RawFormat)
	for _, f := range allowedFormats {
		if f == format {
			p.Format = format
			return nil
		}
	}
	return errors.Errorf("invalid format: %s", p.RawFormat)
}

type PrintOptions struct {
	RawFormat string
	Format
}
