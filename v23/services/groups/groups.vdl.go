// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: groups

// Package groups defines interfaces for managing access control groups.  Groups
// can be referenced by BlessingPatterns (e.g. in AccessLists).
//nolint:golint
package groups

import (
	"fmt"

	v23 "v.io/v23"
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/rpc"
	"v.io/v23/security/access"
	"v.io/v23/services/permissions"
	"v.io/v23/vdl"
	"v.io/v23/verror"
)

var _ = initializeVDL() // Must be first; see initializeVDL comments for details.

//////////////////////////////////////////////////
// Type definitions

// BlessingPatternChunk is a substring of a BlessingPattern. As with
// BlessingPatterns, BlessingPatternChunks may contain references to groups.
// However, they may be restricted in other ways. For example, in the future
// BlessingPatterns may support "$" terminators, but these may be disallowed for
// BlessingPatternChunks.
type BlessingPatternChunk string

func (BlessingPatternChunk) VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/groups.BlessingPatternChunk"`
}) {
}

func (x BlessingPatternChunk) VDLIsZero() bool { //nolint:gocyclo
	return x == ""
}

func (x BlessingPatternChunk) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.WriteValueString(vdlTypeString1, string(x)); err != nil {
		return err
	}
	return nil
}

func (x *BlessingPatternChunk) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	switch value, err := dec.ReadValueString(); {
	case err != nil:
		return err
	default:
		*x = BlessingPatternChunk(value)
	}
	return nil
}

type GetRequest struct {
}

func (GetRequest) VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/groups.GetRequest"`
}) {
}

func (x GetRequest) VDLIsZero() bool { //nolint:gocyclo
	return x == GetRequest{}
}

func (x GetRequest) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct2); err != nil {
		return err
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *GetRequest) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = GetRequest{}
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
		}
	}
}

type GetResponse struct {
	Entries map[BlessingPatternChunk]struct{}
}

func (GetResponse) VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/groups.GetResponse"`
}) {
}

func (x GetResponse) VDLIsZero() bool { //nolint:gocyclo
	return len(x.Entries) == 0
}

func (x GetResponse) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct3); err != nil {
		return err
	}
	if len(x.Entries) != 0 {
		if err := enc.NextField(0); err != nil {
			return err
		}
		if err := vdlWriteAnonSet1(enc, x.Entries); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func vdlWriteAnonSet1(enc vdl.Encoder, x map[BlessingPatternChunk]struct{}) error {
	if err := enc.StartValue(vdlTypeSet4); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for key := range x {
		if err := enc.NextEntryValueString(vdlTypeString1, string(key)); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *GetResponse) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = GetResponse{}
	if err := dec.StartValue(vdlTypeStruct3); err != nil {
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
		if decType != vdlTypeStruct3 {
			index = vdlTypeStruct3.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		if index == 0 {

			if err := vdlReadAnonSet1(dec, &x.Entries); err != nil {
				return err
			}
		}
	}
}

func vdlReadAnonSet1(dec vdl.Decoder, x *map[BlessingPatternChunk]struct{}) error {
	if err := dec.StartValue(vdlTypeSet4); err != nil {
		return err
	}
	var tmpMap map[BlessingPatternChunk]struct{}
	if len := dec.LenHint(); len > 0 {
		tmpMap = make(map[BlessingPatternChunk]struct{}, len)
	}
	for {
		switch done, key, err := dec.NextEntryValueString(); {
		case err != nil:
			return err
		case done:
			*x = tmpMap
			return dec.FinishValue()
		default:
			if tmpMap == nil {
				tmpMap = make(map[BlessingPatternChunk]struct{})
			}
			tmpMap[BlessingPatternChunk(key)] = struct{}{}
		}
	}
}

// ApproximationType defines the type of approximation desired when a Relate
// call encounters an error (inaccessible or undefined group in a blessing
// pattern, cyclic group definitions, storage errors, invalid patterns
// etc). "Under" is used for blessing patterns in "Allow" clauses in an
// AccessList, while "Over" is used for blessing patterns in "Deny" clauses.
type ApproximationType int

const (
	ApproximationTypeUnder ApproximationType = iota
	ApproximationTypeOver
)

// ApproximationTypeAll holds all labels for ApproximationType.
var ApproximationTypeAll = [...]ApproximationType{ApproximationTypeUnder, ApproximationTypeOver}

// ApproximationTypeFromString creates a ApproximationType from a string label.
//nolint:deadcode,unused
func ApproximationTypeFromString(label string) (x ApproximationType, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *ApproximationType) Set(label string) error {
	switch label {
	case "Under", "under":
		*x = ApproximationTypeUnder
		return nil
	case "Over", "over":
		*x = ApproximationTypeOver
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in groups.ApproximationType", label)
}

// String returns the string label of x.
func (x ApproximationType) String() string {
	switch x {
	case ApproximationTypeUnder:
		return "Under"
	case ApproximationTypeOver:
		return "Over"
	}
	return ""
}

func (ApproximationType) VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/groups.ApproximationType"`
	Enum struct{ Under, Over string }
}) {
}

func (x ApproximationType) VDLIsZero() bool { //nolint:gocyclo
	return x == ApproximationTypeUnder
}

func (x ApproximationType) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.WriteValueString(vdlTypeEnum5, x.String()); err != nil {
		return err
	}
	return nil
}

