package stringx

import (
	"reflect"
	"strconv"
)

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

func FormatFloat64(i float64) string {
	return strconv.FormatFloat(i, 'f', -1, 64)
}

func FormatFloat32(i float32) string {
	return strconv.FormatFloat(float64(i), 'f', -1, 32)
}

func Int2S[T Int](i T) string {
	return strconv.FormatInt(int64(i), 10)
}

func Uint2S[T Uint](i T) string {
	return strconv.FormatInt(int64(i), 10)
}

func N2S[T Int | Uint](i T) string {
	return strconv.FormatInt(int64(i), 10)
}

func Float2S[T Float](i T) string {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Float32:
		return strconv.FormatFloat(float64(i), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(float64(i), 'f', -1, 64)
	}
	return ""
}
