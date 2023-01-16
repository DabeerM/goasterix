package transform

import (
	"reflect"
	"testing"

	"github.com/mokhtarimokhtar/goasterix"
)

func Test_TargetReportDescriptors(t *testing.T) {
	// Arrange 
	inputBytestream := []byte{0x2E}
	input := goasterix.Compound{
		Primary: inputBytestream,
	}
	output := TargetReportDescriptor{
		ATP: "Duplicate address",
		ARC: "100ft",
		RC: "Range Check passed, CPR Validation pending",
		RAB: "Report from field monitor (fixed transponder)",
	}

	// Act 
	res := targetReportDescriptor(input)

	// Assert
	if reflect.DeepEqual(output, res) {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
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

func Test_BasicGeometricHeight(t *testing.T) {
	// Arrange
	input := [2]byte{0x09, 0x60}
	output := GeometricHeight{
		Height:      15000.0,
		GreaterThan: false,
	}

	// Act
	res := geometricHeight(input)

	// Assert
	if res != output {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_MaxGeometricHeight(t *testing.T) {
	// Arrange
	input := [2]byte{0x5D, 0xC0}
	output := GeometricHeight{
		Height:      150000.0,
		GreaterThan: false,
	}

	// Act
	res := geometricHeight(input)

	// Assert
	if res != output {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_MinGeometricHeight(t *testing.T) {
	// Arrange
	input := [2]byte{0xFF, 0x10}
	output := GeometricHeight{
		Height:      -1500.0,
		GreaterThan: false,
	}

	// Act
	res := geometricHeight(input)

	// Assert
	if res != output {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_GreaterThanGeometricHeight(t *testing.T) {
	// Arrange
	input := [2]byte{0x7F, 0xFF}
	output := GeometricHeight{
		Height:      204793.75,
		GreaterThan: true,
	}

	// Act
	res := geometricHeight(input)

	// Assert
	if res != output {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}
