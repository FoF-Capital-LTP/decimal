package decimal

import (
	"fmt"
)

var (
	MaxNumber = Decimal{neg: false, scale: 0, coef: maxFint} // Max represents the maximum decimal value.
	powerOf10 = []Decimal{
		{neg: false, scale: 18, coef: 1},
		{neg: false, scale: 17, coef: 1},
		{neg: false, scale: 16, coef: 1},
		{neg: false, scale: 15, coef: 1},
		{neg: false, scale: 14, coef: 1},
		{neg: false, scale: 13, coef: 1},
		{neg: false, scale: 12, coef: 1},
		{neg: false, scale: 11, coef: 1},
		{neg: false, scale: 10, coef: 1},
		{neg: false, scale: 9, coef: 1},
		{neg: false, scale: 8, coef: 1},
		{neg: false, scale: 7, coef: 1},
		{neg: false, scale: 6, coef: 1},
		{neg: false, scale: 5, coef: 1},
		{neg: false, scale: 4, coef: 1},
		{neg: false, scale: 3, coef: 1},
		{neg: false, scale: 2, coef: 1},
		{neg: false, scale: 1, coef: 1},
		{neg: false, scale: 0, coef: 1},
		{neg: false, scale: 0, coef: 10},
		{neg: false, scale: 0, coef: 100},
		{neg: false, scale: 0, coef: 1_000},
		{neg: false, scale: 0, coef: 10_000},
		{neg: false, scale: 0, coef: 100_000},
		{neg: false, scale: 0, coef: 1_000_000},
		{neg: false, scale: 0, coef: 10_000_000},
		{neg: false, scale: 0, coef: 100_000_000},
		{neg: false, scale: 0, coef: 1_000_000_000},
		{neg: false, scale: 0, coef: 10_000_000_000},
		{neg: false, scale: 0, coef: 100_000_000_000},
		{neg: false, scale: 0, coef: 1_000_000_000_000},
		{neg: false, scale: 0, coef: 10_000_000_000_000},
		{neg: false, scale: 0, coef: 100_000_000_000_000},
		{neg: false, scale: 0, coef: 1_000_000_000_000_000},
		{neg: false, scale: 0, coef: 10_000_000_000_000_000},
		{neg: false, scale: 0, coef: 100_000_000_000_000_000},
		{neg: false, scale: 0, coef: 1_000_000_000_000_000_000},
	}
)

func PowOf10(exp int) Decimal {
	if exp < -18 || exp > 18 {
		panic("exponent out of range")
	}
	return powerOf10[exp+18]
}

func MustNewFromFloat64(f float64) Decimal {
	d, err := NewFromFloat64(f)
	if err != nil {
		panic(fmt.Sprintf("NewFromFloat64(%v) failed: %v", f, err))
	}
	return d
}

func MustNewFromString(s string) Decimal {
	d, err := Parse(s)
	if err != nil {
		panic(fmt.Sprintf("NewFromString(%q) failed: %v", s, err))
	}
	return d
}

func NewFromString(s string) (Decimal, error) {
	return ParseExact(s, 0)
}

func NewUnsafe(neg bool, scale int8, coef uint64) Decimal {
	return Decimal{
		neg:   neg,
		scale: scale,
		coef:  fint(coef),
	}
}

func RequireFromString(s string) Decimal {
	d, err := NewFromString(s)
	if err != nil {
		panic(fmt.Sprintf("NewFromString(%q) failed: %v", s, err))
	}
	return d
}

func (d Decimal) IntPart() int64 {
	whole, _, _ := d.Int64(0)
	return whole
}

