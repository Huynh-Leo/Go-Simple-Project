package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	Controller "gapp.go/Controller"
	Model "gapp.go/Model"
)

func MenuGiangvien() {
	reader := bufio.NewReader(os.Stdin)
	var danhSachGV []Model.GiangVien
	for {
		Controller.ClearScreen()
		fmt.Println("**** Quản lý giảng viên ****")
		fmt.Println("1. Thêm giảng viên")
		fmt.Println("2. Xóa giảng viên")
		fmt.Println("3. Sửa giảng viên")
		fmt.Println("4. Danh sách giảng viên")
		fmt.Println("5. Tìm kiếm giảng viên")
		fmt.Println("6. Quay lại")
		fmt.Println("============== Hãy chọn chức năng ====================")
		line, _ := reader.ReadString('\n')
		line = Controller.TextInt(line)
		switch line {
		case "1":
			fmt.Print("Hãy nhập Id của giảng viên:")
			lineId, _ := reader.ReadString('\n')
			lineId = Controller.TextInt(lineId)

			fmt.Print("Hãy nhập Tên của giảng viên:")
			lineName, _ := reader.ReadString('\n')
			// lineName = Controller.TextInt(lineName)

			fmt.Print("Hãy nhập lớp của giảng viên:")
			lineGrade, _ := reader.ReadString('\n')
			// lineGrade = Controller.TextInt(lineGrade)

			id, _ := strconv.Atoi(lineId)

			sv := Model.GiangVien{
				ID:    id,
				Name:  lineName,
				Grade: lineGrade,
			}
			danhSachGV = append(danhSachGV, sv)
		case "2":
		case "3":
		case "4":
			for _, sv := range danhSachGV {
				fmt.Printf(
					"*ID: %d\n *Name: %s\n *Grade: %s\n",
					sv.ID, sv.Name, sv.Grade,
				)
			}
			reader.ReadString('\n')
		case "5":
		case "6":
			return
		default:
			fmt.Println("Hãy nhập lại dữ liệu")
		}
	}

}
