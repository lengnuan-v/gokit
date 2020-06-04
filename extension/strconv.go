// +----------------------------------------------------------------------
// | strconv方法
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年04月02日
// +----------------------------------------------------------------------

package gohelp

import "strconv"

// string转成int
func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// string转成int64
func StringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// int转成string：
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// int64转成string
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// string转float
// size 32 -> float32 64 -> float64
func StringToFloat(str string, size int) (float64, error) {
	return strconv.ParseFloat(str, size)
}

// float转string
// size 32 -> float32 64 -> float64
func FloatToString(f float64 , size int) string {
	return strconv.FormatFloat(f, 'f', -1, size)
}
