// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: discovery
//
//nolint:revive
package discovery

import (
	"v.io/v23/vdl"
)

var initializeVDLCalled = false
var _ = initializeVDL() // Must be first; see initializeVDL comments for details.

// Hold type definitions in package-level variables, for better performance.
// Declare and initialize with default values here so that the initializeVDL
// method will be considered ready to initialize before any of the type
// definitions that appear below.
//
//nolint:unused
var (
	vdlTypeArray1  *vdl.Type = nil
	vdlTypeMap2    *vdl.Type = nil
	vdlTypeMap3    *vdl.Type = nil
	vdlTypeList4   *vdl.Type = nil
	vdlTypeStruct5 *vdl.Type = nil
	vdlTypeList6   *vdl.Type = nil
)

// Type definitions
// ================
// An AdId is a globally unique identifier of an advertisement.
type AdId [16]byte

func (AdId) VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.AdId"`
}) {
}

func (x AdId) VDLIsZero() bool { //nolint:gocyclo
	return x == AdId{}
}

func (x AdId) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.WriteValueBytes(vdlTypeArray1, x[:]); err != nil {
		return err
	}
	return nil
}

func (x *AdId) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	bytes := x[:]
	if err := dec.ReadValueBytes(16, &bytes); err != nil {
		return err
	}
	return nil
}

// Attributes represents service attributes as a key/value pair.
type Attributes map[string]string

func (Attributes) VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.Attributes"`
}) {
}

func (x Attributes) VDLIsZero() bool { //nolint:gocyclo
	return len(x) == 0
}

func (x Attributes) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeMap2); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for key, elem := range x {
		if err := enc.NextEntryValueString(vdl.StringType, key); err != nil {
			return err
		}
		if err := enc.WriteValueString(vdl.StringType, elem); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Attributes) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	if err := dec.StartValue(vdlTypeMap2); err != nil {
		return err
	}
	var tmpMap Attributes
	if len := dec.LenHint(); len > 0 {
		tmpMap = make(Attributes, len)
	}
	for {
		switch done, key, err := dec.NextEntryValueString(); {
		case err != nil:
			return err
		case done:
			*x = tmpMap
			return dec.FinishValue()
		default:
			var elem string
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				elem = value
			}
			if tmpMap == nil {
				tmpMap = make(Attributes)
			}
			tmpMap[key] = elem
		}
	}
}

// Attachments represents service attachments as a key/value pair.
type Attachments map[string][]byte

func (Attachments) VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.Attachments"`
}) {
}

func (x Attachments) VDLIsZero() bool { //nolint:gocyclo
	return len(x) == 0
}

func (x Attachments) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeMap3); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for key, elem := range x {
		if err := enc.NextEntryValueString(vdl.StringType, key); err != nil {
			return err
		}
		if err := enc.WriteValueBytes(vdlTypeList4, elem); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Attachments) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	if err := dec.StartValue(vdlTypeMap3); err != nil {
		return err
	}
	var tmpMap Attachments
	if len := dec.LenHint(); len > 0 {
		tmpMap = make(Attachments, len)
	}
	for {
		switch done, key, err := dec.NextEntryValueString(); {
		case err != nil:
			return err
		case done:
			*x = tmpMap
			return dec.FinishValue()
		default:
			var elem []byte
			if err := dec.ReadValueBytes(-1, &elem); err != nil {
				return err
			}
			if tmpMap == nil {
				tmpMap = make(Attachments)
			}
			tmpMap[key] = elem
		}
	}
}

// Advertisement represents a feed into advertiser to broadcast its contents
// to scanners.
//
// A large advertisement may require additional RPC calls causing delay in
// discovery. We limit the maximum size of an advertisement to 512 bytes
// excluding id and attachments.
type Advertisement struct {
	// Universal unique identifier of the advertisement.
	// If this is not specified, a random unique identifier will be assigned.
	Id AdId
	// Interface name that the advertised service implements.
	// E.g., 'v.io/v23/services/vtrace.Store'.
	InterfaceName string
	// Addresses (vanadium object names) that the advertised service is served on.
	// E.g., '/host:port/a/b/c', '/ns.dev.v.io:8101/blah/blah'.
	Addresses []string
	// Attributes as a key/value pair.
	// E.g., {'resolution': '1024x768'}.
	//
	// The key must be US-ASCII printable characters, excluding the '=' character
	// and should not start with '_' character.
	//
	// We limit the maximum number of attachments to 32.
	Attributes Attributes
	// Attachments as a key/value pair.
	// E.g., {'thumbnail': binary_data }.
	//
	// Unlike attributes, attachments are for binary data and they are not queryable.
	//
	// The key must be US-ASCII printable characters, excluding the '=' character
	// and should not start with '_' character.
	//
	// We limit the maximum number of attachments to 32 and the maximum size of each
	// attachment is 4K bytes.
	Attachments Attachments
}

