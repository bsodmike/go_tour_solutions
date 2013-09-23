package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := float64(1)
  prevZ := 0.0
  delta := 1.0
  i := 0
  for i < 10 {
    prevZ = z
    z = z - (((z*z)-x)/(2*z))
    if i > 0 {
      delta = prevZ - z
      fmt.Printf("---> Delta: %v\n", delta)
    }
    fmt.Printf("#%v - Approx: %v (Target delta: %v)\n", i, z, z - math.Sqrt(2))
    i++
    if delta >= 0 && delta <= 0.00001 { break }
  }
  return z
}

func main() {
  fmt.Printf("Target sqrt(2): %v\n", math.Sqrt(2))
  fmt.Println(Sqrt(2))
}
