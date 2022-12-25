// Copyright 2022 Giuseppe Calabrese.
// This file is distributed under the terms of the ISC License.

// Package strong implements a strong source of randomness for use with the
// [math/rand] package.
//
// The [NewSource] method returns a random source of type [math/rand.Source64]
// backed by a cryptographically-secure random number generator.
//
// The output of a [math/rand.Rand] initialized using such a source is as
// unpredictable as reasonably achievable subject to the imposed probability
// distribution and the quality of [math/rand]'s generation methods.
//
// A source returned by this package is safe for concurrent use.
// A [math/rand.Rand] initialized with it, however, is not.
//
// # Implementation notes
//
// The [math/rand.Source64.Seed] and [math/rand.Rand.Seed] methods for sources
// returned by this package are no-op.
//
// This package trades efficiency for randomness quality. It currently uses
// [crypto/rand.Reader] as its source or randomness, which on common platforms
// requires a system call.
package strong

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

// NewSource returns a [math/rand.Source64] that outputs
// cryptographically-secure uniform random integers.
func NewSource() rand.Source64 {
	return source64{}
}

// A [math/rand.Source64] backed by [crypto/rand.Reader].
type source64 struct{}

var _ rand.Source64 = source64{}

func (s source64) Seed(_ int64) {
	// Do nothing.
}

func (s source64) Uint64() uint64 {
	var bs [8]byte
	_, err := cryptorand.Read(bs[:])
	if err != nil {
		panic(fmt.Errorf(
			"secure: failed to draw entropy from rand.Reader: %w",
			err))
	}
	return binary.LittleEndian.Uint64(bs[:])
}

func (s source64) Int63() int64 {
	return int64(s.Uint64() & (1<<63 - 1))
}