func (Advertisement) VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.Advertisement"`
}) {
}

func (x Advertisement) VDLIsZero() bool { //nolint:gocyclo
	if x.Id != (AdId{}) {
		return false
	}
	if x.InterfaceName != "" {
		return false
	}
	if len(x.Addresses) != 0 {
		return false
	}
	if len(x.Attributes) != 0 {
		return false
	}
	if len(x.Attachments) != 0 {
		return false
	}
	return true
}

func (x Advertisement) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct5); err != nil {
		return err
	}
	if x.Id != (AdId{}) {
		if err := enc.NextFieldValueBytes(0, vdlTypeArray1, x.Id[:]); err != nil {
			return err
		}
	}
	if x.InterfaceName != "" {
		if err := enc.NextFieldValueString(1, vdl.StringType, x.InterfaceName); err != nil {
			return err
		}
	}
	if len(x.Addresses) != 0 {
		if err := enc.NextField(2); err != nil {
			return err
		}
		if err := vdlWriteAnonList1(enc, x.Addresses); err != nil {
			return err
		}
	}
	if len(x.Attributes) != 0 {
		if err := enc.NextField(3); err != nil {
			return err
		}
		if err := x.Attributes.VDLWrite(enc); err != nil {
			return err
		}
	}
	if len(x.Attachments) != 0 {
		if err := enc.NextField(4); err != nil {
			return err
		}
		if err := x.Attachments.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func vdlWriteAnonList1(enc vdl.Encoder, x []string) error {
	if err := enc.StartValue(vdlTypeList6); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for _, elem := range x {
		if err := enc.NextEntryValueString(vdl.StringType, elem); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Advertisement) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = Advertisement{}
	if err := dec.StartValue(vdlTypeStruct5); err != nil {
		return err
	}
	decType := dec.Type()
	for {
		index, err := dec.NextField()
		switch {
		case err != nil:
			return err
		case index == -1:
			return dec.FinishValue()
		}
		if decType != vdlTypeStruct5 {
			index = vdlTypeStruct5.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			bytes := x.Id[:]
			if err := dec.ReadValueBytes(16, &bytes); err != nil {
				return err
			}
		case 1:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.InterfaceName = value
			}
		case 2:
			if err := vdlReadAnonList1(dec, &x.Addresses); err != nil {
				return err
			}
		case 3:
			if err := x.Attributes.VDLRead(dec); err != nil {
				return err
			}
		case 4:
			if err := x.Attachments.VDLRead(dec); err != nil {
				return err
			}
		}
	}
}

func vdlReadAnonList1(dec vdl.Decoder, x *[]string) error {
	if err := dec.StartValue(vdlTypeList6); err != nil {
		return err
	}
	if len := dec.LenHint(); len > 0 {
		*x = make([]string, 0, len)
	} else {
		*x = nil
	}
	for {
		switch done, elem, err := dec.NextEntryValueString(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		default:
			*x = append(*x, elem)
		}
	}
}

// initializeVDL performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
// var _ = initializeVDL()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func initializeVDL() struct{} {
	if initializeVDLCalled {
		return struct{}{}
	}
	initializeVDLCalled = true

	// Register types.
	vdl.Register((*AdId)(nil))
	vdl.Register((*Attributes)(nil))
	vdl.Register((*Attachments)(nil))
	vdl.Register((*Advertisement)(nil))

	// Initialize type definitions.
	vdlTypeArray1 = vdl.TypeOf((*AdId)(nil))
	vdlTypeMap2 = vdl.TypeOf((*Attributes)(nil))
	vdlTypeMap3 = vdl.TypeOf((*Attachments)(nil))
	vdlTypeList4 = vdl.TypeOf((*[]byte)(nil))
	vdlTypeStruct5 = vdl.TypeOf((*Advertisement)(nil)).Elem()
	vdlTypeList6 = vdl.TypeOf((*[]string)(nil))

	return struct{}{}
}
