// package stringx

// func padStart(s string, length int, pad string) string {
// 	if len(s) >= length {
// 		return s
// 	}
// 	l := len(pad)
// 	if l == 0 {
// 		return s
// 	}
// 	n := length - len(s)
// 	if n <= l {
// 		return pad[:n] + s
// 	}
// 	for i := 0; i < n; i += l {
// 		s = pad + s
// 	}
// 	return s
// }
