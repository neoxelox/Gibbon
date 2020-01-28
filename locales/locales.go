package locales

import (
	"os"
	"strings"
)

// Locale describes a Locale instance.
type Locale struct {
	Type    LocaleType
	Version string
	Author  string
	copies  map[string]string
}

// LocaleType defines the supported locales.
type LocaleType string

const (
	// -- English --

	// EnUs - English from United States
	EnUs LocaleType = "en_US"

	// -- Spanish --

	// EsEs - Spanish from Spain
	EsEs LocaleType = "es_ES"

	// -- Catalan --

	// CaEs - Catalan from Spain
	CaEs LocaleType = "ca_ES"
)

// New creates a Locale instance.
func New(localeType LocaleType) (*Locale, error) {
	data, err := os.Open("./locales/" + string(localeType) + ".po")
	if err != nil {
		return nil, err
	}
	defer data.Close()

	// -- WIP --
	// DELETE!
	copies := map[string]string{
		"REPL_INIT": "Gibbon {version} ({date})\n{hostname} on {platform}",
	}

	l := &Locale{
		Type:    localeType,
		Version: "WIP",
		Author:  "Alex",
		copies:  copies,
	}

	// -- WIP --

	return l, nil
}

// F retrieves and formats a Copi.
func (l *Locale) F(copi string, params map[string]string) string {
	if cp, exists := l.copies[copi]; exists {
		for key, param := range params {
			cp = strings.ReplaceAll(cp, "{"+key+"}", param)
		}
		return cp
	}
	return copi
}