func (x *ApproximationType) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	switch value, err := dec.ReadValueString(); {
	case err != nil:
		return err
	default:
		if err := x.Set(value); err != nil {
			return err
		}
	}
	return nil
}

// Approximation contains information about membership approximations made
// during a Relate call.
type Approximation struct {
	Reason  string
	Details string
}

func (Approximation) VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/groups.Approximation"`
}) {
}

func (x Approximation) VDLIsZero() bool { //nolint:gocyclo
	return x == Approximation{}
}

func (x Approximation) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct6); err != nil {
		return err
	}
	if x.Reason != "" {
		if err := enc.NextFieldValueString(0, vdl.StringType, x.Reason); err != nil {
			return err
		}
	}
	if x.Details != "" {
		if err := enc.NextFieldValueString(1, vdl.StringType, x.Details); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Approximation) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = Approximation{}
	if err := dec.StartValue(vdlTypeStruct6); err != nil {
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
		if decType != vdlTypeStruct6 {
			index = vdlTypeStruct6.FieldIndexByName(decType.Field(index).Name)
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
				x.Reason = value
			}
		case 1:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Details = value
			}
		}
	}
}

//////////////////////////////////////////////////
// Error definitions

var (
	ErrNoBlessings         = verror.NewIDAction("v.io/v23/services/groups.NoBlessings", verror.NoRetry)
	ErrExcessiveContention = verror.NewIDAction("v.io/v23/services/groups.ExcessiveContention", verror.RetryBackoff)
	ErrCycleFound          = verror.NewIDAction("v.io/v23/services/groups.CycleFound", verror.NoRetry)
)

// NewErrNoBlessings returns an error with the ErrNoBlessings ID.
// WARNING: this function is deprecated and will be removed in the future,
// use ErrorfNoBlessings or MessageNoBlessings instead.
func NewErrNoBlessings(ctx *context.T) error {
	return verror.New(ErrNoBlessings, ctx)
}

// ErrorfNoBlessings calls ErrNoBlessings.Errorf with the supplied arguments.
func ErrorfNoBlessings(ctx *context.T, format string) error {
	return ErrNoBlessings.Errorf(ctx, format)
}

// MessageNoBlessings calls ErrNoBlessings.Message with the supplied arguments.
func MessageNoBlessings(ctx *context.T, message string) error {
	return ErrNoBlessings.Message(ctx, message)
}

// ParamsErrNoBlessings extracts the expected parameters from the error's ParameterList.
func ParamsErrNoBlessings(argumentError error) (verrorComponent string, verrorOperation string, returnErr error) {
	params := verror.Params(argumentError)
	if params == nil {
		returnErr = fmt.Errorf("no parameters found in: %T: %v", argumentError, argumentError)
		return
	}
	iter := &paramListIterator{params: params, max: len(params)}

	if verrorComponent, verrorOperation, returnErr = iter.preamble(); returnErr != nil {
		return
	}

	return
}

// NewErrExcessiveContention returns an error with the ErrExcessiveContention ID.
// WARNING: this function is deprecated and will be removed in the future,
// use ErrorfExcessiveContention or MessageExcessiveContention instead.
func NewErrExcessiveContention(ctx *context.T) error {
	return verror.New(ErrExcessiveContention, ctx)
}

// ErrorfExcessiveContention calls ErrExcessiveContention.Errorf with the supplied arguments.
func ErrorfExcessiveContention(ctx *context.T, format string) error {
	return ErrExcessiveContention.Errorf(ctx, format)
}

