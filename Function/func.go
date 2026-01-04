package function

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

// Hàm kiểm tra có phải là số nguyên không
func TextInt(value string) string {
	value = EraseSpace(value) // gán kết quả của hàm vào value
	// Kiểm tra trong chuỗi chỉ được có số
	for _, num := range value {
		if !unicode.IsDigit(num) || num == 0 {
			fmt.Println("Dữ liệu không hợp lệ")
			return ""
		}
	}
	return value
}

// Hàm clear màn hình nhằm khi switch case
// thì sẽ xóa đoạn hiển thị ban đầu
func ClearScreen() {
	// khai báo biến cmd để chứa lệnh hệ thống cần chạy
	var cmd *exec.Cmd
	// Kiểm tra hệ điều hành đang chạy chương trình
	// runtime.GOOS trả về ở đây là window
	switch runtime.GOOS {
	case "windows":
		// Dùng lệnh cmd /c cls để xóa màn hình
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		// Với những hệ điều hành khác chạy lệnh clear
		cmd = exec.Command("clear")
	}
	// Gửi kết quả lệnh về terminal hiện tại
	cmd.Stdout = os.Stdout
	// thực thi lệnh
	if err := cmd.Run(); err != nil {
		fmt.Println("Error clear screan: ", err)
	}
}
