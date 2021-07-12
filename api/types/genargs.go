// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"strconv"
	"strings"
)

// GenArgs is a facade over GeneratorArgs, exposing a few readonly properties.
type GenArgs struct {
	Args *GeneratorArgs
}

// NewGenArgs returns a new instance of GenArgs.
func NewGenArgs(args *GeneratorArgs) *GenArgs {
	return &GenArgs{Args: args}
}

func (g *GenArgs) String() string {
	if g == nil {
		return "{nilGenArgs}"
	}
	return "{" +
		strings.Join([]string{
			"nsfx:" + strconv.FormatBool(g.ShouldAddHashSuffixToName()),
			"beh:" + g.Behavior().String()},
			",") +
		"}"
}

// ShouldAddHashSuffixToName returns true if a resource
// content hash should be appended to the name of the resource.
func (g *GenArgs) ShouldAddHashSuffixToName() bool {
	return g.Args != nil &&
		(g.Args.Options == nil || !g.Args.Options.DisableNameSuffixHash)
}

// Behavior returns Behavior field of GeneratorArgs
func (g *GenArgs) Behavior() GenerationBehavior {
	if g == nil || g.Args == nil {
		return BehaviorUnspecified
	}
	return NewGenerationBehavior(g.Args.Behavior)
}
