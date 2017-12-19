package dns

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *A) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "A"
	o = append(o, 0xa1, 0x41)
	o = msgp.AppendBytes(o, ipToByte(z.A))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *A) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "A":
			{
				var zb0003 []byte
				zb0003, bts, err = msgp.ReadBytesBytes(bts, ipToByte(z.A))
				if err != nil {
					return
				}
				z.A = byteToIP(zb0003)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *A) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 2 + msgp.BytesPrefixSize + len(ipToByte(z.A))
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *AAAA) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "AAAA"
	o = append(o, 0xa4, 0x41, 0x41, 0x41, 0x41)
	o = msgp.AppendBytes(o, ipToByte(z.AAAA))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AAAA) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "AAAA":
			{
				var zb0003 []byte
				zb0003, bts, err = msgp.ReadBytesBytes(bts, ipToByte(z.AAAA))
				if err != nil {
					return
				}
				z.AAAA = byteToIP(zb0003)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AAAA) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 5 + msgp.BytesPrefixSize + len(ipToByte(z.AAAA))
	return
}

// MarshalMsg implements msgp.Marshaler
func (z AFSDB) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Hdr"
	o = append(o, 0x83, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Subtype"
	o = append(o, 0xa7, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendUint16(o, z.Subtype)
	// string "Hostname"
	o = append(o, 0xa8, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Hostname)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AFSDB) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Subtype":
			z.Subtype, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Hostname":
			z.Hostname, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z AFSDB) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 8 + msgp.Uint16Size + 9 + msgp.StringPrefixSize + len(z.Hostname)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ANY) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Hdr"
	o = append(o, 0x81, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ANY) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ANY) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr))
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *AVC) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Txt"
	o = append(o, 0xa3, 0x54, 0x78, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Txt)))
	for za0001 := range z.Txt {
		o = msgp.AppendString(o, z.Txt[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AVC) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Txt":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Txt) >= int(zb0003) {
				z.Txt = (z.Txt)[:zb0003]
			} else {
				z.Txt = make([]string, zb0003)
			}
			for za0001 := range z.Txt {
				z.Txt[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AVC) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 4 + msgp.ArrayHeaderSize
	for za0001 := range z.Txt {
		s += msgp.StringPrefixSize + len(z.Txt[za0001])
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CAA) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Hdr"
	o = append(o, 0x84, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Flag"
	o = append(o, 0xa4, 0x46, 0x6c, 0x61, 0x67)
	o = msgp.AppendUint8(o, z.Flag)
	// string "Tag"
	o = append(o, 0xa3, 0x54, 0x61, 0x67)
	o = msgp.AppendString(o, z.Tag)
	// string "Value"
	o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendString(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CAA) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Flag":
			z.Flag, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Tag":
			z.Tag, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Value":
			z.Value, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CAA) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 5 + msgp.Uint8Size + 4 + msgp.StringPrefixSize + len(z.Tag) + 6 + msgp.StringPrefixSize + len(z.Value)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CDNSKEY) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "DNSKEY"
	o = append(o, 0x81, 0xa6, 0x44, 0x4e, 0x53, 0x4b, 0x45, 0x59)
	o, err = z.DNSKEY.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CDNSKEY) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "DNSKEY":
			bts, err = z.DNSKEY.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CDNSKEY) Msgsize() (s int) {
	s = 1 + 7 + z.DNSKEY.Msgsize()
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CDS) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "DS"
	o = append(o, 0x81, 0xa2, 0x44, 0x53)
	o, err = z.DS.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CDS) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "DS":
			bts, err = z.DS.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CDS) Msgsize() (s int) {
	s = 1 + 3 + z.DS.Msgsize()
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CERT) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Hdr"
	o = append(o, 0x85, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendUint16(o, z.Type)
	// string "KeyTag"
	o = append(o, 0xa6, 0x4b, 0x65, 0x79, 0x54, 0x61, 0x67)
	o = msgp.AppendUint16(o, z.KeyTag)
	// string "Algorithm"
	o = append(o, 0xa9, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d)
	o = msgp.AppendUint8(o, z.Algorithm)
	// string "Certificate"
	o = append(o, 0xab, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65)
	o = msgp.AppendString(o, z.Certificate)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CERT) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Type":
			z.Type, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "KeyTag":
			z.KeyTag, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Algorithm":
			z.Algorithm, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Certificate":
			z.Certificate, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CERT) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 5 + msgp.Uint16Size + 7 + msgp.Uint16Size + 10 + msgp.Uint8Size + 12 + msgp.StringPrefixSize + len(z.Certificate)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z CNAME) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Target"
	o = append(o, 0xa6, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74)
	o = msgp.AppendString(o, z.Target)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CNAME) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Target":
			z.Target, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z CNAME) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 7 + msgp.StringPrefixSize + len(z.Target)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CSYNC) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Hdr"
	o = append(o, 0x84, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Serial"
	o = append(o, 0xa6, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c)
	o = msgp.AppendUint32(o, z.Serial)
	// string "Flags"
	o = append(o, 0xa5, 0x46, 0x6c, 0x61, 0x67, 0x73)
	o = msgp.AppendUint16(o, z.Flags)
	// string "TypeBitMap"
	o = append(o, 0xaa, 0x54, 0x79, 0x70, 0x65, 0x42, 0x69, 0x74, 0x4d, 0x61, 0x70)
	o = msgp.AppendArrayHeader(o, uint32(len(z.TypeBitMap)))
	for za0001 := range z.TypeBitMap {
		o = msgp.AppendUint16(o, z.TypeBitMap[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CSYNC) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Serial":
			z.Serial, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "Flags":
			z.Flags, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "TypeBitMap":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.TypeBitMap) >= int(zb0003) {
				z.TypeBitMap = (z.TypeBitMap)[:zb0003]
			} else {
				z.TypeBitMap = make([]uint16, zb0003)
			}
			for za0001 := range z.TypeBitMap {
				z.TypeBitMap[za0001], bts, err = msgp.ReadUint16Bytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CSYNC) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 7 + msgp.Uint32Size + 6 + msgp.Uint16Size + 11 + msgp.ArrayHeaderSize + (len(z.TypeBitMap) * (msgp.Uint16Size))
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Class) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint16(o, uint16(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Class) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint16
		zb0001, bts, err = msgp.ReadUint16Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Class(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Class) Msgsize() (s int) {
	s = msgp.Uint16Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z DHCID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Digest"
	o = append(o, 0xa6, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74)
	o = msgp.AppendString(o, z.Digest)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DHCID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Digest":
			z.Digest, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z DHCID) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 7 + msgp.StringPrefixSize + len(z.Digest)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DLV) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "DS"
	o = append(o, 0x81, 0xa2, 0x44, 0x53)
	o, err = z.DS.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DLV) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "DS":
			bts, err = z.DS.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *DLV) Msgsize() (s int) {
	s = 1 + 3 + z.DS.Msgsize()
	return
}

