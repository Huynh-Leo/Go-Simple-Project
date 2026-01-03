package function

import (
	"fmt"
	"strings"
	"unicode"
)

// chuyển mọi input thành string
func NormalizeInput(value interface{}) string {
	// fmt.sprint chuyển bất kỳ kiểu nào thành string
	// strings.Trimspace loại bỏ space/tab đầu và cuối
	return strings.TrimSpace(fmt.Sprint(value))
}

// Xóa bỏ toàn bộ khoảng trắng ở giữa
// hàm nhận kiểu dữ liệu string và trả về kiểu dữ liệu string
func EraseSpace(value string) string {
	str := NormalizeInput(value)
	// Hàm strings.ReplaceAll thay tất cả chuỗi con trong một chuổi
	// trong trường hợp này thay tất cả khoảng trắng " " thành ""
	str = strings.ReplaceAll(str, " ", "")
	return str
}
func TextInt(value string) string {
	value = EraseSpace(value) // gán kết quả của hàm vào value
	// Kiểm tra trong chuỗi chỉ được có số
	for _, num := range value {
		if !unicode.IsDigit(num) {
			fmt.Println("Dữ liệu không hợp lệ")
			return ""
		}
	}
	return value
}
