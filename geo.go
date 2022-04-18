// Copyright 2021 Hugo Melder. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package geo

import (
	"math"
)

const (
  // Coordinate type information
  CoordinatesRadians = 1
  CoordinatesDegrees = 2
)

type Coordinates struct {
  Latitude  float64
  Longitude float64

  // Coordinate type information
  Type      int
}

func ConvertToRadians(in *Coordinates) *Coordinates {
  var coordRad Coordinates

  if in.Type == CoordinatesDegrees {
    coordRad.Latitude = in.Latitude * (math.Pi/180)
    coordRad.Longitude = in.Longitude * (math.Pi/180)
    coordRad.Type = CoordinatesRadians
  
    return &coordRad
  } else {
    return in
  }
}

func ConvertToDegrees(in *Coordinates) *Coordinates {
  var coordDeg Coordinates

  if in.Type == CoordinatesRadians {
    coordDeg.Latitude = in.Latitude * (180/math.Pi)
    coordDeg.Longitude = in.Longitude * (180/math.Pi)
    coordDeg.Type = CoordinatesDegrees
  
    return &coordDeg
  } else {
    return in
  }
}


func GreatCircleDistance(p *Coordinates, q *Coordinates, r int) float64 {
  c1 := ConvertToRadians(p)
  c2 := ConvertToRadians(q)

	diffLat := c2.Latitude - c1.Latitude
	diffLon := c2.Longitude - c1.Longitude

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(c1.Latitude)*math.Cos(c2.Latitude)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return c * float64(r)
}

func GreatCircleIntermediate(p *Coordinates, q *Coordinates, n float64) *[]Coordinates {
  c1 := ConvertToRadians(p)
  c2 := ConvertToRadians(q)

  d := GreatCircleDistance(p, q, 1)

  // Create a regular sequence and divide it by max + 1 
  // The calculated coordinates are a fraction along the great circle route.
  seqCoords := func(min, max int) []Coordinates {
    a := make([]Coordinates, max-min+1)
    for i := range a {
      f := float64(min + i) / float64((max + 1))

      A := math.Sin((1-f)*d) / math.Sin(d)
      B := math.Sin(f*d) / math.Sin(d)
      x := A * math.Cos(c1.Latitude) * math.Cos(c1.Longitude) +
           B * math.Cos(c2.Latitude) * math.Cos(c2.Longitude) 
      y := A * math.Cos(c1.Latitude) * math.Sin(c1.Longitude) +
           B * math.Cos(c2.Latitude) * math.Sin(c2.Longitude)
      z := A * math.Sin(c1.Latitude) + B * math.Sin(c2.Latitude)

      a[i].Latitude = math.Atan2(z, math.Sqrt(math.Pow(x,2)+math.Pow(y,2)))
      a[i].Longitude = math.Atan2(y,x)
      a[i].Type = CoordinatesRadians
      a[i] = *ConvertToDegrees(&a[i])
    }

    return a
  }
  
  n2 := int(math.Max(math.Round(n), 1))
  s := seqCoords(1, n2)
  
  return &s
}
