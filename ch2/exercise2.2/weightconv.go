package main

import "fmt"

type Kilogram float64
type Pound float64

const wFactor float64 = 2.2045

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }

// KToP converts Kilogram to Pound
func KToP(k Kilogram) Pound { return Pound(k * Kilogram(wFactor)) }

// PToK converts Pound to Kilogram
func PToK(p Pound) Kilogram { return Kilogram(p / Pound(wFactor)) }
