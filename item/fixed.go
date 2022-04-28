package item

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

func newFixed(field DataItem) DataItem {
	f := &Fixed{}
	f.Base.NewBase(field)
	f.Size = field.GetSize().ForFixed
	f.SubItems = field.GetSubItem()
	return f
}

func (f Fixed) GetSize() SizeField {
	s := SizeField{}
	s.ForFixed = f.Size
	return s
}

func (f Fixed) GetSubItem() []SubItem {
	return f.SubItems
}
func (f Fixed) GetCompound() []DataItem {
	return nil // not used, it's for implement DataItem interface
}

// Reader extracts a number(nb) of bytes(size) and returns a slice of bytes(data of item).
func (f *Fixed) Reader(rb *bytes.Reader) error {
	var err error
	tmp := make([]byte, f.Size)
	err = binary.Read(rb, binary.BigEndian, &tmp)
	if err != nil {
		return err
	}
	// check if they are defined
	if f.SubItems != nil {
		for _, subItem := range f.SubItems {
			err = subItem.Reader(tmp)
		}
	} else {
		f.Data = tmp
	}

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
	buf.WriteString(f.Base.DataItemName)
	buf.WriteByte(':')
	if f.SubItems != nil {
		for _, subItem := range f.SubItems {
			buf.WriteByte('[')
			buf.WriteString(subItem.String())
			buf.WriteByte(']')
		}
	} else {
		buf.WriteByte('[')
		buf.WriteString(hex.EncodeToString(f.Data))
		buf.WriteByte(']')
	}
	return buf.String()
}
