package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	Model "gapp.go/Model"
)

var danhSachSV []Model.SinhVien

func AddSinhVien() {
	reader := bufio.NewReader(os.Stdin)

	var lineId string
	for {
		fmt.Print("Hãy nhập Id của sinh viên:")
		lineId, _ = reader.ReadString('\n')
		if !TextIntId(danhSachSV, lineId) {
			fmt.Println("ID không được trùng với ID có trong dữ liệu")
			continue
		}
		break
	}
	/*
		Vì khi nhập giá trị thì sẽ có \n kế bên dữ liệu nhập
		để tránh dữ liệu khi chuyển đổi thành 0
		thì phải loại bỏ để có thể chuyển từ string sang int
	*/
	lineId = NormalizeInput(lineId)
	id, _ := strconv.Atoi(lineId)

	// Đặt tên biến nhập
	var lineName string
	for {
		fmt.Print("Hãy nhập Tên của sinh viên: ")
		// Đọc tất cả dữ liệu nhập cho đến khi Enter
		lineName, _ = reader.ReadString('\n')
		// Nếu giá trị trả về là true thì sẽ break
		// lúc đó dữ liệu sẽ được ghi vào biến
		if TextStringName(lineName) {
			break
		}
		fmt.Println("❌ Tên không hợp lệ, vui lòng nhập lại!")
	}

	fmt.Print("Hãy nhập Năm của sinh viên:")
	lineYear, _ := reader.ReadString('\n')
	// lineYear = TextStringName(lineYear)

	fmt.Print("Hãy nhập lớp của sinh viên:")
	lineGrade, _ := reader.ReadString('\n')

	// Convert từ string sang lại int
	year, _ := strconv.Atoi(lineYear)

	sv := Model.SinhVien{
		ID:    id,
		Name:  lineName,
		Year:  year,
		Grade: lineGrade,
	}

	danhSachSV = append(danhSachSV, sv)
}

func DeleteSinhVien() {

}
func EditSinhVien() {

}
func ViewSinhVien() {
	reader := bufio.NewReader(os.Stdin)
	for _, sv := range danhSachSV {
		fmt.Printf(
			"*ID: %d\n *Name: %s\n *Year: %d\n *Grade: %s\n",
			sv.ID, sv.Name, sv.Year, sv.Grade,
		)
	}
	// Nhằm giữ màn hình lại để không clear ngay khi hiện kết quả
	// Chỉ khi nhấn thêm Enter sẽ kết thúc Case này và quay lại switch
	reader.ReadString('\n')
}
func FindSinhVien() {

}