// MarshalMsg implements msgp.Marshaler
func (z DNAME) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Target"
	o = append(o, 0xa6, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74)
	o = msgp.AppendString(o, z.Target)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DNAME) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Target":
			z.Target, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z DNAME) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 7 + msgp.StringPrefixSize + len(z.Target)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DNSKEY) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Hdr"
	o = append(o, 0x85, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Flags"
	o = append(o, 0xa5, 0x46, 0x6c, 0x61, 0x67, 0x73)
	o = msgp.AppendUint16(o, z.Flags)
	// string "Protocol"
	o = append(o, 0xa8, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c)
	o = msgp.AppendUint8(o, z.Protocol)
	// string "Algorithm"
	o = append(o, 0xa9, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d)
	o = msgp.AppendUint8(o, z.Algorithm)
	// string "PublicKey"
	o = append(o, 0xa9, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.PublicKey)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DNSKEY) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Flags":
			z.Flags, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Protocol":
			z.Protocol, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Algorithm":
			z.Algorithm, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "PublicKey":
			z.PublicKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *DNSKEY) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 6 + msgp.Uint16Size + 9 + msgp.Uint8Size + 10 + msgp.Uint8Size + 10 + msgp.StringPrefixSize + len(z.PublicKey)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DS) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Hdr"
	o = append(o, 0x85, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "KeyTag"
	o = append(o, 0xa6, 0x4b, 0x65, 0x79, 0x54, 0x61, 0x67)
	o = msgp.AppendUint16(o, z.KeyTag)
	// string "Algorithm"
	o = append(o, 0xa9, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d)
	o = msgp.AppendUint8(o, z.Algorithm)
	// string "DigestType"
	o = append(o, 0xaa, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendUint8(o, z.DigestType)
	// string "Digest"
	o = append(o, 0xa6, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74)
	o = msgp.AppendString(o, z.Digest)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DS) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "KeyTag":
			z.KeyTag, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Algorithm":
			z.Algorithm, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "DigestType":
			z.DigestType, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Digest":
			z.Digest, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *DS) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 7 + msgp.Uint16Size + 10 + msgp.Uint8Size + 11 + msgp.Uint8Size + 7 + msgp.StringPrefixSize + len(z.Digest)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z EID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Endpoint"
	o = append(o, 0xa8, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Endpoint)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *EID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Endpoint":
			z.Endpoint, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z EID) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 9 + msgp.StringPrefixSize + len(z.Endpoint)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z EUI48) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Address"
	o = append(o, 0xa7, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
	o = msgp.AppendUint64(o, z.Address)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *EUI48) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Address":
			z.Address, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z EUI48) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 8 + msgp.Uint64Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z EUI64) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Address"
	o = append(o, 0xa7, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
	o = msgp.AppendUint64(o, z.Address)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *EUI64) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Address":
			z.Address, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z EUI64) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 8 + msgp.Uint64Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z GID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Gid"
	o = append(o, 0xa3, 0x47, 0x69, 0x64)
	o = msgp.AppendUint32(o, z.Gid)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Gid":
			z.Gid, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z GID) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 4 + msgp.Uint32Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GPOS) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Hdr"
	o = append(o, 0x84, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Longitude"
	o = append(o, 0xa9, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65)
	o = msgp.AppendString(o, z.Longitude)
	// string "Latitude"
	o = append(o, 0xa8, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65)
	o = msgp.AppendString(o, z.Latitude)
	// string "Altitude"
	o = append(o, 0xa8, 0x41, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65)
	o = msgp.AppendString(o, z.Altitude)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GPOS) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Longitude":
			z.Longitude, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Latitude":
			z.Latitude, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Altitude":
			z.Altitude, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *GPOS) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 10 + msgp.StringPrefixSize + len(z.Longitude) + 9 + msgp.StringPrefixSize + len(z.Latitude) + 9 + msgp.StringPrefixSize + len(z.Altitude)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z HINFO) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Hdr"
	o = append(o, 0x83, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Cpu"
	o = append(o, 0xa3, 0x43, 0x70, 0x75)
	o = msgp.AppendString(o, z.Cpu)
	// string "Os"
	o = append(o, 0xa2, 0x4f, 0x73)
	o = msgp.AppendString(o, z.Os)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *HINFO) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Cpu":
			z.Cpu, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Os":
			z.Os, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z HINFO) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 4 + msgp.StringPrefixSize + len(z.Cpu) + 3 + msgp.StringPrefixSize + len(z.Os)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *HIP) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "Hdr"
	o = append(o, 0x87, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "HitLength"
	o = append(o, 0xa9, 0x48, 0x69, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68)
	o = msgp.AppendUint8(o, z.HitLength)
	// string "PublicKeyAlgorithm"
	o = append(o, 0xb2, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d)
	o = msgp.AppendUint8(o, z.PublicKeyAlgorithm)
	// string "PublicKeyLength"
	o = append(o, 0xaf, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68)
	o = msgp.AppendUint16(o, z.PublicKeyLength)
	// string "Hit"
	o = append(o, 0xa3, 0x48, 0x69, 0x74)
	o = msgp.AppendString(o, z.Hit)
	// string "PublicKey"
	o = append(o, 0xa9, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.PublicKey)
	// string "RendezvousServers"
	o = append(o, 0xb1, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.RendezvousServers)))
	for za0001 := range z.RendezvousServers {
		o = msgp.AppendString(o, z.RendezvousServers[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *HIP) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "HitLength":
			z.HitLength, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "PublicKeyAlgorithm":
			z.PublicKeyAlgorithm, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "PublicKeyLength":
			z.PublicKeyLength, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Hit":
			z.Hit, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "PublicKey":
			z.PublicKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RendezvousServers":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.RendezvousServers) >= int(zb0003) {
				z.RendezvousServers = (z.RendezvousServers)[:zb0003]
			} else {
				z.RendezvousServers = make([]string, zb0003)
			}
			for za0001 := range z.RendezvousServers {
				z.RendezvousServers[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *HIP) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 10 + msgp.Uint8Size + 19 + msgp.Uint8Size + 16 + msgp.Uint16Size + 4 + msgp.StringPrefixSize + len(z.Hit) + 10 + msgp.StringPrefixSize + len(z.PublicKey) + 18 + msgp.ArrayHeaderSize
	for za0001 := range z.RendezvousServers {
		s += msgp.StringPrefixSize + len(z.RendezvousServers[za0001])
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Header) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "Id"
	o = append(o, 0x86, 0xa2, 0x49, 0x64)
	o = msgp.AppendUint16(o, z.Id)
	// string "Bits"
	o = append(o, 0xa4, 0x42, 0x69, 0x74, 0x73)
	o = msgp.AppendUint16(o, z.Bits)
	// string "Qdcount"
	o = append(o, 0xa7, 0x51, 0x64, 0x63, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendUint16(o, z.Qdcount)
	// string "Ancount"
	o = append(o, 0xa7, 0x41, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendUint16(o, z.Ancount)
	// string "Nscount"
	o = append(o, 0xa7, 0x4e, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendUint16(o, z.Nscount)
	// string "Arcount"
	o = append(o, 0xa7, 0x41, 0x72, 0x63, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendUint16(o, z.Arcount)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Header) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Bits":
			z.Bits, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Qdcount":
			z.Qdcount, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Ancount":
			z.Ancount, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Nscount":
			z.Nscount, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Arcount":
			z.Arcount, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Header) Msgsize() (s int) {
	s = 1 + 3 + msgp.Uint16Size + 5 + msgp.Uint16Size + 8 + msgp.Uint16Size + 8 + msgp.Uint16Size + 8 + msgp.Uint16Size + 8 + msgp.Uint16Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *KEY) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "DNSKEY"
	o = append(o, 0x81, 0xa6, 0x44, 0x4e, 0x53, 0x4b, 0x45, 0x59)
	o, err = z.DNSKEY.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *KEY) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "DNSKEY":
			bts, err = z.DNSKEY.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *KEY) Msgsize() (s int) {
	s = 1 + 7 + z.DNSKEY.Msgsize()
	return
}

// MarshalMsg implements msgp.Marshaler
func (z KX) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Hdr"
	o = append(o, 0x83, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Preference"
	o = append(o, 0xaa, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65)
	o = msgp.AppendUint16(o, z.Preference)
	// string "Exchanger"
	o = append(o, 0xa9, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x72)
	o = msgp.AppendString(o, z.Exchanger)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *KX) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Preference":
			z.Preference, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Exchanger":
			z.Exchanger, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z KX) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 11 + msgp.Uint16Size + 10 + msgp.StringPrefixSize + len(z.Exchanger)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *L32) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Hdr"
	o = append(o, 0x83, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Preference"
	o = append(o, 0xaa, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65)
	o = msgp.AppendUint16(o, z.Preference)
	// string "Locator32"
	o = append(o, 0xa9, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x33, 0x32)
	o = msgp.AppendBytes(o, ipToByte(z.Locator32))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *L32) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Preference":
			z.Preference, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Locator32":
			{
				var zb0003 []byte
				zb0003, bts, err = msgp.ReadBytesBytes(bts, ipToByte(z.Locator32))
				if err != nil {
					return
				}
				z.Locator32 = byteToIP(zb0003)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *L32) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 11 + msgp.Uint16Size + 10 + msgp.BytesPrefixSize + len(ipToByte(z.Locator32))
	return
}

