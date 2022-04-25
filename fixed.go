package goasterix

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

// Fixed length Data Fields shall comprise a fixed number of octets.
type Fixed struct {
	Base
	Data     []byte
	Size     uint8
	SubItems []SubItem
}

//func newFixed(field _uap.IDataField) Item {
func newFixed(field Item) Item {
	f := &Fixed{}
	f.Base.NewBase(field)
	f.Size = field.GetSize().ForFixed
	return f
}
func (f Fixed) GetSize() SizeField {
	s := SizeField{}
	s.ForFixed = f.Size
	return s
}
func (f Fixed) GetCompound() []Item {
	return nil // not used, it's for implement Item interface
}

// Reader extracts a number(nb) of bytes(size) and returns a slice of bytes(data of item).
func (f *Fixed) Reader(rb *bytes.Reader) error {
	var err error
	f.Data = make([]byte, f.Size)
	err = binary.Read(rb, binary.BigEndian, &f.Data)
	if err != nil {
		f.Data = nil
		return err
	}
	//tmp := f.Data
	//for _, subItem := range f.SubItems {
	//	//tmp := make([]byte, subItem.SizeBit)
	//	//subItem.Pos.From
	//}

	return err
}

// Payload returns this dataField as bytes.
func (f Fixed) Payload() []byte {
	var p []byte
	p = append(p, f.Data...)
	return p
}

// String implements fmt.Stringer in hexadecimal
func (f Fixed) String() string {
	var buf bytes.Buffer
	buf.Reset()
	buf.WriteString(f.Base.DataItem)
	buf.WriteByte(':')
	buf.WriteString(hex.EncodeToString(f.Data))
	return buf.String()
}