// GobEncode implements the gob.GobEncoder interface for gob serialization.
func (d Decimal) GobEncode() ([]byte, error) {
	return d.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface for gob serialization.
func (d *Decimal) GobDecode(data []byte) error {
	return d.UnmarshalBinary(data)
}

func unquoteIfQuoted(value []byte) (string, error) {
	// If the amount is quoted, strip the quotes
	if len(value) > 2 && value[0] == '"' && value[len(value)-1] == '"' {
		value = value[1 : len(value)-1]
	}
	return string(value), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *Decimal) UnmarshalJSON(decimalBytes []byte) error {
	if string(decimalBytes) == "null" {
		return nil
	}

	str, err := unquoteIfQuoted(decimalBytes)
	if err != nil {
		return fmt.Errorf("error decoding string '%s': %s", decimalBytes, err)
	}

	decimal, err := NewFromString(str)
	*d = decimal
	if err != nil {
		return fmt.Errorf("error decoding string '%s': %s", str, err)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (d Decimal) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d Decimal) MulIgnoreError(e Decimal) Decimal {
	res, _ := d.Mul(e)
	return res
}

func (d Decimal) MustMul(e Decimal) Decimal {
	res, err := d.Mul(e)
	if err != nil {
		panic(err)
	}
	return res
}

func (d Decimal) PowIntIgnoreError(e int) Decimal {
	res, _ := d.PowInt(e)
	return res
}

func (d Decimal) MustPowInt(e int) Decimal {
	res, err := d.PowInt(e)
	if err != nil {
		panic(err)
	}
	return res
}
func (d Decimal) SqrtIgnoreError() Decimal {
	res, _ := d.Sqrt()
	return res
}

func (d Decimal) MustSqrt() Decimal {
	res, err := d.Sqrt()
	if err != nil {
		panic(err)
	}
	return res
}

func (d Decimal) ExpIgnoreError() Decimal {
	res, _ := d.Exp()
	return res
}

func (d Decimal) MustExp() Decimal {
	res, err := d.Exp()
	if err != nil {
		panic(err)
	}
	return res
}

func (d Decimal) LogIgnoreError() Decimal {
	res, _ := d.Log()
	return res
}

func (d Decimal) MustLog() Decimal {
	res, err := d.Log()
	if err != nil {
		panic(err)
	}
	return res
}

func (d Decimal) SubIgnoreError(e Decimal) Decimal {
	res, _ := d.Sub(e)
	return res
}

func (d Decimal) MustSub(e Decimal) Decimal {
	res, err := d.Sub(e)
	if err != nil {
		panic(err)
	}
	return res
}

func (d Decimal) AddIgnoreError(e Decimal) Decimal {
	res, _ := d.Add(e)
	return res
}

func (d Decimal) MustAdd(e Decimal) Decimal {
	res, err := d.Add(e)
	if err != nil {
		panic(err)
	}
	return res
}

func (d Decimal) SubMulIgnoreError(e, f Decimal) Decimal {
	res, _ := d.SubMul(e, f)
	return res
}

func (d Decimal) MustSubMul(e, f Decimal) Decimal {
	res, err := d.SubMul(e, f)
	if err != nil {
		panic(err)
	}
	return res
}

func (d Decimal) AddMulIgnoreError(e, f Decimal) Decimal {
	res, _ := d.AddMul(e, f)
	return res
}

func (d Decimal) MustAddMul(e, f Decimal) Decimal {
	res, err := d.AddMul(e, f)
	if err != nil {
		panic(err)
	}
	return res
}

func (d Decimal) SubQuoIgnoreError(e, f Decimal) Decimal {
	res, _ := d.SubQuo(e, f)
	return res
}

func (d Decimal) MustSubQuo(e, f Decimal) Decimal {
	res, err := d.SubQuo(e, f)
	if err != nil {
		panic(err)
	}
	return res
}

func (d Decimal) AddQuoIgnoreError(e, f Decimal) Decimal {
	res, _ := d.AddQuo(e, f)
	return res
}

func (d Decimal) MustAddQuo(e, f Decimal) Decimal {
	res, err := d.AddQuo(e, f)
	if err != nil {
		panic(err)
	}
	return res
}

func (d Decimal) InvIgnoreError() Decimal {
	res, _ := d.Inv()
	return res
}

func (d Decimal) MustInv() Decimal {
	res, err := d.Inv()
	if err != nil {
		panic(err)
	}
	return res
}

func (d Decimal) QuoIgnoreError(e Decimal) Decimal {
	res, _ := d.Quo(e)
	return res
}

func (d Decimal) MustQuo(e Decimal) Decimal {
	res, err := d.Quo(e)
	if err != nil {
		panic(err)
	}
	return res
}

func (d Decimal) QuoRemIgnoreError(e Decimal) (Decimal, Decimal) {
	q, r, _ := d.QuoRem(e)
	return q, r
}

func (d Decimal) MustQuoRem(e Decimal) (Decimal, Decimal) {
	q, r, err := d.QuoRem(e)
	if err != nil {
		panic(err)
	}
	return q, r
}

// Equals compares decimals and returns:
//
//	 true if d = e
//	false otherwise
//
// See also method [Decimal.Cmp].
func (d Decimal) Equals(e Decimal) bool {
	return d.Cmp(e) == 0
}

// GreaterThan (GT) returns true when d is greater than d2.
func (d Decimal) GreaterThan(d2 Decimal) bool {
	return d.Cmp(d2) == 1
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (d Decimal) GreaterThanOrEqual(d2 Decimal) bool {
	cmp := d.Cmp(d2)
	return cmp == 1 || cmp == 0
}

// LessThan (LT) returns true when d is less than d2.
func (d Decimal) LessThan(d2 Decimal) bool {
	return d.Cmp(d2) == -1
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (d Decimal) LessThanOrEqual(d2 Decimal) bool {
	cmp := d.Cmp(d2)
	return cmp == -1 || cmp == 0
}

// IsPositive return
//
//	true if d > 0
//	false if d == 0
//	false if d < 0
func (d Decimal) IsPositive() bool {
	return d.Sign() == 1
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (d Decimal) IsNegative() bool {
	return d.Sign() == -1
}