// MarshalMsg implements msgp.Marshaler
func (z L64) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Hdr"
	o = append(o, 0x83, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Preference"
	o = append(o, 0xaa, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65)
	o = msgp.AppendUint16(o, z.Preference)
	// string "Locator64"
	o = append(o, 0xa9, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x36, 0x34)
	o = msgp.AppendUint64(o, z.Locator64)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *L64) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Preference":
			z.Preference, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Locator64":
			z.Locator64, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z L64) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 11 + msgp.Uint16Size + 10 + msgp.Uint64Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *LOC) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 8
	// string "Hdr"
	o = append(o, 0x88, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Version"
	o = append(o, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendUint8(o, z.Version)
	// string "Size"
	o = append(o, 0xa4, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendUint8(o, z.Size)
	// string "HorizPre"
	o = append(o, 0xa8, 0x48, 0x6f, 0x72, 0x69, 0x7a, 0x50, 0x72, 0x65)
	o = msgp.AppendUint8(o, z.HorizPre)
	// string "VertPre"
	o = append(o, 0xa7, 0x56, 0x65, 0x72, 0x74, 0x50, 0x72, 0x65)
	o = msgp.AppendUint8(o, z.VertPre)
	// string "Latitude"
	o = append(o, 0xa8, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65)
	o = msgp.AppendUint32(o, z.Latitude)
	// string "Longitude"
	o = append(o, 0xa9, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65)
	o = msgp.AppendUint32(o, z.Longitude)
	// string "Altitude"
	o = append(o, 0xa8, 0x41, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65)
	o = msgp.AppendUint32(o, z.Altitude)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LOC) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Version":
			z.Version, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Size":
			z.Size, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "HorizPre":
			z.HorizPre, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "VertPre":
			z.VertPre, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Latitude":
			z.Latitude, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "Longitude":
			z.Longitude, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "Altitude":
			z.Altitude, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *LOC) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 8 + msgp.Uint8Size + 5 + msgp.Uint8Size + 9 + msgp.Uint8Size + 8 + msgp.Uint8Size + 9 + msgp.Uint32Size + 10 + msgp.Uint32Size + 9 + msgp.Uint32Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z LP) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Hdr"
	o = append(o, 0x83, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Preference"
	o = append(o, 0xaa, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65)
	o = msgp.AppendUint16(o, z.Preference)
	// string "Fqdn"
	o = append(o, 0xa4, 0x46, 0x71, 0x64, 0x6e)
	o = msgp.AppendString(o, z.Fqdn)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LP) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Preference":
			z.Preference, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Fqdn":
			z.Fqdn, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z LP) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 11 + msgp.Uint16Size + 5 + msgp.StringPrefixSize + len(z.Fqdn)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MB) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Mb"
	o = append(o, 0xa2, 0x4d, 0x62)
	o = msgp.AppendString(o, z.Mb)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MB) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Mb":
			z.Mb, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MB) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 3 + msgp.StringPrefixSize + len(z.Mb)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MD) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Md"
	o = append(o, 0xa2, 0x4d, 0x64)
	o = msgp.AppendString(o, z.Md)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MD) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Md":
			z.Md, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MD) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 3 + msgp.StringPrefixSize + len(z.Md)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MF) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Mf"
	o = append(o, 0xa2, 0x4d, 0x66)
	o = msgp.AppendString(o, z.Mf)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MF) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Mf":
			z.Mf, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MF) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 3 + msgp.StringPrefixSize + len(z.Mf)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MG) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Mg"
	o = append(o, 0xa2, 0x4d, 0x67)
	o = msgp.AppendString(o, z.Mg)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MG) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Mg":
			z.Mg, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MG) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 3 + msgp.StringPrefixSize + len(z.Mg)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MINFO) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Hdr"
	o = append(o, 0x83, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Rmail"
	o = append(o, 0xa5, 0x52, 0x6d, 0x61, 0x69, 0x6c)
	o = msgp.AppendString(o, z.Rmail)
	// string "Email"
	o = append(o, 0xa5, 0x45, 0x6d, 0x61, 0x69, 0x6c)
	o = msgp.AppendString(o, z.Email)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MINFO) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Rmail":
			z.Rmail, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Email":
			z.Email, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MINFO) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 6 + msgp.StringPrefixSize + len(z.Rmail) + 6 + msgp.StringPrefixSize + len(z.Email)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MR) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Mr"
	o = append(o, 0xa2, 0x4d, 0x72)
	o = msgp.AppendString(o, z.Mr)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MR) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Mr":
			z.Mr, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MR) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 3 + msgp.StringPrefixSize + len(z.Mr)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MX) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Hdr"
	o = append(o, 0x83, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Preference"
	o = append(o, 0xaa, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65)
	o = msgp.AppendUint16(o, z.Preference)
	// string "Mx"
	o = append(o, 0xa2, 0x4d, 0x78)
	o = msgp.AppendString(o, z.Mx)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MX) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Preference":
			z.Preference, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Mx":
			z.Mx, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MX) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 11 + msgp.Uint16Size + 3 + msgp.StringPrefixSize + len(z.Mx)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NAPTR) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "Hdr"
	o = append(o, 0x87, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Order"
	o = append(o, 0xa5, 0x4f, 0x72, 0x64, 0x65, 0x72)
	o = msgp.AppendUint16(o, z.Order)
	// string "Preference"
	o = append(o, 0xaa, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65)
	o = msgp.AppendUint16(o, z.Preference)
	// string "Flags"
	o = append(o, 0xa5, 0x46, 0x6c, 0x61, 0x67, 0x73)
	o = msgp.AppendString(o, z.Flags)
	// string "Service"
	o = append(o, 0xa7, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65)
	o = msgp.AppendString(o, z.Service)
	// string "Regexp"
	o = append(o, 0xa6, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70)
	o = msgp.AppendString(o, z.Regexp)
	// string "Replacement"
	o = append(o, 0xab, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Replacement)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NAPTR) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Order":
			z.Order, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Preference":
			z.Preference, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Flags":
			z.Flags, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Service":
			z.Service, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Regexp":
			z.Regexp, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Replacement":
			z.Replacement, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NAPTR) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 6 + msgp.Uint16Size + 11 + msgp.Uint16Size + 6 + msgp.StringPrefixSize + len(z.Flags) + 8 + msgp.StringPrefixSize + len(z.Service) + 7 + msgp.StringPrefixSize + len(z.Regexp) + 12 + msgp.StringPrefixSize + len(z.Replacement)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z NID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Hdr"
	o = append(o, 0x83, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Preference"
	o = append(o, 0xaa, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65)
	o = msgp.AppendUint16(o, z.Preference)
	// string "NodeID"
	o = append(o, 0xa6, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44)
	o = msgp.AppendUint64(o, z.NodeID)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Preference":
			z.Preference, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "NodeID":
			z.NodeID, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z NID) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 11 + msgp.Uint16Size + 7 + msgp.Uint64Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z NIMLOC) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Locator"
	o = append(o, 0xa7, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72)
	o = msgp.AppendString(o, z.Locator)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NIMLOC) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Locator":
			z.Locator, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z NIMLOC) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 8 + msgp.StringPrefixSize + len(z.Locator)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NINFO) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "ZSData"
	o = append(o, 0xa6, 0x5a, 0x53, 0x44, 0x61, 0x74, 0x61)
	o = msgp.AppendArrayHeader(o, uint32(len(z.ZSData)))
	for za0001 := range z.ZSData {
		o = msgp.AppendString(o, z.ZSData[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NINFO) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "ZSData":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.ZSData) >= int(zb0003) {
				z.ZSData = (z.ZSData)[:zb0003]
			} else {
				z.ZSData = make([]string, zb0003)
			}
			for za0001 := range z.ZSData {
				z.ZSData[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NINFO) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 7 + msgp.ArrayHeaderSize
	for za0001 := range z.ZSData {
		s += msgp.StringPrefixSize + len(z.ZSData[za0001])
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z NS) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Ns"
	o = append(o, 0xa2, 0x4e, 0x73)
	o = msgp.AppendString(o, z.Ns)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NS) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Ns":
			z.Ns, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z NS) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 3 + msgp.StringPrefixSize + len(z.Ns)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z NSAPPTR) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Ptr"
	o = append(o, 0xa3, 0x50, 0x74, 0x72)
	o = msgp.AppendString(o, z.Ptr)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NSAPPTR) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Ptr":
			z.Ptr, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z NSAPPTR) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 4 + msgp.StringPrefixSize + len(z.Ptr)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NSEC) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Hdr"
	o = append(o, 0x83, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "NextDomain"
	o = append(o, 0xaa, 0x4e, 0x65, 0x78, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e)
	o = msgp.AppendString(o, z.NextDomain)
	// string "TypeBitMap"
	o = append(o, 0xaa, 0x54, 0x79, 0x70, 0x65, 0x42, 0x69, 0x74, 0x4d, 0x61, 0x70)
	o = msgp.AppendArrayHeader(o, uint32(len(z.TypeBitMap)))
	for za0001 := range z.TypeBitMap {
		o = msgp.AppendUint16(o, z.TypeBitMap[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NSEC) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "NextDomain":
			z.NextDomain, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "TypeBitMap":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.TypeBitMap) >= int(zb0003) {
				z.TypeBitMap = (z.TypeBitMap)[:zb0003]
			} else {
				z.TypeBitMap = make([]uint16, zb0003)
			}
			for za0001 := range z.TypeBitMap {
				z.TypeBitMap[za0001], bts, err = msgp.ReadUint16Bytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NSEC) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 11 + msgp.StringPrefixSize + len(z.NextDomain) + 11 + msgp.ArrayHeaderSize + (len(z.TypeBitMap) * (msgp.Uint16Size))
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NSEC3) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "Hdr"
	o = append(o, 0x89, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Hash"
	o = append(o, 0xa4, 0x48, 0x61, 0x73, 0x68)
	o = msgp.AppendUint8(o, z.Hash)
	// string "Flags"
	o = append(o, 0xa5, 0x46, 0x6c, 0x61, 0x67, 0x73)
	o = msgp.AppendUint8(o, z.Flags)
	// string "Iterations"
	o = append(o, 0xaa, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73)
	o = msgp.AppendUint16(o, z.Iterations)
	// string "SaltLength"
	o = append(o, 0xaa, 0x53, 0x61, 0x6c, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68)
	o = msgp.AppendUint8(o, z.SaltLength)
	// string "Salt"
	o = append(o, 0xa4, 0x53, 0x61, 0x6c, 0x74)
	o = msgp.AppendString(o, z.Salt)
	// string "HashLength"
	o = append(o, 0xaa, 0x48, 0x61, 0x73, 0x68, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68)
	o = msgp.AppendUint8(o, z.HashLength)
	// string "NextDomain"
	o = append(o, 0xaa, 0x4e, 0x65, 0x78, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e)
	o = msgp.AppendString(o, z.NextDomain)
	// string "TypeBitMap"
	o = append(o, 0xaa, 0x54, 0x79, 0x70, 0x65, 0x42, 0x69, 0x74, 0x4d, 0x61, 0x70)
	o = msgp.AppendArrayHeader(o, uint32(len(z.TypeBitMap)))
	for za0001 := range z.TypeBitMap {
		o = msgp.AppendUint16(o, z.TypeBitMap[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NSEC3) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Hash":
			z.Hash, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Flags":
			z.Flags, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Iterations":
			z.Iterations, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "SaltLength":
			z.SaltLength, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Salt":
			z.Salt, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "HashLength":
			z.HashLength, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "NextDomain":
			z.NextDomain, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "TypeBitMap":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.TypeBitMap) >= int(zb0003) {
				z.TypeBitMap = (z.TypeBitMap)[:zb0003]
			} else {
				z.TypeBitMap = make([]uint16, zb0003)
			}
			for za0001 := range z.TypeBitMap {
				z.TypeBitMap[za0001], bts, err = msgp.ReadUint16Bytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NSEC3) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 5 + msgp.Uint8Size + 6 + msgp.Uint8Size + 11 + msgp.Uint16Size + 11 + msgp.Uint8Size + 5 + msgp.StringPrefixSize + len(z.Salt) + 11 + msgp.Uint8Size + 11 + msgp.StringPrefixSize + len(z.NextDomain) + 11 + msgp.ArrayHeaderSize + (len(z.TypeBitMap) * (msgp.Uint16Size))
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NSEC3PARAM) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "Hdr"
	o = append(o, 0x86, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Hash"
	o = append(o, 0xa4, 0x48, 0x61, 0x73, 0x68)
	o = msgp.AppendUint8(o, z.Hash)
	// string "Flags"
	o = append(o, 0xa5, 0x46, 0x6c, 0x61, 0x67, 0x73)
	o = msgp.AppendUint8(o, z.Flags)
	// string "Iterations"
	o = append(o, 0xaa, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73)
	o = msgp.AppendUint16(o, z.Iterations)
	// string "SaltLength"
	o = append(o, 0xaa, 0x53, 0x61, 0x6c, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68)
	o = msgp.AppendUint8(o, z.SaltLength)
	// string "Salt"
	o = append(o, 0xa4, 0x53, 0x61, 0x6c, 0x74)
	o = msgp.AppendString(o, z.Salt)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NSEC3PARAM) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Hash":
			z.Hash, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Flags":
			z.Flags, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Iterations":
			z.Iterations, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "SaltLength":
			z.SaltLength, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Salt":
			z.Salt, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NSEC3PARAM) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 5 + msgp.Uint8Size + 6 + msgp.Uint8Size + 11 + msgp.Uint16Size + 11 + msgp.Uint8Size + 5 + msgp.StringPrefixSize + len(z.Salt)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Name) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Name) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			return
		}
		(*z) = Name(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Name) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}

