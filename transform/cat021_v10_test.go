package transform

import (
	"math"
	"testing"
)

func equalWithinErrorBounds(actualValue float64, targetValue float64, epsilon float64) bool {
	return math.Abs(targetValue-actualValue) < epsilon
}

func checkEqualLatLong(resultCoordinates WGS84Coordinates, actualCoordinates WGS84Coordinates, epsilon float64) bool {
	compareLatitudes := equalWithinErrorBounds(float64(resultCoordinates.Latitude), float64(actualCoordinates.Latitude), epsilon)
	compareLongitudes := equalWithinErrorBounds(float64(resultCoordinates.Longitude), float64(actualCoordinates.Longitude), epsilon)
	return compareLatitudes && compareLongitudes
}

func Test_wgs84Coordinates_LowPrecision(t *testing.T) {
	// Arrange
	input := []byte{0x24, 0x0, 0x0, 0x07, 0x00, 0x00}
	output := WGS84Coordinates{
		Latitude:  51.0,
		Longitude: 10.0,
	}
	epsilon := 1.0

	// Act
	res := wgs84Coordinates(input)

	// Assert
	if !checkEqualLatLong(res, output, epsilon) {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_wgs84Coordinates_HighPrecision(t *testing.T) {
	// Arrange
	input := []byte{0x24, 0x00, 0x0, 0x0, 0x07, 0x00, 0x00, 0x0}
	output := WGS84Coordinates{
		Latitude:  51.0,
		Longitude: 10.0,
	}
	epsilon := 1.0

	// Act
	res := wgs84Coordinates(input)

	// Assert
	if !checkEqualLatLong(res, output, epsilon) {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_BasicGeometricHeight(t *testing.T){
	// Arrange 
	input := [2]byte{0x09, 0x60}
	output := GeometricHeight{
		Height: 15000.0,
		GreaterThan: false,
	}

	// Act
	res := getGeometricHeight(input)

	// Assert
	if res != output {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_MaxGeometricHeight(t *testing.T){
	// Arrange 
	input := [2]byte{0x5D, 0xC0}
	output := GeometricHeight{
		Height: 150000.0,
		GreaterThan: false,
	}

	// Act
	res := getGeometricHeight(input)

	// Assert
	if res != output {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_MinGeometricHeight(t *testing.T){
	// Arrange 
	input := [2]byte{0xFF, 0x10}
	output := GeometricHeight{
		Height: -1500.0,
		GreaterThan: false,
	}

	// Act
	res := getGeometricHeight(input)

	// Assert
	if res != output {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_GreaterThanGeometricHeight(t *testing.T){
	// Arrange 
	input := [2]byte{0x7F, 0xFF}
	output := GeometricHeight{
		Height: 204793.75,
		GreaterThan: true,
	}

	// Act
	res := getGeometricHeight(input)

	// Assert
	if res != output {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}