// MessageExcessiveContention calls ErrExcessiveContention.Message with the supplied arguments.
func MessageExcessiveContention(ctx *context.T, message string) error {
	return ErrExcessiveContention.Message(ctx, message)
}

// ParamsErrExcessiveContention extracts the expected parameters from the error's ParameterList.
func ParamsErrExcessiveContention(argumentError error) (verrorComponent string, verrorOperation string, returnErr error) {
	params := verror.Params(argumentError)
	if params == nil {
		returnErr = fmt.Errorf("no parameters found in: %T: %v", argumentError, argumentError)
		return
	}
	iter := &paramListIterator{params: params, max: len(params)}

	if verrorComponent, verrorOperation, returnErr = iter.preamble(); returnErr != nil {
		return
	}

	return
}

// NewErrCycleFound returns an error with the ErrCycleFound ID.
// WARNING: this function is deprecated and will be removed in the future,
// use ErrorfCycleFound or MessageCycleFound instead.
func NewErrCycleFound(ctx *context.T, name string, visited string) error {
	return verror.New(ErrCycleFound, ctx, name, visited)
}

// ErrorfCycleFound calls ErrCycleFound.Errorf with the supplied arguments.
func ErrorfCycleFound(ctx *context.T, format string, name string, visited string) error {
	return ErrCycleFound.Errorf(ctx, format, name, visited)
}

// MessageCycleFound calls ErrCycleFound.Message with the supplied arguments.
func MessageCycleFound(ctx *context.T, message string, name string, visited string) error {
	return ErrCycleFound.Message(ctx, message, name, visited)
}

// ParamsErrCycleFound extracts the expected parameters from the error's ParameterList.
func ParamsErrCycleFound(argumentError error) (verrorComponent string, verrorOperation string, name string, visited string, returnErr error) {
	params := verror.Params(argumentError)
	if params == nil {
		returnErr = fmt.Errorf("no parameters found in: %T: %v", argumentError, argumentError)
		return
	}
	iter := &paramListIterator{params: params, max: len(params)}

	if verrorComponent, verrorOperation, returnErr = iter.preamble(); returnErr != nil {
		return
	}

	var (
		tmp interface{}
		ok  bool
	)
	tmp, returnErr = iter.next()
	if name, ok = tmp.(string); !ok {
		if returnErr != nil {
			return
		}
		returnErr = fmt.Errorf("parameter list contains the wrong type for return value name, has %T and not string", tmp)
		return
	}
	tmp, returnErr = iter.next()
	if visited, ok = tmp.(string); !ok {
		if returnErr != nil {
			return
		}
		returnErr = fmt.Errorf("parameter list contains the wrong type for return value visited, has %T and not string", tmp)
		return
	}

	return
}

type paramListIterator struct {
	err      error
	idx, max int
	params   []interface{}
}

func (pl *paramListIterator) next() (interface{}, error) {
	if pl.err != nil {
		return nil, pl.err
	}
	if pl.idx+1 > pl.max {
		pl.err = fmt.Errorf("too few parameters: have %v", pl.max)
		return nil, pl.err
	}
	pl.idx++
	return pl.params[pl.idx-1], nil
}

func (pl *paramListIterator) preamble() (component, operation string, err error) {
	var tmp interface{}
	if tmp, err = pl.next(); err != nil {
		return
	}
	var ok bool
	if component, ok = tmp.(string); !ok {
		return "", "", fmt.Errorf("ParamList[0]: component name is not a string: %T", tmp)
	}
	if tmp, err = pl.next(); err != nil {
		return
	}
	if operation, ok = tmp.(string); !ok {
		return "", "", fmt.Errorf("ParamList[1]: operation name is not a string: %T", tmp)
	}
	return
}

//////////////////////////////////////////////////
// Interface definitions