// MarshalMsg implements msgp.Marshaler
func (z OPENPGPKEY) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "PublicKey"
	o = append(o, 0xa9, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.PublicKey)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OPENPGPKEY) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "PublicKey":
			z.PublicKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z OPENPGPKEY) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 10 + msgp.StringPrefixSize + len(z.PublicKey)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z PTR) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Ptr"
	o = append(o, 0xa3, 0x50, 0x74, 0x72)
	o = msgp.AppendString(o, z.Ptr)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PTR) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Ptr":
			z.Ptr, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z PTR) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 4 + msgp.StringPrefixSize + len(z.Ptr)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PX) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Hdr"
	o = append(o, 0x84, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Preference"
	o = append(o, 0xaa, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65)
	o = msgp.AppendUint16(o, z.Preference)
	// string "Map822"
	o = append(o, 0xa6, 0x4d, 0x61, 0x70, 0x38, 0x32, 0x32)
	o = msgp.AppendString(o, z.Map822)
	// string "Mapx400"
	o = append(o, 0xa7, 0x4d, 0x61, 0x70, 0x78, 0x34, 0x30, 0x30)
	o = msgp.AppendString(o, z.Mapx400)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PX) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Preference":
			z.Preference, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Map822":
			z.Map822, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Mapx400":
			z.Mapx400, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PX) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 11 + msgp.Uint16Size + 7 + msgp.StringPrefixSize + len(z.Map822) + 8 + msgp.StringPrefixSize + len(z.Mapx400)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Question) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Name"
	o = append(o, 0x83, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Qtype"
	o = append(o, 0xa5, 0x51, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendUint16(o, z.Qtype)
	// string "Qclass"
	o = append(o, 0xa6, 0x51, 0x63, 0x6c, 0x61, 0x73, 0x73)
	o = msgp.AppendUint16(o, z.Qclass)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Question) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Qtype":
			z.Qtype, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Qclass":
			z.Qclass, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Question) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 6 + msgp.Uint16Size + 7 + msgp.Uint16Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z RFC3597) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Rdata"
	o = append(o, 0xa5, 0x52, 0x64, 0x61, 0x74, 0x61)
	o = msgp.AppendString(o, z.Rdata)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RFC3597) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Rdata":
			z.Rdata, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z RFC3597) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 6 + msgp.StringPrefixSize + len(z.Rdata)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *RKEY) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Hdr"
	o = append(o, 0x85, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Flags"
	o = append(o, 0xa5, 0x46, 0x6c, 0x61, 0x67, 0x73)
	o = msgp.AppendUint16(o, z.Flags)
	// string "Protocol"
	o = append(o, 0xa8, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c)
	o = msgp.AppendUint8(o, z.Protocol)
	// string "Algorithm"
	o = append(o, 0xa9, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d)
	o = msgp.AppendUint8(o, z.Algorithm)
	// string "PublicKey"
	o = append(o, 0xa9, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.PublicKey)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RKEY) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Flags":
			z.Flags, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Protocol":
			z.Protocol, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Algorithm":
			z.Algorithm, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "PublicKey":
			z.PublicKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *RKEY) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 6 + msgp.Uint16Size + 9 + msgp.Uint8Size + 10 + msgp.Uint8Size + 10 + msgp.StringPrefixSize + len(z.PublicKey)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z RP) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Hdr"
	o = append(o, 0x83, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Mbox"
	o = append(o, 0xa4, 0x4d, 0x62, 0x6f, 0x78)
	o = msgp.AppendString(o, z.Mbox)
	// string "Txt"
	o = append(o, 0xa3, 0x54, 0x78, 0x74)
	o = msgp.AppendString(o, z.Txt)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RP) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Mbox":
			z.Mbox, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Txt":
			z.Txt, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z RP) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 5 + msgp.StringPrefixSize + len(z.Mbox) + 4 + msgp.StringPrefixSize + len(z.Txt)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *RRSIG) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "Hdr"
	o = append(o, 0x8a, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "TypeCovered"
	o = append(o, 0xab, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64)
	o = msgp.AppendUint16(o, z.TypeCovered)
	// string "Algorithm"
	o = append(o, 0xa9, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d)
	o = msgp.AppendUint8(o, z.Algorithm)
	// string "Labels"
	o = append(o, 0xa6, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73)
	o = msgp.AppendUint8(o, z.Labels)
	// string "OrigTtl"
	o = append(o, 0xa7, 0x4f, 0x72, 0x69, 0x67, 0x54, 0x74, 0x6c)
	o = msgp.AppendUint32(o, z.OrigTtl)
	// string "Expiration"
	o = append(o, 0xaa, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendUint32(o, z.Expiration)
	// string "Inception"
	o = append(o, 0xa9, 0x49, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendUint32(o, z.Inception)
	// string "KeyTag"
	o = append(o, 0xa6, 0x4b, 0x65, 0x79, 0x54, 0x61, 0x67)
	o = msgp.AppendUint16(o, z.KeyTag)
	// string "SignerName"
	o = append(o, 0xaa, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.SignerName)
	// string "Signature"
	o = append(o, 0xa9, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65)
	o = msgp.AppendString(o, z.Signature)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RRSIG) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "TypeCovered":
			z.TypeCovered, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Algorithm":
			z.Algorithm, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Labels":
			z.Labels, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "OrigTtl":
			z.OrigTtl, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "Expiration":
			z.Expiration, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "Inception":
			z.Inception, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "KeyTag":
			z.KeyTag, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "SignerName":
			z.SignerName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Signature":
			z.Signature, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *RRSIG) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 12 + msgp.Uint16Size + 10 + msgp.Uint8Size + 7 + msgp.Uint8Size + 8 + msgp.Uint32Size + 11 + msgp.Uint32Size + 10 + msgp.Uint32Size + 7 + msgp.Uint16Size + 11 + msgp.StringPrefixSize + len(z.SignerName) + 10 + msgp.StringPrefixSize + len(z.Signature)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z RR_Header) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, hdrToEmpty(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RR_Header) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			return
		}
		(*z) = emptyToHDR(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z RR_Header) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(hdrToEmpty(z))
	return
}

