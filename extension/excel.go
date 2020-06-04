// +----------------------------------------------------------------------
// | Excel 方法
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年04月02日
// +----------------------------------------------------------------------

package gohelp

func Index2ExcelRow(index int) string {
	var Letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	result := Letters[index%26]
	index = index / 26
	for index > 0 {
		index = index - 1
		result = Letters[index%26] + result
		index = index / 26
	}
	return result
}
