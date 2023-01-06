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
