package sinhvien

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	function "gapp.go/Function"
)

var danhSachSV []SinhVien

func AddSinhVien() {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Hãy nhập Id của sinh viên:")
	// lineId, _ := reader.ReadString('\n')
	// lineId = function.TextInt(lineId)

	// Đặt tên biến nhập
	var lineName string
	for {
		fmt.Print("Hãy nhập Tên của sinh viên: ")
		// Đọc tất cả dữ liệu nhập cho đến khi Enter
		lineName, _ = reader.ReadString('\n')
		// Nếu giá trị trả về là true thì sẽ break
		// lúc đó dữ liệu sẽ được ghi vào biến
		if function.TextStringName(lineName) {
			break
		}
		fmt.Println("❌ Tên không hợp lệ, vui lòng nhập lại!")
	}

	fmt.Print("Hãy nhập Năm của sinh viên:")
	lineYear, _ := reader.ReadString('\n')
	// lineYear = function.TextStringName(lineYear)

	fmt.Print("Hãy nhập lớp của sinh viên:")
	lineGrade, _ := reader.ReadString('\n')

	// Convert từ string sang lại int
	id, _ := strconv.Atoi(lineId)
	year, _ := strconv.Atoi(lineYear)

	sv := SinhVien{
		ID:    id,
		Name:  lineName,
		Year:  year,
		Grade: lineGrade,
	}
	danhSachSV = append(danhSachSV, sv)
}

func MenuSinhvien() {
	reader := bufio.NewReader(os.Stdin)
	for {
		function.ClearScreen()

		fmt.Println("**** Quản lý sinh viên ****")
		fmt.Println("1. Thêm sinh viên")
		fmt.Println("2. Xóa sinh viên")
		fmt.Println("3. Sửa sinh viên")
		fmt.Println("4. Danh sách sinh viên")
		fmt.Println("5. Tìm kiếm sinh viên")
		fmt.Println("6. Quay lại")
		fmt.Println("============== Hãy chọn chức năng ====================")
		fmt.Print("Hãy nhập chức năng:")
		line, _ := reader.ReadString('\n')
		line = function.TextInt(line)
		switch line {
		case "1":
			AddSinhVien()
		case "2":
		case "3":
		case "4":
			for _, sv := range danhSachSV {
				fmt.Printf(
					"*ID: %d\n *Name: %s\n *Year: %d\n *Grade: %s\n",
					sv.ID, sv.Name, sv.Year, sv.Grade,
				)
			}
			// Nhằm giữ màn hình lại để không clear ngay khi hiện kết quả
			// Chỉ khi nhấn thêm Enter sẽ kết thúc Case này và quay lại switch
			reader.ReadString('\n')
		case "5":
		case "6":
			return
		default:
			fmt.Println("Hãy nhập lại dữ liệu")
		}
	}
}

func DeleteSinhVien() {

}
func EditSinhVien() {

}
func ViewSinhVien() {

}
func FindSinhVien() {

}
