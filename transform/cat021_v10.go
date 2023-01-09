package transform

import (
	"errors"
	"math"

	"github.com/mokhtarimokhtar/goasterix"
)

const (
	BYTESIZE = 8
)

var ErrTypeUnknown021 = errors.New("[ASTERIX Error CAT021] Message TYPE Unknown")

type WGS84Coordinates struct {
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}

type TargetAddress struct {
	Target  byte  `json:"target,omitempty"`
	Address int16 `json:"address,omitempty"` // TODO: Check if this is the best type?
}

type GeometricHeight struct {
	Height      float64 `json:"height,omitempty"`
	GreaterThan bool    `json:"greaterthan,omitempty"`
}

type AirSpeed struct {
	IM       string  `json:"im,omitempty"`
	AirSpeed float64 `json:"airspeed,omitempty"`
}
type TrueAirSpeed struct {
	RE    int   `json:"re,omitempty"`
	Speed int16 `json:"speed,omitempty"`
}

// TODO: Write the potential messages/states in const
// TODO: Look into next extensions ( not clearly defined by spec)
type SecondExtension struct {
	LLC            string `json:"llc,omitempty"`
	IPC            string `json:"ipc,omitempty"`
	NOGO           string `json:"nogo,omitempty"`
	CPR            string `json:"cpr,omitempty"`
	LDPJ           string `json:"ldpj,omitempty"`
	RCF            string `json:"rcf,omitempty"`
	ThirdExtension byte   `json:"fx,omitempty"`
}
type FirstExtension struct {
	DCR             string           `json:"dcr,omitempty"`
	GBS             string           `json:"gbs,omitempty"`
	SIM             string           `json:"sim,omitempty"`
	TST             string           `json:"tst,omitempty"`
	SAA             string           `json:"saa,omitempty"`
	CL              string           `json:"cl,omitempty"`
	SecondExtension *SecondExtension `json:"fx,omitempty"`
}
type TargetReportDescriptor struct {
	ATP string          `json:"atp,omitempty"`
	ARC string          `json:"arc,omitempty"`
	RC  string          `json:"rc,omitempty"`
	RAB string          `json:"rab,omitempty"`
	FX  *FirstExtension `json:"fx,omitempty"`
}