// GroupReaderClientMethods is the client interface
// containing GroupReader methods.
//
// GroupReader implements methods to read or query a group's membership
// information.
type GroupReaderClientMethods interface {
	// Relate determines the relationships between the provided blessing
	// names and the members of the group.
	//
	// Given an input set of blessing names and a group defined by a set of
	// blessing patterns S, for each blessing name B in the input, Relate(B)
	// returns a set of "remainders" consisting of every blessing name B"
	// such that there exists some B' for which B = B' B" and B' is in S,
	// and "" if B is a member of S.
	//
	// For example, if a group is defined as S = {n1, n1:n2, n1:n2:n3}, then
	// Relate(n1:n2) = {n2, ""}.
	//
	// reqVersion specifies the expected version of the group's membership
	// information. If this version is set and matches the Group's current
	// version, the response will indicate that fact but will otherwise be
	// empty.
	//
	// visitedGroups is the set of groups already visited in a particular
	// chain of Relate calls, and is used to detect the presence of
	// cycles. When a cycle is detected, it is treated just like any other
	// error, and the result is approximated.
	//
	// Relate also returns information about all the errors encountered that
	// resulted in approximations, if any.
	//
	// TODO(hpucha): scrub "Approximation" for preserving privacy. Flesh
	// versioning out further. Other args we may need: option to Get() the
	// membership set when allowed (to avoid an extra RPC), options related
	// to caching this information.
	Relate(_ *context.T, blessings map[string]struct{}, hint ApproximationType, reqVersion string, visitedGroups map[string]struct{}, _ ...rpc.CallOpt) (remainder map[string]struct{}, approximations []Approximation, version string, _ error)
	// Get returns all entries in the group.
	// TODO(sadovsky): Flesh out this API.
	Get(_ *context.T, req GetRequest, reqVersion string, _ ...rpc.CallOpt) (res GetResponse, version string, _ error)
}

// GroupReaderClientStub embeds GroupReaderClientMethods and is a
// placeholder for additional management operations.
type GroupReaderClientStub interface {
	GroupReaderClientMethods
}

// GroupReaderClient returns a client stub for GroupReader.
func GroupReaderClient(name string) GroupReaderClientStub {
	return implGroupReaderClientStub{name}
}

type implGroupReaderClientStub struct {
	name string
}

func (c implGroupReaderClientStub) Relate(ctx *context.T, i0 map[string]struct{}, i1 ApproximationType, i2 string, i3 map[string]struct{}, opts ...rpc.CallOpt) (o0 map[string]struct{}, o1 []Approximation, o2 string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Relate", []interface{}{i0, i1, i2, i3}, []interface{}{&o0, &o1, &o2}, opts...)
	return
}

func (c implGroupReaderClientStub) Get(ctx *context.T, i0 GetRequest, i1 string, opts ...rpc.CallOpt) (o0 GetResponse, o1 string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Get", []interface{}{i0, i1}, []interface{}{&o0, &o1}, opts...)
	return
}

// GroupReaderServerMethods is the interface a server writer
// implements for GroupReader.
//
// GroupReader implements methods to read or query a group's membership
// information.
type GroupReaderServerMethods interface {
	// Relate determines the relationships between the provided blessing
	// names and the members of the group.
	//
	// Given an input set of blessing names and a group defined by a set of
	// blessing patterns S, for each blessing name B in the input, Relate(B)
	// returns a set of "remainders" consisting of every blessing name B"
	// such that there exists some B' for which B = B' B" and B' is in S,
	// and "" if B is a member of S.
	//
	// For example, if a group is defined as S = {n1, n1:n2, n1:n2:n3}, then
	// Relate(n1:n2) = {n2, ""}.
	//
	// reqVersion specifies the expected version of the group's membership
	// information. If this version is set and matches the Group's current
	// version, the response will indicate that fact but will otherwise be
	// empty.
	//
	// visitedGroups is the set of groups already visited in a particular
	// chain of Relate calls, and is used to detect the presence of
	// cycles. When a cycle is detected, it is treated just like any other
	// error, and the result is approximated.
	//
	// Relate also returns information about all the errors encountered that
	// resulted in approximations, if any.
	//
	// TODO(hpucha): scrub "Approximation" for preserving privacy. Flesh
	// versioning out further. Other args we may need: option to Get() the
	// membership set when allowed (to avoid an extra RPC), options related
	// to caching this information.
	Relate(_ *context.T, _ rpc.ServerCall, blessings map[string]struct{}, hint ApproximationType, reqVersion string, visitedGroups map[string]struct{}) (remainder map[string]struct{}, approximations []Approximation, version string, _ error)
	// Get returns all entries in the group.
	// TODO(sadovsky): Flesh out this API.
	Get(_ *context.T, _ rpc.ServerCall, req GetRequest, reqVersion string) (res GetResponse, version string, _ error)
}

// GroupReaderServerStubMethods is the server interface containing
// GroupReader methods, as expected by rpc.Server.
// There is no difference between this interface and GroupReaderServerMethods
// since there are no streaming methods.
type GroupReaderServerStubMethods GroupReaderServerMethods

