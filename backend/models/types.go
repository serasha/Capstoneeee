package models

import (
	"strconv"
	"strings"
)

// FlexibleInt adalah tipe int yang bisa di-unmarshal dari string atau number JSON
// Contoh yang diterima: 123 atau "123"
type FlexibleInt int

func (fi *FlexibleInt) UnmarshalJSON(b []byte) error {
	s := strings.TrimSpace(string(b))
	if s == "" || s == "null" {
		*fi = 0
		return nil
	}
	// Jika dalam tanda kutip, treat sebagai string number
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
		if s == "" {
			*fi = 0
			return nil
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		*fi = FlexibleInt(n)
		return nil
	}
	// Jika bukan string, parse sebagai number
	if strings.ContainsAny(s, ".eE") {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		*fi = FlexibleInt(int(f))
		return nil
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*fi = FlexibleInt(n)
	return nil
} 