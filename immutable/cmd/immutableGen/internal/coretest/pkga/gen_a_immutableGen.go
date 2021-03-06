// Code generated by immutableGen. DO NOT EDIT.

package pkga

//immutableVet:skipFile

import (
	"myitcv.io/immutable"

	"myitcv.io/immutable/cmd/immutableGen/internal/coretest/pkgb"
)

//
// PkgA is an immutable type and has the following template:
//
// 	struct {
// 		*pkgb.PkgB
// 		Address	string
// 	}
//
type PkgA struct {
	anonfield_PkgB *pkgb.PkgB
	field_Address  string

	mutable bool
	__tmpl  *_Imm_PkgA
}

var _ immutable.Immutable = new(PkgA)
var _ = new(PkgA).__tmpl

func (s *PkgA) AsMutable() *PkgA {
	if s.Mutable() {
		return s
	}

	res := *s
	res.mutable = true
	return &res
}

func (s *PkgA) AsImmutable(v *PkgA) *PkgA {
	if s == nil {
		return nil
	}

	if s == v {
		return s
	}

	s.mutable = false
	return s
}

func (s *PkgA) Mutable() bool {
	return s.mutable
}

func (s *PkgA) WithMutable(f func(si *PkgA)) *PkgA {
	res := s.AsMutable()
	f(res)
	res = res.AsImmutable(s)

	return res
}

func (s *PkgA) WithImmutable(f func(si *PkgA)) *PkgA {
	prev := s.mutable
	s.mutable = false
	f(s)
	s.mutable = prev

	return s
}

func (s *PkgA) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}

	if seen == nil {
		return s.IsDeeplyNonMutable(make(map[interface{}]bool))
	}

	if seen[s] {
		return true
	}

	seen[s] = true
	{
		v := s.anonfield_PkgB

		if v != nil && !v.IsDeeplyNonMutable(seen) {
			return false
		}
	}
	return true
}
func (s *PkgA) Address() string {
	return s.field_Address
}

// SetAddress is the setter for Address()
func (s *PkgA) SetAddress(n string) *PkgA {
	if s.mutable {
		s.field_Address = n
		return s
	}

	res := *s
	res.field_Address = n
	return &res
}
func (s *PkgA) PkgB() *pkgb.PkgB {
	return s.anonfield_PkgB
}

// SetPkgB is the setter for PkgB()
func (s *PkgA) SetPkgB(n *pkgb.PkgB) *PkgA {
	if s.mutable {
		s.anonfield_PkgB = n
		return s
	}

	res := *s
	res.anonfield_PkgB = n
	return &res
}
func (s *PkgA) Postcode() string {
	return s.PkgB().Postcode()
}
func (s *PkgA) SetPostcode(n string) *PkgA {
	v1 := s.PkgB().SetPostcode(n)
	v0 := s.SetPkgB(v1)
	return v0
}

//
// Clash2 is an immutable type and has the following template:
//
// 	struct {
// 		Clash		string
// 		NoClash2	string
// 	}
//
type Clash2 struct {
	field_Clash    string
	field_NoClash2 string

	mutable bool
	__tmpl  *_Imm_Clash2
}

var _ immutable.Immutable = new(Clash2)
var _ = new(Clash2).__tmpl

func (s *Clash2) AsMutable() *Clash2 {
	if s.Mutable() {
		return s
	}

	res := *s
	res.mutable = true
	return &res
}

func (s *Clash2) AsImmutable(v *Clash2) *Clash2 {
	if s == nil {
		return nil
	}

	if s == v {
		return s
	}

	s.mutable = false
	return s
}

func (s *Clash2) Mutable() bool {
	return s.mutable
}

func (s *Clash2) WithMutable(f func(si *Clash2)) *Clash2 {
	res := s.AsMutable()
	f(res)
	res = res.AsImmutable(s)

	return res
}

func (s *Clash2) WithImmutable(f func(si *Clash2)) *Clash2 {
	prev := s.mutable
	s.mutable = false
	f(s)
	s.mutable = prev

	return s
}

func (s *Clash2) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}

	if seen == nil {
		return s.IsDeeplyNonMutable(make(map[interface{}]bool))
	}

	if seen[s] {
		return true
	}

	seen[s] = true
	return true
}
func (s *Clash2) Clash() string {
	return s.field_Clash
}

// SetClash is the setter for Clash()
func (s *Clash2) SetClash(n string) *Clash2 {
	if s.mutable {
		s.field_Clash = n
		return s
	}

	res := *s
	res.field_Clash = n
	return &res
}
func (s *Clash2) NoClash2() string {
	return s.field_NoClash2
}

// SetNoClash2 is the setter for NoClash2()
func (s *Clash2) SetNoClash2(n string) *Clash2 {
	if s.mutable {
		s.field_NoClash2 = n
		return s
	}

	res := *s
	res.field_NoClash2 = n
	return &res
}

//
// OtherA is an immutable type and has the following template:
//
// 	struct {
// 		OtherNameA string
// 	}
//
type OtherA struct {
	field_OtherNameA string

	mutable bool
	__tmpl  *_Imm_OtherA
}

var _ immutable.Immutable = new(OtherA)
var _ = new(OtherA).__tmpl

func (s *OtherA) AsMutable() *OtherA {
	if s.Mutable() {
		return s
	}

	res := *s
	res.mutable = true
	return &res
}

func (s *OtherA) AsImmutable(v *OtherA) *OtherA {
	if s == nil {
		return nil
	}

	if s == v {
		return s
	}

	s.mutable = false
	return s
}

func (s *OtherA) Mutable() bool {
	return s.mutable
}

func (s *OtherA) WithMutable(f func(si *OtherA)) *OtherA {
	res := s.AsMutable()
	f(res)
	res = res.AsImmutable(s)

	return res
}

func (s *OtherA) WithImmutable(f func(si *OtherA)) *OtherA {
	prev := s.mutable
	s.mutable = false
	f(s)
	s.mutable = prev

	return s
}

func (s *OtherA) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}

	if seen == nil {
		return s.IsDeeplyNonMutable(make(map[interface{}]bool))
	}

	if seen[s] {
		return true
	}

	seen[s] = true
	return true
}
func (s *OtherA) OtherNameA() string {
	return s.field_OtherNameA
}

// SetOtherNameA is the setter for OtherNameA()
func (s *OtherA) SetOtherNameA(n string) *OtherA {
	if s.mutable {
		s.field_OtherNameA = n
		return s
	}

	res := *s
	res.field_OtherNameA = n
	return &res
}