// MarshalMsg implements msgp.Marshaler
func (z RT) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Hdr"
	o = append(o, 0x83, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Preference"
	o = append(o, 0xaa, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65)
	o = msgp.AppendUint16(o, z.Preference)
	// string "Host"
	o = append(o, 0xa4, 0x48, 0x6f, 0x73, 0x74)
	o = msgp.AppendString(o, z.Host)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RT) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Preference":
			z.Preference, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Host":
			z.Host, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z RT) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 11 + msgp.Uint16Size + 5 + msgp.StringPrefixSize + len(z.Host)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SIG) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "RRSIG"
	o = append(o, 0x81, 0xa5, 0x52, 0x52, 0x53, 0x49, 0x47)
	o, err = z.RRSIG.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SIG) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "RRSIG":
			bts, err = z.RRSIG.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SIG) Msgsize() (s int) {
	s = 1 + 6 + z.RRSIG.Msgsize()
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SMIMEA) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Hdr"
	o = append(o, 0x85, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Usage"
	o = append(o, 0xa5, 0x55, 0x73, 0x61, 0x67, 0x65)
	o = msgp.AppendUint8(o, z.Usage)
	// string "Selector"
	o = append(o, 0xa8, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72)
	o = msgp.AppendUint8(o, z.Selector)
	// string "MatchingType"
	o = append(o, 0xac, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendUint8(o, z.MatchingType)
	// string "Certificate"
	o = append(o, 0xab, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65)
	o = msgp.AppendString(o, z.Certificate)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SMIMEA) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Usage":
			z.Usage, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Selector":
			z.Selector, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "MatchingType":
			z.MatchingType, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Certificate":
			z.Certificate, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SMIMEA) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 6 + msgp.Uint8Size + 9 + msgp.Uint8Size + 13 + msgp.Uint8Size + 12 + msgp.StringPrefixSize + len(z.Certificate)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SOA) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 8
	// string "Hdr"
	o = append(o, 0x88, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Ns"
	o = append(o, 0xa2, 0x4e, 0x73)
	o = msgp.AppendString(o, z.Ns)
	// string "Mbox"
	o = append(o, 0xa4, 0x4d, 0x62, 0x6f, 0x78)
	o = msgp.AppendString(o, z.Mbox)
	// string "Serial"
	o = append(o, 0xa6, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c)
	o = msgp.AppendUint32(o, z.Serial)
	// string "Refresh"
	o = append(o, 0xa7, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68)
	o = msgp.AppendUint32(o, z.Refresh)
	// string "Retry"
	o = append(o, 0xa5, 0x52, 0x65, 0x74, 0x72, 0x79)
	o = msgp.AppendUint32(o, z.Retry)
	// string "Expire"
	o = append(o, 0xa6, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65)
	o = msgp.AppendUint32(o, z.Expire)
	// string "Minttl"
	o = append(o, 0xa6, 0x4d, 0x69, 0x6e, 0x74, 0x74, 0x6c)
	o = msgp.AppendUint32(o, z.Minttl)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SOA) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Ns":
			z.Ns, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Mbox":
			z.Mbox, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Serial":
			z.Serial, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "Refresh":
			z.Refresh, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "Retry":
			z.Retry, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "Expire":
			z.Expire, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "Minttl":
			z.Minttl, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SOA) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 3 + msgp.StringPrefixSize + len(z.Ns) + 5 + msgp.StringPrefixSize + len(z.Mbox) + 7 + msgp.Uint32Size + 8 + msgp.Uint32Size + 6 + msgp.Uint32Size + 7 + msgp.Uint32Size + 7 + msgp.Uint32Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SPF) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Txt"
	o = append(o, 0xa3, 0x54, 0x78, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Txt)))
	for za0001 := range z.Txt {
		o = msgp.AppendString(o, z.Txt[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SPF) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Txt":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Txt) >= int(zb0003) {
				z.Txt = (z.Txt)[:zb0003]
			} else {
				z.Txt = make([]string, zb0003)
			}
			for za0001 := range z.Txt {
				z.Txt[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SPF) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 4 + msgp.ArrayHeaderSize
	for za0001 := range z.Txt {
		s += msgp.StringPrefixSize + len(z.Txt[za0001])
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SRV) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Hdr"
	o = append(o, 0x85, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Priority"
	o = append(o, 0xa8, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79)
	o = msgp.AppendUint16(o, z.Priority)
	// string "Weight"
	o = append(o, 0xa6, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74)
	o = msgp.AppendUint16(o, z.Weight)
	// string "Port"
	o = append(o, 0xa4, 0x50, 0x6f, 0x72, 0x74)
	o = msgp.AppendUint16(o, z.Port)
	// string "Target"
	o = append(o, 0xa6, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74)
	o = msgp.AppendString(o, z.Target)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SRV) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Priority":
			z.Priority, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Weight":
			z.Weight, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Port":
			z.Port, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Target":
			z.Target, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SRV) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 9 + msgp.Uint16Size + 7 + msgp.Uint16Size + 5 + msgp.Uint16Size + 7 + msgp.StringPrefixSize + len(z.Target)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SSHFP) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Hdr"
	o = append(o, 0x84, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Algorithm"
	o = append(o, 0xa9, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d)
	o = msgp.AppendUint8(o, z.Algorithm)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendUint8(o, z.Type)
	// string "FingerPrint"
	o = append(o, 0xab, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6e, 0x74)
	o = msgp.AppendString(o, z.FingerPrint)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SSHFP) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Algorithm":
			z.Algorithm, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Type":
			z.Type, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "FingerPrint":
			z.FingerPrint, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SSHFP) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 10 + msgp.Uint8Size + 5 + msgp.Uint8Size + 12 + msgp.StringPrefixSize + len(z.FingerPrint)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TA) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Hdr"
	o = append(o, 0x85, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "KeyTag"
	o = append(o, 0xa6, 0x4b, 0x65, 0x79, 0x54, 0x61, 0x67)
	o = msgp.AppendUint16(o, z.KeyTag)
	// string "Algorithm"
	o = append(o, 0xa9, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d)
	o = msgp.AppendUint8(o, z.Algorithm)
	// string "DigestType"
	o = append(o, 0xaa, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendUint8(o, z.DigestType)
	// string "Digest"
	o = append(o, 0xa6, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74)
	o = msgp.AppendString(o, z.Digest)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TA) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "KeyTag":
			z.KeyTag, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Algorithm":
			z.Algorithm, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "DigestType":
			z.DigestType, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Digest":
			z.Digest, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TA) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 7 + msgp.Uint16Size + 10 + msgp.Uint8Size + 11 + msgp.Uint8Size + 7 + msgp.StringPrefixSize + len(z.Digest)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z TALINK) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Hdr"
	o = append(o, 0x83, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "PreviousName"
	o = append(o, 0xac, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.PreviousName)
	// string "NextName"
	o = append(o, 0xa8, 0x4e, 0x65, 0x78, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.NextName)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TALINK) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "PreviousName":
			z.PreviousName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "NextName":
			z.NextName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z TALINK) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 13 + msgp.StringPrefixSize + len(z.PreviousName) + 9 + msgp.StringPrefixSize + len(z.NextName)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TKEY) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "Hdr"
	o = append(o, 0x8a, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Algorithm"
	o = append(o, 0xa9, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d)
	o = msgp.AppendString(o, z.Algorithm)
	// string "Inception"
	o = append(o, 0xa9, 0x49, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendUint32(o, z.Inception)
	// string "Expiration"
	o = append(o, 0xaa, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendUint32(o, z.Expiration)
	// string "Mode"
	o = append(o, 0xa4, 0x4d, 0x6f, 0x64, 0x65)
	o = msgp.AppendUint16(o, z.Mode)
	// string "Error"
	o = append(o, 0xa5, 0x45, 0x72, 0x72, 0x6f, 0x72)
	o = msgp.AppendUint16(o, z.Error)
	// string "KeySize"
	o = append(o, 0xa7, 0x4b, 0x65, 0x79, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendUint16(o, z.KeySize)
	// string "Key"
	o = append(o, 0xa3, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.Key)
	// string "OtherLen"
	o = append(o, 0xa8, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x4c, 0x65, 0x6e)
	o = msgp.AppendUint16(o, z.OtherLen)
	// string "OtherData"
	o = append(o, 0xa9, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61)
	o = msgp.AppendString(o, z.OtherData)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TKEY) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Algorithm":
			z.Algorithm, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Inception":
			z.Inception, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "Expiration":
			z.Expiration, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "Mode":
			z.Mode, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Error":
			z.Error, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "KeySize":
			z.KeySize, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Key":
			z.Key, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "OtherLen":
			z.OtherLen, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "OtherData":
			z.OtherData, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TKEY) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 10 + msgp.StringPrefixSize + len(z.Algorithm) + 10 + msgp.Uint32Size + 11 + msgp.Uint32Size + 5 + msgp.Uint16Size + 6 + msgp.Uint16Size + 8 + msgp.Uint16Size + 4 + msgp.StringPrefixSize + len(z.Key) + 9 + msgp.Uint16Size + 10 + msgp.StringPrefixSize + len(z.OtherData)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TLSA) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Hdr"
	o = append(o, 0x85, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Usage"
	o = append(o, 0xa5, 0x55, 0x73, 0x61, 0x67, 0x65)
	o = msgp.AppendUint8(o, z.Usage)
	// string "Selector"
	o = append(o, 0xa8, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72)
	o = msgp.AppendUint8(o, z.Selector)
	// string "MatchingType"
	o = append(o, 0xac, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendUint8(o, z.MatchingType)
	// string "Certificate"
	o = append(o, 0xab, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65)
	o = msgp.AppendString(o, z.Certificate)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TLSA) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Usage":
			z.Usage, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Selector":
			z.Selector, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "MatchingType":
			z.MatchingType, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Certificate":
			z.Certificate, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TLSA) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 6 + msgp.Uint8Size + 9 + msgp.Uint8Size + 13 + msgp.Uint8Size + 12 + msgp.StringPrefixSize + len(z.Certificate)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TXT) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Txt"
	o = append(o, 0xa3, 0x54, 0x78, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Txt)))
	for za0001 := range z.Txt {
		o = msgp.AppendString(o, z.Txt[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TXT) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Txt":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Txt) >= int(zb0003) {
				z.Txt = (z.Txt)[:zb0003]
			} else {
				z.Txt = make([]string, zb0003)
			}
			for za0001 := range z.Txt {
				z.Txt[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TXT) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 4 + msgp.ArrayHeaderSize
	for za0001 := range z.Txt {
		s += msgp.StringPrefixSize + len(z.Txt[za0001])
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Type) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint16(o, uint16(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Type) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint16
		zb0001, bts, err = msgp.ReadUint16Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Type(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Type) Msgsize() (s int) {
	s = msgp.Uint16Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z UID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Uid"
	o = append(o, 0xa3, 0x55, 0x69, 0x64)
	o = msgp.AppendUint32(o, z.Uid)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Uid":
			z.Uid, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z UID) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 4 + msgp.Uint32Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z UINFO) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Uinfo"
	o = append(o, 0xa5, 0x55, 0x69, 0x6e, 0x66, 0x6f)
	o = msgp.AppendString(o, z.Uinfo)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UINFO) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Uinfo":
			z.Uinfo, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z UINFO) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 6 + msgp.StringPrefixSize + len(z.Uinfo)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *URI) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Hdr"
	o = append(o, 0x84, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "Priority"
	o = append(o, 0xa8, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79)
	o = msgp.AppendUint16(o, z.Priority)
	// string "Weight"
	o = append(o, 0xa6, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74)
	o = msgp.AppendUint16(o, z.Weight)
	// string "Target"
	o = append(o, 0xa6, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74)
	o = msgp.AppendString(o, z.Target)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *URI) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "Priority":
			z.Priority, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Weight":
			z.Weight, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "Target":
			z.Target, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *URI) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 9 + msgp.Uint16Size + 7 + msgp.Uint16Size + 7 + msgp.StringPrefixSize + len(z.Target)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z X25) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Hdr"
	o = append(o, 0x82, 0xa3, 0x48, 0x64, 0x72)
	o = msgp.AppendString(o, hdrToEmpty(z.Hdr))
	// string "PSDNAddress"
	o = append(o, 0xab, 0x50, 0x53, 0x44, 0x4e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
	o = msgp.AppendString(o, z.PSDNAddress)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *X25) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hdr":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Hdr = emptyToHDR(zb0002)
			}
		case "PSDNAddress":
			z.PSDNAddress, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z X25) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(hdrToEmpty(z.Hdr)) + 12 + msgp.StringPrefixSize + len(z.PSDNAddress)
	return
}
