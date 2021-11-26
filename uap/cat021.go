package uap

// Cat021v10 User Application Profile
// version 2.5
/*
var Cat021v10 = StandardUAP{
	Name:     "cat021_2.5",
	Category: 21,
	Version:  2.5,
	Items: []DataField{
		{
			FRN:      1,
			DataItem: "I021/010",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:         2,
			DataItem:    "I021/040",
			Description: "Target Report Descriptor",
			Type: TypeField{
				NameType:      Extended,
				PrimarySize:   1,
				SecondarySize: 1,
			},
		},
		{
			FRN:         3,
			DataItem:    "I021/161",
			Description: "Track Number",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:         4,
			DataItem:    "I021/015",
			Description: "Service Identification",
			Type: TypeField{
				NameType: Fixed,
				Size:     1,
			},
		},
		{
			FRN:         5,
			DataItem:    "I021/071",
			Description: "Time of Applicability for Position",
			Type: TypeField{
				NameType: Fixed,
				Size:     3,
			},
		},
		{
			FRN:         6,
			DataItem:    "I021/130",
			Description: "Position in WGS-84 co-ordinates",
			Type: TypeField{
				NameType: Fixed,
				Size:     6,
			},
		},
		{
			FRN:         7,
			DataItem:    "I021/131",
			Description: "Position in WGS-84 co-ordinates, high res",
			Type: TypeField{
				NameType: Fixed,
				Size:     8,
			},
		},
		{
			FRN:         8,
			DataItem:    "I021/072",
			Description: "Time of Applicability for Velocity",
			Type: TypeField{
				NameType: Fixed,
				Size:     3,
			},
		},
		{
			FRN:         9,
			DataItem:    "I021/150",
			Description: "Air Speed",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:         10,
			DataItem:    "I021/151",
			Description: "True Air Speed",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:         11,
			DataItem:    "I021/080",
			Description: "Target Address",
			Type: TypeField{
				NameType: Fixed,
				Size:     3,
			},
		},
		{
			FRN:         12,
			DataItem:    "I021/073",
			Description: "Time of Message Reception of Position",
			Type: TypeField{
				NameType: Fixed,
				Size:     3,
			},
		},
		{
			FRN:         13,
			DataItem:    "I021/074",
			Description: "Time of Message Reception of Position-High Precision",
			Type: TypeField{
				NameType: Fixed,
				Size:     4,
			},
		},
		{
			FRN:         14,
			DataItem:    "I021/075",
			Description: "Time of Message Reception of Velocity",
			Type: TypeField{
				NameType: Fixed,
				Size:     3,
			},
		},
		{
			FRN:         15,
			DataItem:    "I021/076",
			Description: "Time of Message Reception of Velocity-High Precision",
			Type: TypeField{
				NameType: Fixed,
				Size:     4,
			},
		},
		{
			FRN:         16,
			DataItem:    "I021/140",
			Description: "Geometric Height",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:         17,
			DataItem:    "I021/090",
			Description: "Quality Indicators",
			Type: TypeField{
				NameType:      Extended,
				PrimarySize:   1,
				SecondarySize: 1,
			},
		},
		{
			FRN:         18,
			DataItem:    "I021/210",
			Description: "MOPS Version",
			Type: TypeField{
				NameType: Fixed,
				Size:     1,
			},
		},
		{
			FRN:         19,
			DataItem:    "I021/070",
			Description: "Mode 3/A Code",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:         20,
			DataItem:    "I021/230",
			Description: "Roll Angle",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:         21,
			DataItem:    "I021/145",
			Description: "Flight Level",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:         22,
			DataItem:    "I021/152",
			Description: "Magnetic Heading",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:         23,
			DataItem:    "I021/200",
			Description: "Target Status",
			Type: TypeField{
				NameType: Fixed,
				Size:     1,
			},
		},
		{
			FRN:         24,
			DataItem:    "I021/155",
			Description: "Barometric Vertical Rate",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:         25,
			DataItem:    "I021/157",
			Description: "Geometric Vertical Rate",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:         26,
			DataItem:    "I021/160",
			Description: "Airborne Ground Vector",
			Type: TypeField{
				NameType: Fixed,
				Size:     4,
			},
		},
		{
			FRN:         27,
			DataItem:    "I021/165",
			Description: "Track Angle Rate",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:         28,
			DataItem:    "I021/177",
			Description: "Time of Report Transmission",
			Type: TypeField{
				NameType: Fixed,
				Size:     3,
			},
		},
		{
			FRN:         29,
			DataItem:    "I021/170",
			Description: "Target Identification",
			Type: TypeField{
				NameType: Fixed,
				Size:     6,
			},
		},
		{
			FRN:         30,
			DataItem:    "I021/020",
			Description: "Emitter Category",
			Type: TypeField{
				NameType: Fixed,
				Size:     1,
			},
		},
		{
			FRN:         31,
			DataItem:    "I021/220",
			Description: "Met Information",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {NameType: Fixed, Size: 2},
						7: {NameType: Fixed, Size: 2},
						6: {NameType: Fixed, Size: 2},
						5: {NameType: Fixed, Size: 1},
						4: {NameType: Spare},
						3: {NameType: Spare},
						2: {NameType: Spare},
					},
				},
			},
		},
		{
			FRN:         32,
			DataItem:    "I021/146",
			Description: "Selected Altitude",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:         33,
			DataItem:    "I021/148",
			Description: "Final State Selected Altitude",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:         34,
			DataItem:    "I021/110",
			Description: "Trajectory Intent",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {NameType: Fixed, Size: 1},
						7: {NameType: Repetitive, Size: 15},
						6: {NameType: Spare},
						5: {NameType: Spare},
						4: {NameType: Spare},
						3: {NameType: Spare},
						2: {NameType: Spare},
					},
				},
			},
		},
		{
			FRN:         35,
			DataItem:    "I021/016",
			Description: "Service Management",
			Type: TypeField{
				NameType: Fixed,
				Size:     1,
			},
		},
		{
			FRN:         36,
			DataItem:    "I021/008",
			Description: "Aircraft Operational Status",
			Type: TypeField{
				NameType: Fixed,
				Size:     1,
			},
		},
		{
			FRN:         37,
			DataItem:    "I021/271",
			Description: "Surface Capabilities and Characteristics",
			Type: TypeField{
				NameType:      Extended,
				PrimarySize:   1,
				SecondarySize: 1,
			},
		},
		{
			FRN:         38,
			DataItem:    "I021/132",
			Description: "Message Amplitude",
			Type: TypeField{
				NameType: Fixed,
				Size:     1,
			},
		},
		{
			FRN:         38,
			DataItem:    "I021/250",
			Description: "Mode S MB Data",
			Type: TypeField{
				NameType: Repetitive,
				Size:     8,
			},
		},
		{
			FRN:         40,
			DataItem:    "I021/260",
			Description: "ACAS Resolution Advisory Report",
			Type: TypeField{
				NameType: Fixed,
				Size:     7,
			},
		},
		{
			FRN:         41,
			DataItem:    "I021/400",
			Description: "Receiver ID",
			Type: TypeField{
				NameType: Fixed,
				Size:     1,
			},
		},
		{
			FRN:         42,
			DataItem:    "I021/295",
			Description: "Data Ages",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {NameType: Fixed, Size: 1},
						7: {NameType: Fixed, Size: 1},
						6: {NameType: Fixed, Size: 1},
						5: {NameType: Fixed, Size: 1},
						4: {NameType: Fixed, Size: 1},
						3: {NameType: Fixed, Size: 1},
						2: {NameType: Fixed, Size: 1},
					},
				},
			},
		},
		{
			FRN: 43, DataItem: "NA", Type: TypeField{NameType: Spare},
		},
		{
			FRN: 44, DataItem: "NA", Type: TypeField{NameType: Spare},
		},
		{
			FRN: 45, DataItem: "NA", Type: TypeField{NameType: Spare},
		},
		{
			FRN: 46, DataItem: "NA", Type: TypeField{NameType: Spare},
		},
		{
			FRN: 47, DataItem: "NA", Type: TypeField{NameType: Spare},
		},
		{
			FRN:         48,
			DataItem:    "RE-Data Item",
			Description: "Reserved Expansion Field",
			Type: TypeField{
				NameType: RE,
			},
		},
		{
			FRN:         49,
			DataItem:    "SP-Data Item",
			Description: "Special Purpose Field",
			Type: TypeField{
				NameType: SP,
			},
		},
	},
}
*/
