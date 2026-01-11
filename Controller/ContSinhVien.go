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
		/*
			CHỈ ĐƯỢC GÁN GIÁ TRỊ =
			KHÔNG KHAI BÁO VÀ GÀN GIÁ TRỊ :=
			SẼ LÀM CHO BIẾN BÊN NGOÀI VÒNG LÒNG = ""
		*/
		lineId, _ = reader.ReadString('\n')
		if !TextIntIdSV(danhSachSV, NormalizeInput(lineId)) {
			fmt.Println("ID không được trùng với ID có trong dữ liệu")
			continue
		} else if NormalizeInput(lineId) == "" {
			fmt.Println("ID không được để trống")
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

	var lineYear string
	for {
		fmt.Print("Hãy nhập Năm của sinh viên:")
		lineYear, _ = reader.ReadString('\n')
		if TextIntYear(NormalizeInput(lineYear)) {
			break
		}
		fmt.Println("❌ Năm không hợp lệ, vui lòng nhập lại!")
		continue
	}
	lineYear = NormalizeInput(lineYear)
	year, _ := strconv.Atoi(lineYear)

	var lineGrade string
	for {
		fmt.Print("Hãy nhập lớp của sinh viên:")
		lineGrade, _ = reader.ReadString('\n')
		if TextStringClass(NormalizeInput(lineGrade)) {
			break
		}
		fmt.Println("❌ Lớp không hợp lệ, vui lòng nhập lại!")
		continue
	}

	sv := Model.SinhVien{
		ID:    id,
		Name:  lineName,
		Year:  year,
		Grade: lineGrade,
	}

	danhSachSV = append(danhSachSV, sv)
}

func DeleteSinhVien() {
	var valId string
	reader := bufio.NewReader(os.Stdin)

	idx := 0
	for {
		fmt.Println("Hãy nhập ID cần xóa thông tin")
		valId, _ = reader.ReadString('\n')
		if TextIntIdSV(danhSachSV, NormalizeInput(valId)) {
			fmt.Println("ID không tồn tại ")
			continue
		}
		id, _ := strconv.Atoi(NormalizeInput(valId))
		for _, user := range danhSachSV {
			if user.ID != id {
				danhSachSV[idx] = user
				idx++
			}
		}
		danhSachSV = danhSachSV[:idx]
		break
	}
}
func EditSinhVien() {
	var valId string
	reader := bufio.NewReader(os.Stdin)
	ViewSinhVien()
	for {
		fmt.Println("Hãy nhập ID cần chỉnh sửa thông tin")
		valId, _ = reader.ReadString('\n')
		if TextIntIdSV(danhSachSV, NormalizeInput(valId)) {
			fmt.Println("ID không tồn tại ")
			continue
		}
		AddSinhVien()
		break
	}
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
	var valId string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Hãy nhập ID để tìm kiếm thông tin")
		valId, _ = reader.ReadString('\n')
		if TextIntIdSV(danhSachSV, NormalizeInput(valId)) {
			fmt.Println("ID không tồn tại ")
			continue
		}
		id, _ := strconv.Atoi(NormalizeInput(valId))
		for _, sv := range danhSachSV {
			if sv.ID == id {
				fmt.Println("ID:", sv.ID)
				fmt.Println("Name:", sv.Name)
				fmt.Println("Year:", sv.Year)
				fmt.Println("Grade:", sv.Grade)
			}
		}
		reader.ReadString('\n')
		break
	}

}