// GroupReaderServerStub adds universal methods to GroupReaderServerStubMethods.
type GroupReaderServerStub interface {
	GroupReaderServerStubMethods
	// DescribeInterfaces the GroupReader interfaces.
	Describe__() []rpc.InterfaceDesc
}

// GroupReaderServer returns a server stub for GroupReader.
// It converts an implementation of GroupReaderServerMethods into
// an object that may be used by rpc.Server.
func GroupReaderServer(impl GroupReaderServerMethods) GroupReaderServerStub {
	stub := implGroupReaderServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implGroupReaderServerStub struct {
	impl GroupReaderServerMethods
	gs   *rpc.GlobState
}

func (s implGroupReaderServerStub) Relate(ctx *context.T, call rpc.ServerCall, i0 map[string]struct{}, i1 ApproximationType, i2 string, i3 map[string]struct{}) (map[string]struct{}, []Approximation, string, error) {
	return s.impl.Relate(ctx, call, i0, i1, i2, i3)
}

func (s implGroupReaderServerStub) Get(ctx *context.T, call rpc.ServerCall, i0 GetRequest, i1 string) (GetResponse, string, error) {
	return s.impl.Get(ctx, call, i0, i1)
}

func (s implGroupReaderServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implGroupReaderServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{GroupReaderDesc}
}

// GroupReaderDesc describes the GroupReader interface.
var GroupReaderDesc rpc.InterfaceDesc = descGroupReader

// descGroupReader hides the desc to keep godoc clean.
var descGroupReader = rpc.InterfaceDesc{
	Name:    "GroupReader",
	PkgPath: "v.io/v23/services/groups",
	Doc:     "// GroupReader implements methods to read or query a group's membership\n// information.",
	Methods: []rpc.MethodDesc{
		{
			Name: "Relate",
			Doc:  "// Relate determines the relationships between the provided blessing\n// names and the members of the group.\n//\n// Given an input set of blessing names and a group defined by a set of\n// blessing patterns S, for each blessing name B in the input, Relate(B)\n// returns a set of \"remainders\" consisting of every blessing name B\"\n// such that there exists some B' for which B = B' B\" and B' is in S,\n// and \"\" if B is a member of S.\n//\n// For example, if a group is defined as S = {n1, n1:n2, n1:n2:n3}, then\n// Relate(n1:n2) = {n2, \"\"}.\n//\n// reqVersion specifies the expected version of the group's membership\n// information. If this version is set and matches the Group's current\n// version, the response will indicate that fact but will otherwise be\n// empty.\n//\n// visitedGroups is the set of groups already visited in a particular\n// chain of Relate calls, and is used to detect the presence of\n// cycles. When a cycle is detected, it is treated just like any other\n// error, and the result is approximated.\n//\n// Relate also returns information about all the errors encountered that\n// resulted in approximations, if any.\n//\n// TODO(hpucha): scrub \"Approximation\" for preserving privacy. Flesh\n// versioning out further. Other args we may need: option to Get() the\n// membership set when allowed (to avoid an extra RPC), options related\n// to caching this information.",
			InArgs: []rpc.ArgDesc{
				{Name: "blessings", Doc: ``},     // map[string]struct{}
				{Name: "hint", Doc: ``},          // ApproximationType
				{Name: "reqVersion", Doc: ``},    // string
				{Name: "visitedGroups", Doc: ``}, // map[string]struct{}
			},
			OutArgs: []rpc.ArgDesc{
				{Name: "remainder", Doc: ``},      // map[string]struct{}
				{Name: "approximations", Doc: ``}, // []Approximation
				{Name: "version", Doc: ``},        // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Resolve"))},
		},
		{
			Name: "Get",
			Doc:  "// Get returns all entries in the group.\n// TODO(sadovsky): Flesh out this API.",
			InArgs: []rpc.ArgDesc{
				{Name: "req", Doc: ``},        // GetRequest
				{Name: "reqVersion", Doc: ``}, // string
			},
			OutArgs: []rpc.ArgDesc{
				{Name: "res", Doc: ``},     // GetResponse
				{Name: "version", Doc: ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
	},
}

// GroupClientMethods is the client interface
// containing Group methods.
//
// A group's version covers its Permissions as well as any other data stored in
// the group. Clients should treat versions as opaque identifiers. For both Get
// and Relate, if version is set and matches the Group's current version, the
// response will indicate that fact but will otherwise be empty.
type GroupClientMethods interface {
	// GroupReader implements methods to read or query a group's membership
	// information.
	GroupReaderClientMethods
	// Object provides access control for Vanadium objects.
	//
	// Vanadium services implementing dynamic access control would typically embed
	// this interface and tag additional methods defined by the service with one of
	// Admin, Read, Write, Resolve etc. For example, the VDL definition of the
	// object would be:
	//
	//   package mypackage
	//
	//   import "v.io/v23/security/access"
	//   import "v.io/v23/services/permissions"
	//
	//   type MyObject interface {
	//     permissions.Object
	//     MyRead() (string, error) {access.Read}
	//     MyWrite(string) error    {access.Write}
	//   }
	//
	// If the set of pre-defined tags is insufficient, services may define their
	// own tag type and annotate all methods with this new type.
	//
	// Instead of embedding this Object interface, define SetPermissions and
	// GetPermissions in their own interface. Authorization policies will typically
	// respect annotations of a single type. For example, the VDL definition of an
	// object would be:
	//
	//  package mypackage
	//
	//  import "v.io/v23/security/access"
	//
	//  type MyTag string
	//
	//  const (
	//    Blue = MyTag("Blue")
	//    Red  = MyTag("Red")
	//  )
	//
	//  type MyObject interface {
	//    MyMethod() (string, error) {Blue}
	//
	//    // Allow clients to change access via the access.Object interface:
	//    SetPermissions(perms access.Permissions, version string) error         {Red}
	//    GetPermissions() (perms access.Permissions, version string, err error) {Blue}
	//  }
	permissions.ObjectClientMethods
	// Create creates a new group if it doesn't already exist.
	// If perms is nil, a default Permissions is used, providing Admin access to
	// the caller.
	// Create requires the caller to have Write permission at the GroupServer.
	Create(_ *context.T, perms access.Permissions, entries []BlessingPatternChunk, _ ...rpc.CallOpt) error
	// Delete deletes the group.
	// Permissions for all group-related methods except Create() are checked
	// against the Group object.
	Delete(_ *context.T, version string, _ ...rpc.CallOpt) error
	// Add adds an entry to the group.
	Add(_ *context.T, entry BlessingPatternChunk, version string, _ ...rpc.CallOpt) error
	// Remove removes an entry from the group.
	Remove(_ *context.T, entry BlessingPatternChunk, version string, _ ...rpc.CallOpt) error
}

// GroupClientStub embeds GroupClientMethods and is a
// placeholder for additional management operations.
type GroupClientStub interface {
	GroupClientMethods
}

// GroupClient returns a client stub for Group.
func GroupClient(name string) GroupClientStub {
	return implGroupClientStub{name, GroupReaderClient(name), permissions.ObjectClient(name)}
}

type implGroupClientStub struct {
	name string

	GroupReaderClientStub
	permissions.ObjectClientStub
}

func (c implGroupClientStub) Create(ctx *context.T, i0 access.Permissions, i1 []BlessingPatternChunk, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Create", []interface{}{i0, i1}, nil, opts...)
	return
}

func (c implGroupClientStub) Delete(ctx *context.T, i0 string, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Delete", []interface{}{i0}, nil, opts...)
	return
}

func (c implGroupClientStub) Add(ctx *context.T, i0 BlessingPatternChunk, i1 string, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Add", []interface{}{i0, i1}, nil, opts...)
	return
}

func (c implGroupClientStub) Remove(ctx *context.T, i0 BlessingPatternChunk, i1 string, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Remove", []interface{}{i0, i1}, nil, opts...)
	return
}

// GroupServerMethods is the interface a server writer
// implements for Group.
//
// A group's version covers its Permissions as well as any other data stored in
// the group. Clients should treat versions as opaque identifiers. For both Get
// and Relate, if version is set and matches the Group's current version, the
// response will indicate that fact but will otherwise be empty.
type GroupServerMethods interface {
	// GroupReader implements methods to read or query a group's membership
	// information.
	GroupReaderServerMethods
	// Object provides access control for Vanadium objects.
	//
	// Vanadium services implementing dynamic access control would typically embed
	// this interface and tag additional methods defined by the service with one of
	// Admin, Read, Write, Resolve etc. For example, the VDL definition of the
	// object would be:
	//
	//   package mypackage
	//
	//   import "v.io/v23/security/access"
	//   import "v.io/v23/services/permissions"
	//
	//   type MyObject interface {
	//     permissions.Object
	//     MyRead() (string, error) {access.Read}
	//     MyWrite(string) error    {access.Write}
	//   }
	//
	// If the set of pre-defined tags is insufficient, services may define their
	// own tag type and annotate all methods with this new type.
	//
	// Instead of embedding this Object interface, define SetPermissions and
	// GetPermissions in their own interface. Authorization policies will typically
	// respect annotations of a single type. For example, the VDL definition of an
	// object would be:
	//
	//  package mypackage
	//
	//  import "v.io/v23/security/access"
	//
	//  type MyTag string
	//
	//  const (
	//    Blue = MyTag("Blue")
	//    Red  = MyTag("Red")
	//  )
	//
	//  type MyObject interface {
	//    MyMethod() (string, error) {Blue}
	//
	//    // Allow clients to change access via the access.Object interface:
	//    SetPermissions(perms access.Permissions, version string) error         {Red}
	//    GetPermissions() (perms access.Permissions, version string, err error) {Blue}
	//  }
	permissions.ObjectServerMethods
	// Create creates a new group if it doesn't already exist.
	// If perms is nil, a default Permissions is used, providing Admin access to
	// the caller.
	// Create requires the caller to have Write permission at the GroupServer.
	Create(_ *context.T, _ rpc.ServerCall, perms access.Permissions, entries []BlessingPatternChunk) error
	// Delete deletes the group.
	// Permissions for all group-related methods except Create() are checked
	// against the Group object.
	Delete(_ *context.T, _ rpc.ServerCall, version string) error
	// Add adds an entry to the group.
	Add(_ *context.T, _ rpc.ServerCall, entry BlessingPatternChunk, version string) error
	// Remove removes an entry from the group.
	Remove(_ *context.T, _ rpc.ServerCall, entry BlessingPatternChunk, version string) error
}

// GroupServerStubMethods is the server interface containing
// Group methods, as expected by rpc.Server.
// There is no difference between this interface and GroupServerMethods
// since there are no streaming methods.
type GroupServerStubMethods GroupServerMethods

// GroupServerStub adds universal methods to GroupServerStubMethods.
type GroupServerStub interface {
	GroupServerStubMethods
	// DescribeInterfaces the Group interfaces.
	Describe__() []rpc.InterfaceDesc
}

// GroupServer returns a server stub for Group.
// It converts an implementation of GroupServerMethods into
// an object that may be used by rpc.Server.
func GroupServer(impl GroupServerMethods) GroupServerStub {
	stub := implGroupServerStub{
		impl:                  impl,
		GroupReaderServerStub: GroupReaderServer(impl),
		ObjectServerStub:      permissions.ObjectServer(impl),
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implGroupServerStub struct {
	impl GroupServerMethods
	GroupReaderServerStub
	permissions.ObjectServerStub
	gs *rpc.GlobState
}

func (s implGroupServerStub) Create(ctx *context.T, call rpc.ServerCall, i0 access.Permissions, i1 []BlessingPatternChunk) error {
	return s.impl.Create(ctx, call, i0, i1)
}

func (s implGroupServerStub) Delete(ctx *context.T, call rpc.ServerCall, i0 string) error {
	return s.impl.Delete(ctx, call, i0)
}

func (s implGroupServerStub) Add(ctx *context.T, call rpc.ServerCall, i0 BlessingPatternChunk, i1 string) error {
	return s.impl.Add(ctx, call, i0, i1)
}

func (s implGroupServerStub) Remove(ctx *context.T, call rpc.ServerCall, i0 BlessingPatternChunk, i1 string) error {
	return s.impl.Remove(ctx, call, i0, i1)
}

func (s implGroupServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implGroupServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{GroupDesc, GroupReaderDesc, permissions.ObjectDesc}
}

// GroupDesc describes the Group interface.
var GroupDesc rpc.InterfaceDesc = descGroup

// descGroup hides the desc to keep godoc clean.
var descGroup = rpc.InterfaceDesc{
	Name:    "Group",
	PkgPath: "v.io/v23/services/groups",
	Doc:     "// A group's version covers its Permissions as well as any other data stored in\n// the group. Clients should treat versions as opaque identifiers. For both Get\n// and Relate, if version is set and matches the Group's current version, the\n// response will indicate that fact but will otherwise be empty.",
	Embeds: []rpc.EmbedDesc{
		{Name: "GroupReader", PkgPath: "v.io/v23/services/groups", Doc: "// GroupReader implements methods to read or query a group's membership\n// information."},
		{Name: "Object", PkgPath: "v.io/v23/services/permissions", Doc: "// Object provides access control for Vanadium objects.\n//\n// Vanadium services implementing dynamic access control would typically embed\n// this interface and tag additional methods defined by the service with one of\n// Admin, Read, Write, Resolve etc. For example, the VDL definition of the\n// object would be:\n//\n//   package mypackage\n//\n//   import \"v.io/v23/security/access\"\n//   import \"v.io/v23/services/permissions\"\n//\n//   type MyObject interface {\n//     permissions.Object\n//     MyRead() (string, error) {access.Read}\n//     MyWrite(string) error    {access.Write}\n//   }\n//\n// If the set of pre-defined tags is insufficient, services may define their\n// own tag type and annotate all methods with this new type.\n//\n// Instead of embedding this Object interface, define SetPermissions and\n// GetPermissions in their own interface. Authorization policies will typically\n// respect annotations of a single type. For example, the VDL definition of an\n// object would be:\n//\n//  package mypackage\n//\n//  import \"v.io/v23/security/access\"\n//\n//  type MyTag string\n//\n//  const (\n//    Blue = MyTag(\"Blue\")\n//    Red  = MyTag(\"Red\")\n//  )\n//\n//  type MyObject interface {\n//    MyMethod() (string, error) {Blue}\n//\n//    // Allow clients to change access via the access.Object interface:\n//    SetPermissions(perms access.Permissions, version string) error         {Red}\n//    GetPermissions() (perms access.Permissions, version string, err error) {Blue}\n//  }"},
	},
	Methods: []rpc.MethodDesc{
		{
			Name: "Create",
			Doc:  "// Create creates a new group if it doesn't already exist.\n// If perms is nil, a default Permissions is used, providing Admin access to\n// the caller.\n// Create requires the caller to have Write permission at the GroupServer.",
			InArgs: []rpc.ArgDesc{
				{Name: "perms", Doc: ``},   // access.Permissions
				{Name: "entries", Doc: ``}, // []BlessingPatternChunk
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
		{
			Name: "Delete",
			Doc:  "// Delete deletes the group.\n// Permissions for all group-related methods except Create() are checked\n// against the Group object.",
			InArgs: []rpc.ArgDesc{
				{Name: "version", Doc: ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
		{
			Name: "Add",
			Doc:  "// Add adds an entry to the group.",
			InArgs: []rpc.ArgDesc{
				{Name: "entry", Doc: ``},   // BlessingPatternChunk
				{Name: "version", Doc: ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
		{
			Name: "Remove",
			Doc:  "// Remove removes an entry from the group.",
			InArgs: []rpc.ArgDesc{
				{Name: "entry", Doc: ``},   // BlessingPatternChunk
				{Name: "version", Doc: ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
	},
}

// Hold type definitions in package-level variables, for better performance.
//nolint:unused
var (
	vdlTypeString1 *vdl.Type
	vdlTypeStruct2 *vdl.Type
	vdlTypeStruct3 *vdl.Type
	vdlTypeSet4    *vdl.Type
	vdlTypeEnum5   *vdl.Type
	vdlTypeStruct6 *vdl.Type
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
	vdl.Register((*BlessingPatternChunk)(nil))
	vdl.Register((*GetRequest)(nil))
	vdl.Register((*GetResponse)(nil))
	vdl.Register((*ApproximationType)(nil))
	vdl.Register((*Approximation)(nil))

	// Initialize type definitions.
	vdlTypeString1 = vdl.TypeOf((*BlessingPatternChunk)(nil))
	vdlTypeStruct2 = vdl.TypeOf((*GetRequest)(nil)).Elem()
	vdlTypeStruct3 = vdl.TypeOf((*GetResponse)(nil)).Elem()
	vdlTypeSet4 = vdl.TypeOf((*map[BlessingPatternChunk]struct{})(nil))
	vdlTypeEnum5 = vdl.TypeOf((*ApproximationType)(nil))
	vdlTypeStruct6 = vdl.TypeOf((*Approximation)(nil)).Elem()

	// Set error format strings.
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoBlessings.ID), "{1:}{2:} No blessings recognized; cannot create group Permissions")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrExcessiveContention.ID), "{1:}{2:} Gave up after encountering excessive contention; try again later")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrCycleFound.ID), "{1:}{2:} Found cycle in group definitions {3} visited {4}")

	return struct{}{}
}
