// Copyright 2022 Giuseppe Calabrese.
// This file is distributed under the terms of the ISC License.

package strong_test

import (
	"fmt"
	"math/rand"
	"time"
)

// Create a source.
func ExampleNewSource_initialization() {
	src := strong.NewSource()

	// A source supports only two operations.
	_ = src.Int63()  // as per math/rand.Source
	_ = src.Uint64() // as per math/rand.Source64
}

// Use a source with the [math/rand] package.
func ExampleNewSource_expectedUse() {
	// Initialize a Rand.
	rng := rand.New(strong.NewSource())

	// Use it to generate all variates supported by math/rand.
	_ = 1 + rng.Intn(6)
	_ = 1 + rng.Float64()*10
	_ = rng.ExpFloat64() * 1000
	_ = 100 + rng.NormFloat64()*15
}

func Example_bernoulliVariate() {
	rng := rand.New(strong.NewSource())

	// A good Bernoulli variate.
	B := func(p float64) bool {
		return rng.Float64() < p
	}

	// Russian roulette.
	for B(0.9) {
		fmt.Print("Click. ")
		time.Sleep(time.Second)
	}
	fmt.Println("\n\nGame over.")
}

func Example_concurrency() {
	src := strong.NewSource() // Safe for concurrent use.

	// This is OK:
	go src.Uint64()
	go src.Int63()
	go src.Int63()
	go src.Uint64()

	rng := rand.New(src) // NOT safe for concurrent use.

	// DON'T:
	go rng.ExpFloat64()
	go rng.NormFloat64()

	// OK to reuse the same source.
	go rand.New(src).ExpFloat64()
	go rand.New(src).NormFloat64()
}
