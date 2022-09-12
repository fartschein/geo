// Copyright 2021 Hugo Melder. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package geo

import (
	"fmt"
	"testing"
)

const (
  fixedLatitudeDegrees = 52.5185931
  fixedLongitudeDegrees = 13.3761064

  secondFixedLatitudeDegrees = 52.5204445
  secondFixedLongitudeDegrees = 13.4069693

  earthRadiumInKm = 6371
  distanceInKm = 2.0983416370739216

  firstPointLatitude = 52.519211125280364
  firstPointLongitude = 13.386393744246242
  secondPointLatitude = 52.51982825862868
  secondPointLongitude = 13.396681377713959

  fixedLatitudeRadians = 0.9166223681101755
  fixedLongitudeRadians = 0.23345709777708562 
)

func validateCoordinates(in *Coordinates) bool {
  switch in.Type {
  case CoordinatesRadians:
    if in.Latitude == fixedLatitudeRadians &&
       in.Longitude == fixedLongitudeRadians {
      return true
    }
  case CoordinatesDegrees:
    if in.Latitude == fixedLatitudeDegrees &&
       in.Longitude == fixedLongitudeDegrees {
      return true
    }
  }

  return false
}

func TestConvertToRadians(t *testing.T) {
  coordDeg := &Coordinates{
    Latitude: fixedLatitudeDegrees,
    Longitude: fixedLongitudeDegrees,
    Type: CoordinatesDegrees,
  }

  coordRad := ConvertToRadians(coordDeg)
  if validateCoordinates(coordRad) != true {
    t.Errorf("Unexpected result")
  }
}

func TestConvertToDegrees(t *testing.T) {
  coordRad := &Coordinates{
    Latitude: fixedLatitudeRadians,
    Longitude: fixedLongitudeRadians,
    Type: CoordinatesRadians,
  }

  coordDeg := ConvertToDegrees(coordRad)
  if validateCoordinates(coordDeg) != true {
    t.Errorf("Unexpected result")
  }
}

func TestGreatCircleDistance(t *testing.T) {
  coordDegOne := &Coordinates{
    Latitude: fixedLatitudeDegrees,
    Longitude: fixedLongitudeDegrees,
    Type: CoordinatesDegrees,
  }
  coordDegTwo := &Coordinates{
    Latitude: secondFixedLatitudeDegrees,
    Longitude: secondFixedLongitudeDegrees,
    Type: CoordinatesDegrees,
  }

  distance := GreatCircleDistance(coordDegOne, coordDegTwo, earthRadiumInKm)
  if distance != distanceInKm {
    t.Errorf("Unexpected result, got %f expected %f", distance, distanceInKm)
  }
}

func TestGreatCircleIntermediate(t *testing.T) {
  coordDegOne := &Coordinates{
    Latitude: fixedLatitudeDegrees,
    Longitude: fixedLongitudeDegrees,
    Type: CoordinatesDegrees,
  }
  coordDegTwo := &Coordinates{
    Latitude: secondFixedLatitudeDegrees,
    Longitude: secondFixedLongitudeDegrees,
    Type: CoordinatesDegrees,
  }

  points := GreatCircleIntermediate(coordDegOne, coordDegTwo, 2)

  if (*points)[0].Latitude != firstPointLatitude ||
     (*points)[0].Longitude != firstPointLongitude ||
     (*points)[1].Latitude != secondPointLatitude ||
     (*points)[1].Longitude != secondPointLongitude {
    t.Errorf("Unexpected result, points do not match fixed points")
  }
}

func TestCoordinateDisplacement(t *testing.T) {
  coordDeg := &Coordinates{
    Latitude: fixedLatitudeDegrees,
    Longitude: fixedLongitudeDegrees,
    Type: CoordinatesDegrees,
  }

  // Apply a 10m offset
  coords := CoordinateDisplacement(coordDeg, 10)
  fmt.Printf("Original coords: %+v\n Displaced coords (10m): %+v\n", coordDeg, coords)
}
