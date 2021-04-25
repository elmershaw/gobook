package main

import "fmt"

type Meter float64
type Foot float64

const lFactor float64 = 3.2808

func (m Meter) String() string { return fmt.Sprintf("%gM", m) }
func (f Foot) String() string  { return fmt.Sprintf("%gft", f) }

// FToM converts Meter to Foot
func FToM(f Foot) Meter { return Meter(f / Foot(lFactor)) }

// MToF converts Foot to Meter
func MToF(m Meter) Foot { return Foot(m * Meter(lFactor)) }
