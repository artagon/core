// This file was auto-generated by the vanadium vdl tool.
// Package: gotagtest

//nolint:golint
package gotagtest

import (
	"v.io/v23/vdl"
)

var _ = initializeVDL() // Must be first; see initializeVDL comments for details.

// Type definitions
// ================

type TestStructA struct {
	A string `json:"ja,omitempty"`
	B int64
}

func (TestStructA) VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/cmd/vdl/gotagtest.TestStructA"`
}) {
}

func (x TestStructA) VDLIsZero() bool { //nolint:gocyclo
	return x == TestStructA{}
}

func (x TestStructA) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct1); err != nil {
		return err
	}
	if x.A != "" {
		if err := enc.NextFieldValueString(0, vdl.StringType, x.A); err != nil {
			return err
		}
	}
	if x.B != 0 {
		if err := enc.NextFieldValueInt(1, vdl.Int64Type, x.B); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *TestStructA) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = TestStructA{}
	if err := dec.StartValue(vdlTypeStruct1); err != nil {
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
		if decType != vdlTypeStruct1 {
			index = vdlTypeStruct1.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.A = value
			}
		case 1:
			switch value, err := dec.ReadValueInt(64); {
			case err != nil:
				return err
			default:
				x.B = value
			}
		}
	}
}

type TestStructB struct {
	A int64
	B string `json:"jb,omitempty"`
}

func (TestStructB) VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/cmd/vdl/gotagtest.TestStructB"`
}) {
}

func (x TestStructB) VDLIsZero() bool { //nolint:gocyclo
	return x == TestStructB{}
}

func (x TestStructB) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct2); err != nil {
		return err
	}
	if x.A != 0 {
		if err := enc.NextFieldValueInt(0, vdl.Int64Type, x.A); err != nil {
			return err
		}
	}
	if x.B != "" {
		if err := enc.NextFieldValueString(1, vdl.StringType, x.B); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *TestStructB) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = TestStructB{}
	if err := dec.StartValue(vdlTypeStruct2); err != nil {
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
		if decType != vdlTypeStruct2 {
			index = vdlTypeStruct2.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueInt(64); {
			case err != nil:
				return err
			default:
				x.A = value
			}
		case 1:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.B = value
			}
		}
	}
}

// Hold type definitions in package-level variables, for better performance.
//nolint:unused
var (
	vdlTypeStruct1 *vdl.Type
	vdlTypeStruct2 *vdl.Type
)

var initializeVDLCalled bool

// initializeVDL performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = initializeVDL()
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
	vdl.Register((*TestStructA)(nil))
	vdl.Register((*TestStructB)(nil))

	// Initialize type definitions.
	vdlTypeStruct1 = vdl.TypeOf((*TestStructA)(nil)).Elem()
	vdlTypeStruct2 = vdl.TypeOf((*TestStructB)(nil)).Elem()
	return struct{}{}
}
