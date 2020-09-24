// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: stats

// Package stats defines an interface to access statistical information for
// troubleshooting and monitoring purposes.
//nolint:golint
package stats

import (
	"fmt"

	v23 "v.io/v23"
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/rpc"
	"v.io/v23/security/access"
	"v.io/v23/services/watch"
	"v.io/v23/vdl"
	"v.io/v23/verror"
	"v.io/v23/vom"
)

var _ = initializeVDL() // Must be first; see initializeVDL comments for details.

//////////////////////////////////////////////////
// Error definitions

var (
	ErrNoValue = verror.NewIDAction("v.io/v23/services/stats.NoValue", verror.NoRetry)
)

// NewErrNoValue returns an error with the ErrNoValue ID.
// WARNING: this function is deprecated and will be removed in the future,
// use ErrorfErrNoValue or MessageErrNoValue instead.
func NewErrNoValue(ctx *context.T, suffix string) error {
	return verror.New(ErrNoValue, ctx, suffix)
}

// ErrorfErrNoValue calls ErrNoValue.Errorf with the supplied arguments.
func ErrorfErrNoValue(ctx *context.T, format string, suffix string) error {
	return ErrNoValue.Errorf(ctx, format, suffix)
}

// MessageErrNoValue calls ErrNoValue.Message with the supplied arguments.
func MessageErrNoValue(ctx *context.T, message string, suffix string) error {
	return ErrNoValue.Message(ctx, message, suffix)
}

// ParamsErrNoValue extracts the expected parameters from the error's ParameterList.
func ParamsErrNoValue(argumentError error) (verrorComponent string, verrorOperation string, suffix string, returnErr error) {
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
	if suffix, ok = tmp.(string); !ok {
		if returnErr != nil {
			return
		}
		returnErr = fmt.Errorf("parameter list contains the wrong type for return value suffix, has %T and not string", tmp)
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

// StatsClientMethods is the client interface
// containing Stats methods.
//
// The Stats interface is used to access stats for troubleshooting and
// monitoring purposes. The stats objects are discoverable via the Globbable
// interface and watchable via the GlobWatcher interface.
//
// The types of the object values are implementation specific, but should be
// primarily numeric in nature, e.g. counters, memory usage, latency metrics,
// etc.
type StatsClientMethods interface {
	// GlobWatcher allows a client to receive updates for changes to objects
	// that match a pattern.  See the package comments for details.
	watch.GlobWatcherClientMethods
	// Value returns the current value of an object, or an error. The type
	// of the value is implementation specific.
	// Some objects may not have a value, in which case, Value() returns
	// a NoValue error.
	Value(*context.T, ...rpc.CallOpt) (*vom.RawBytes, error)
}

// StatsClientStub embeds StatsClientMethods and is a
// placeholder for additional management operations.
type StatsClientStub interface {
	StatsClientMethods
}

// StatsClient returns a client stub for Stats.
func StatsClient(name string) StatsClientStub {
	return implStatsClientStub{name, watch.GlobWatcherClient(name)}
}

type implStatsClientStub struct {
	name string

	watch.GlobWatcherClientStub
}

func (c implStatsClientStub) Value(ctx *context.T, opts ...rpc.CallOpt) (o0 *vom.RawBytes, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Value", nil, []interface{}{&o0}, opts...)
	return
}

// StatsServerMethods is the interface a server writer
// implements for Stats.
//
// The Stats interface is used to access stats for troubleshooting and
// monitoring purposes. The stats objects are discoverable via the Globbable
// interface and watchable via the GlobWatcher interface.
//
// The types of the object values are implementation specific, but should be
// primarily numeric in nature, e.g. counters, memory usage, latency metrics,
// etc.
type StatsServerMethods interface {
	// GlobWatcher allows a client to receive updates for changes to objects
	// that match a pattern.  See the package comments for details.
	watch.GlobWatcherServerMethods
	// Value returns the current value of an object, or an error. The type
	// of the value is implementation specific.
	// Some objects may not have a value, in which case, Value() returns
	// a NoValue error.
	Value(*context.T, rpc.ServerCall) (*vom.RawBytes, error)
}

// StatsServerStubMethods is the server interface containing
// Stats methods, as expected by rpc.Server.
// The only difference between this interface and StatsServerMethods
// is the streaming methods.
type StatsServerStubMethods interface {
	// GlobWatcher allows a client to receive updates for changes to objects
	// that match a pattern.  See the package comments for details.
	watch.GlobWatcherServerStubMethods
	// Value returns the current value of an object, or an error. The type
	// of the value is implementation specific.
	// Some objects may not have a value, in which case, Value() returns
	// a NoValue error.
	Value(*context.T, rpc.ServerCall) (*vom.RawBytes, error)
}

// StatsServerStub adds universal methods to StatsServerStubMethods.
type StatsServerStub interface {
	StatsServerStubMethods
	// DescribeInterfaces the Stats interfaces.
	Describe__() []rpc.InterfaceDesc
}

// StatsServer returns a server stub for Stats.
// It converts an implementation of StatsServerMethods into
// an object that may be used by rpc.Server.
func StatsServer(impl StatsServerMethods) StatsServerStub {
	stub := implStatsServerStub{
		impl:                  impl,
		GlobWatcherServerStub: watch.GlobWatcherServer(impl),
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

type implStatsServerStub struct {
	impl StatsServerMethods
	watch.GlobWatcherServerStub
	gs *rpc.GlobState
}

func (s implStatsServerStub) Value(ctx *context.T, call rpc.ServerCall) (*vom.RawBytes, error) {
	return s.impl.Value(ctx, call)
}

func (s implStatsServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implStatsServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{StatsDesc, watch.GlobWatcherDesc}
}

// StatsDesc describes the Stats interface.
var StatsDesc rpc.InterfaceDesc = descStats

// descStats hides the desc to keep godoc clean.
var descStats = rpc.InterfaceDesc{
	Name:    "Stats",
	PkgPath: "v.io/v23/services/stats",
	Doc:     "// The Stats interface is used to access stats for troubleshooting and\n// monitoring purposes. The stats objects are discoverable via the Globbable\n// interface and watchable via the GlobWatcher interface.\n//\n// The types of the object values are implementation specific, but should be\n// primarily numeric in nature, e.g. counters, memory usage, latency metrics,\n// etc.",
	Embeds: []rpc.EmbedDesc{
		{Name: "GlobWatcher", PkgPath: "v.io/v23/services/watch", Doc: "// GlobWatcher allows a client to receive updates for changes to objects\n// that match a pattern.  See the package comments for details."},
	},
	Methods: []rpc.MethodDesc{
		{
			Name: "Value",
			Doc:  "// Value returns the current value of an object, or an error. The type\n// of the value is implementation specific.\n// Some objects may not have a value, in which case, Value() returns\n// a NoValue error.",
			OutArgs: []rpc.ArgDesc{
				{Name: "", Doc: ``}, // *vom.RawBytes
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Debug"))},
		},
	},
}

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

	// Set error format strings.
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoValue.ID), "{1:}{2:} object has no value, suffix: {3}")

	return struct{}{}
}
