package stringx

import "strconv"

func FormatUint(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}

func FormatUint8(i uint8) string {
	return strconv.FormatUint(uint64(i), 10)
}

func FormatUint16(i uint16) string {
	return strconv.FormatUint(uint64(i), 10)
}

func FormatUint32(i uint32) string {
	return strconv.FormatUint(uint64(i), 10)
}

func FormatUint64(i uint64) string {
	return strconv.FormatUint(uint64(i), 10)
}

func FormatInt(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatInt8(i int8) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatInt16(i int16) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatInt32(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatInt64(i int64) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatBool(b bool) string {
	return strconv.FormatBool(b)
}
