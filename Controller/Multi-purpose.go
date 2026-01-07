package controller

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"unicode"

	Model "gapp.go/Model"
)

// chuyển mọi input thành string
func NormalizeInput(value interface{}) string {
	/* 	fmt.sprint chuyển bất kỳ kiểu nào thành string
	   	strings.Trimspace loại bỏ space/tab đầu và cuối */
	return strings.TrimSpace(fmt.Sprint(value))
}

/*
Xóa bỏ toàn bộ khoảng trắng ở giữa
hàm nhận kiểu dữ liệu string và trả về kiểu dữ liệu string
*/
func EraseSpace(value string) string {
	str := NormalizeInput(value)
	/* 	Hàm strings.ReplaceAll thay tất cả chuỗi con trong một chuỗi
	   	trong trường hợp này thay tất cả khoảng trắng " " thành "" */
	str = strings.ReplaceAll(str, " ", "")
	return str
}

// Hàm kiểm tra có phải là số nguyên không
func TextInt(value string) string {
	value = EraseSpace(value) // gán kết quả của hàm vào value
	// Kiểm tra trong chuỗi chỉ được có số
	for _, num := range value {
		if !unicode.IsDigit(num) || num <= 0 {
			fmt.Println("Dữ liệu không hợp lệ")
			return ""
		}
	}
	return value
}

/*
Vì biến TextInt trả về là value
nên sẽ tạo thêm 1 phương thức tương tự TextInt nhưng trả về bool
*/
func TextIntBool(value string) bool {
	value = EraseSpace(value)
	for _, num := range value {
		if !unicode.IsDigit(num) || num <= 0 {
			return false
		}
	}
	return true
}

// Hàm kiểm tra có phải là chữ không
func TextStringName(value string) bool {
	// Hàm loại bỏ khoảng trắng đầu và cuối của chuỗi
	value = NormalizeInput(value)
	// vòng lặp for duyệt từng ký tự
	// r là rune( ký tự Unicode) dùng để  hỗ kiểm tra Tiếng Việt
	for _, err := range value {
		// Nếu ký tự hiện tại là dấu space thì chuyển sang ký tự kế
		if err == ' ' {
			continue
		}
		// Nếu ký tự thuốc số hoặc ký tự đặc biệt về false
		if !unicode.IsLetter(err) {
			// fmt.Println("Tên không hợp lệ")
			return false
		}
	}
	/* 	len(value) trả về độ dài của chuỗi
	   	nếu chuối lớn hơn 0 và do kiểu dữ liệu trả về là bool
	   	nên return sẽ là true
	   	Kết hợp giữa việc xét chuỗi phải lớn hơn 0 và trả về true */
	return len(value) > 0
}

// Hàm kiểm tra ID đầu vào
func TextIntId(danhSachSV []Model.SinhVien, Id string) bool {
	if TextIntBool(Id) {
		id, _ := strconv.Atoi(Id)
		for _, sv := range danhSachSV {
			if sv.ID == id {
				return false // nếu trùng id thì trả về false
			}
		}
		return true // không trùng id nào thì trả về true
	}
	/* Nếu dữ liệu nhập khi qua TextIntBool(Id)
	không đạt điều kiện thì sẽ return false */
	return false
}

// Hàm kiểm tra năm nhập vào
func TextIntYear(Year string) bool {
	Year = TextInt(Year)
	/*
		Điều kiện chung giá trị nhập vào chỉ được 4 chữ số
		điều kiện nếu số đầu là chỉ được 19 hoặc 20
	*/
	re := regexp.MustCompile(`^(19\d{2}|20\d{2})$`)
	return re.MatchString(Year)
}

/*
Hàm clear màn hình nhằm khi switch case
thì sẽ xóa đoạn hiển thị ban đầu
*/
func ClearScreen() {
	// khai báo biến cmd để chứa lệnh hệ thống cần chạy
	var cmd *exec.Cmd
	/* 	Kiểm tra hệ điều hành đang chạy chương trình
	   	runtime.GOOS trả về ở đây là window */
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