type Cat021Model struct {
	AircraftOperationStatus                        string                  `json:"aircraftOperationStatus,omitempty"`
	DataSourceIdentification                       *SourceIdentifier       `json:"DataSourceIdentification,omitempty"`
	ServiceIdentification                          byte                    `json:"ServiceIdentification,omitempty"`
	ServiceManagement                              string                  `json:"ServiceManagement,omitempty"`
	EmitterCategory                                string                  `json:"EmitterCategory,omitempty"`
	TargetReportDescriptor                         *TargetReportDescriptor `json:"TargetReportDescriptor,omitempty"`
	Mode3ACode                                     string                  `json:"Mode3ACode,omitempty"`
	TimeOfApplicabilityForPosition                 float64                 `json:"timeOfApplicabilityForPosition,omitempty"`
	TimeOfApplicabilityForVelocity                 float64                 `json:"timeOfApplicabilityForVelocity,omitempty"`
	TimeOfMessageReceptionForPosition              float64                 `json:"TimeOfMessageReceptionForPosition,omitempty"`
	TimeOfMessageReceptionForPositionHighPrecision *TimeOfDayHighPrecision `json:"TimeOfMessageReceptionForPositionHighPrecision,omitempty"`
	TimeOfMessageReceptionForVelocity              float64                 `json:"TimeOfMessageReceptionForVelocity,omitempty"`
	TimeOfMessageReceptionForVelocityHighPrecision *TimeOfDayHighPrecision `json:"TimeOfMessageReceptionForVelocityHighPrecision,omitempty"`
	TimeOfReportTransmission                       float64                 `json:"TimeOfReportTransmission,omitempty"`
	TargetAddress                                  *TargetAddress          `json:"TargetAddress,omitempty"`
	QualityIndicators                              string                  `json:"QualityIndicators,omitempty"`
	TrajectoryIntent                               string                  `json:"TrajectoryIntent,omitempty"`
	PositionWGS84                                  *WGS84Coordinates       `json:"PositionWGS84,omitempty"`
	PositionWGS84HighRes                           *WGS84Coordinates       `json:"PositionWGS84HighRes,omitempty"`
	MessageAmplitude                               int64                   `json:"MessageAmplitude,omitempty"`
	GeometricHeight                                *GeometricHeight         `json:"GeometricHeight,omitempty"`
	FlightLevel                                    float64                 `json:"FlightLevel,omitempty"`
	SelectedAltitude                               int64                   `json:"SelectedAltitude,omitempty"`
	FinalStateSelectedAltitude                     int64                   `json:"FinalStateSelectedAltitude,omitempty"`
	AirSpeed                                       *AirSpeed               `json:"AirSpeed,omitempty"`
	TrueAirSpeed                                   *TrueAirSpeed           `json:"TrueAirSpeed,omitempty"`
	MagneticHeading                                float64                 `json:"MagneticHeading,omitempty"`
	BarometricVerticalRate                         float64                 `json:"BarometricVerticalRate,omitempty"`
	GeometricVerticalRate                          float64                 `json:"GeometricVerticalRate,omitempty"`
	AirborneGroundVector                           string                  `json:"AirborneGroundVector,omitempty"`
	TrackNumber                                    uint16                  `json:"TrackNumber,omitempty"`
	TrackAngleRate                                 float64                 `json:"TrackAngleRate,omitempty"`
	TargetIdentification                           string                  `json:"TargetIdentification,omitempty"`
	TargetStatus                                   string                  `json:"TargetStatus,omitempty"`
	MPOSVersion                                    string                  `json:"MPOSVersion,omitempty"`
	MetInformation                                 string                  `json:"MetInformation,omitempty"`
	RollAngle                                      float64                 `json:"RollAngle,omitempty"`
	ModeSMBData                                    string                  `json:"ModeSMBData,omitempty"`
	ACASResolutionAdvisoryReport                   string                  `json:"ACASResolutionAdvisoryReport,omitempty"`
	ReceiverID                                     string                  `json:"ReceiverID,omitempty"`
}

