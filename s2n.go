package stringx

import "strconv"

func ParseUint(s string) uint {
	i, _ := strconv.ParseUint(s, 10, 0)
	return uint(i)
}

func ParseUint8(s string) uint8 {
	i, _ := strconv.ParseUint(s, 10, 8)
	return uint8(i)
}

func ParseUint16(s string) uint16 {
	i, _ := strconv.ParseUint(s, 10, 16)
	return uint16(i)
}

func ParseUint32(s string) uint32 {
	i, _ := strconv.ParseUint(s, 10, 32)
	return uint32(i)
}

func ParseUint64(s string) uint64 {
	i, _ := strconv.ParseUint(s, 10, 64)
	return i
}

func ParseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ParseInt8(s string) int8 {
	i, _ := strconv.ParseInt(s, 10, 8)
	return int8(i)
}

func ParseInt16(s string) int16 {
	i, _ := strconv.ParseInt(s, 10, 16)
	return int16(i)
}

func ParseInt32(s string) int32 {
	i, _ := strconv.ParseInt(s, 10, 32)
	return int32(i)
}

func ParseInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func ParseBool(s string) bool {
	i, _ := strconv.ParseBool(s)
	return i
}

func ParseFloat32(s string) float32 {
	i, _ := strconv.ParseFloat(s, 32)
	return float32(i)
}

func ParseFloat64(s string) float64 {
	i, _ := strconv.ParseFloat(s, 32)
	return i
}