func (data *Cat021Model) write(rec goasterix.Record) {
	for _, item := range rec.Items {
		switch item.Meta.FRN {
		case 1:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp, _ := sacSic(payload)
			data.DataSourceIdentification = &tmp
		case 2:
			// TODO: Write unit tests
			tmp := getTargetReportDescriptor(*item.Compound)
			data.TargetReportDescriptor = &tmp
		case 3:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := trackNumber(payload)
			data.TrackNumber = tmp
		case 4:
			data.ServiceIdentification = item.Fixed.Data[0] // TODO: Double check?
		case 5:
			// TODO: Check correctness
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TimeOfApplicabilityForPosition, _ = timeOfDay(payload)
		case 6:
			var payload []byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := wgs84Coordinates(payload)
			data.PositionWGS84 = &tmp
		case 7:
			var payload []byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := wgs84Coordinates(payload)
			data.PositionWGS84 = &tmp
		case 8:
			// TODO: Check correctness
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TimeOfApplicabilityForVelocity, _ = timeOfDay(payload)
		case 9:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := getAirSpeed(payload)
			data.AirSpeed = &tmp
		case 10:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := getTrueAirSpeed(payload)
			data.TrueAirSpeed = &tmp
		case 11:
			var tmp TargetAddress
			tmp.Target = item.Fixed.Data[0]
			tmp.Address = int16(item.Fixed.Data[1]) + int16(item.Fixed.Data[2])
			data.TargetAddress = &tmp
		case 12:
			// TODO: Check correctness
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TimeOfMessageReceptionForPosition, _ = timeOfDay(payload)
		case 13:
			// TODO: Check correctness
			var payload [4]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp, _ := timeOfDayHighPrecision(payload)
			data.TimeOfMessageReceptionForPositionHighPrecision = &tmp
		case 14:
			// TODO: Check correctness
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TimeOfMessageReceptionForVelocity, _ = timeOfDay(payload)
		case 15:
			// TODO: Check correctness
			var payload [4]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp, _ := timeOfDayHighPrecision(payload)
			data.TimeOfMessageReceptionForVelocityHighPrecision = &tmp
		case 16:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := getGeometricHeight(payload)
			data.GeometricHeight = &tmp
		case 17:
			// Do stuff
		case 18:
			// Do stuff
		case 19:
			// Do stuff
		case 20:
			// Do stuff
		case 21:
			// Do stuff
		case 22:
			// Do stuff
		case 23:
			// Do stuff
		case 24:
			// Do stuff
		case 25:
			// Do stuff
		case 26:
			// Do stuff
		case 27:
			// Do stuff
		case 28:
			// Do stuff
		case 29:
			// Do stuff
		case 30:
			// Do stuff
		case 31:
			// Do stuff
		case 32:
			// Do stuff
		case 33:
			// Do stuff
		case 34:
			// Do stuff
		case 35:
			// Do stuff
		case 36:
			// Do stuff
		case 37:
			// Do stuff
		case 38:
			// Do stuff
		case 39:
			// Do stuff
		case 40:
			// Do stuff
		case 41:
			// Do stuff
		case 42:
			// Do stuff
		case 43:
			// Do stuff
		case 44:
			// Do stuff
		case 45:
			// Do stuff
		case 46:
			// Do stuff
		case 47:
			// Do stuff
		case 48:
			// Do stuff
		case 49:
			// Do stuff
		}
	}
}

// TODO: Refactor to cover for arbitrary number of extensions (currently only covers
//       two as that's explicitly in the spec)
func getTargetReportDescriptor(cp goasterix.Compound) TargetReportDescriptor {
	trd := new(TargetReportDescriptor)

	tmp := cp.Primary[0]

	switch tmp & 0xE0 >> 5 {
	case 0:
		trd.ATP = "24-Bit ICAO address"
	case 1:
		trd.ATP = "Duplicate Address"
	case 2:
		trd.ATP = "Surface vehicle address"
	case 3:
		trd.ATP = "Anonymous address"
	default: // 4-7
		trd.ATP = "Reserved for future use"
	}

	switch tmp & 0x18 >> 3 {
	case 0:
		trd.ARC = "25ft"
	case 1:
		trd.ARC = "100ft"
	case 2:
		trd.ARC = "Unknown"
	case 3:
		trd.ARC = "Invalid"
	}

	if tmp&0x4 == 0 {
		trd.RC = "Default"
	} else {
		trd.RC = "Range Check passed, CPR Validation pending"
	}

	if tmp&0x2 == 0 {
		trd.RAB = "Report from target transponder"
	} else {
		trd.RAB = "Report from field monitor (fixed transponder)"
	}

	if tmp&0x1 != 0 {
		fx1 := new(FirstExtension)

		fstItem := 0
		fstByte := 0
		tmp = cp.Secondary[fstItem].Payload()[fstByte] //?

		if tmp&0x80 == 0 {
			fx1.DCR = "No differential correction"
		} else {
			fx1.DCR = "Differential correction"
		}

		if tmp&0x40 == 0 {
			fx1.GBS = "Not set"
		} else {
			fx1.GBS = "Set"
		}

		if tmp&0x20 == 0 {
			fx1.SIM = "Actual"
		} else {
			fx1.SIM = "Simulated"
		}

		if tmp&0x10 == 0 {
			fx1.TST = "Default"
		} else {
			fx1.TST = "Test target"
		}

		if tmp&0x8 == 0 {
			fx1.SAA = "Capable"
		} else {
			fx1.SAA = "Not capable"
		}

		switch tmp & 0x6 >> 1 {
		case 0:
			fx1.CL = "Report valid"
		case 1:
			fx1.CL = "Report suspect"
		case 2:
			fx1.CL = "No info"
		case 3:
			fx1.CL = "Reserved for future use"
		}

		if tmp&0x1 != 0 {
			fx2 := new(SecondExtension)

			sndItem := 0
			tmp = cp.Secondary[sndItem].Payload()[fstByte] //?

			if tmp&0x40 == 0 {
				fx2.LLC = "default"
			} else {
				fx2.LLC = "Target is suspect"
			}

			if tmp&0x20 == 0 {
				fx2.IPC = "default"
			} else {
				fx2.IPC = " Independent Position Check failed "
			}

			if tmp&0x10 == 0 {
				fx2.NOGO = "Not set"
			} else {
				fx2.NOGO = "Set"
			}

			if tmp&0x8 == 0 {
				fx2.CPR = "CPR validation correct"
			} else {
				fx2.CPR = "CPR vallidation failed"
			}

			if tmp&0x4 == 0 {
				fx2.LDPJ = "Not detected"
			} else {
				fx2.LDPJ = "Detected"
			}

			if tmp&0x2 == 0 {
				fx2.RCF = "Default"
			} else {
				fx2.RCF = "Range check failed"
			}

			// TODO: Investigate and implement sequential field extensions
			fx2.ThirdExtension = tmp & 0x1
		}

		trd.FX = fx1

	}

	return *trd
}

func wgs84Coordinates(data []byte) WGS84Coordinates {
	var pos WGS84Coordinates

	if len(data) == 6 {
		tmpLatitude := uint32(data[0])<<(2*BYTESIZE) + uint32(data[1])<<BYTESIZE + uint32(data[2])
		pos.Latitude = float32(goasterix.TwoComplement32(24, tmpLatitude)) * 0.00002145767

		tmpLongitude := uint32(data[3])<<(2*BYTESIZE) + uint32(data[4])<<BYTESIZE + uint32(data[5])
		pos.Longitude = float32(goasterix.TwoComplement32(32, tmpLongitude)) * 0.00002145767
	} else { // high precision data
		tmpLatitude := uint32(data[0])<<23 + uint32(data[1])<<15 + uint32(data[2])<<7 + uint32(data[3])
		pos.Latitude = float32(goasterix.TwoComplement32(32, tmpLatitude)) * 0.00000016764

		tmpLongitude := uint32(data[4])<<23 + uint32(data[5])<<15 + uint32(data[6])<<7 + uint32(data[7])
		pos.Longitude = float32(goasterix.TwoComplement32(32, tmpLongitude)) * 0.00000016764
	}

	return pos
}

func getAirSpeed(data [2]byte) AirSpeed {
	var speed AirSpeed

	tmp := data[0]
	speedValue := float64(uint32(data[0]&0x7F)<<8 + uint32(data[1]&0xFF))
	if tmp&0x80 == 0 {
		speed.IM = "IAS"
		speed.AirSpeed = speedValue * math.Pow(2, -14)

	} else {
		speed.IM = "Mach"
		speed.AirSpeed = speedValue * 0.001
	}

	return speed
}

func getTrueAirSpeed(data [2]byte) TrueAirSpeed {
	return TrueAirSpeed{
		RE:    int(data[0] & 0x80),
		Speed: int16(uint32(data[0]&0x7F)<<BYTESIZE + uint32(data[1]&0xFF)),
	}
}

// TODO: Double-check this...
func getGeometricHeight(data [2]byte) GeometricHeight {
	tmpHeight := goasterix.TwoComplement16(16, uint16(data[0])<<BYTESIZE+uint16(data[1]))
	greaterThan := false
	int16Max := int16(32767)
	if tmpHeight == int16Max {
		greaterThan = true
	}
	return GeometricHeight{
		Height:      float64(tmpHeight) * 6.25,
		GreaterThan: greaterThan,
	}
}
