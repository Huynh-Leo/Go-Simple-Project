package view

import (
	"bufio"
	"fmt"
	"os"

	Controller "gapp.go/Controller"
	controller "gapp.go/Controller"
)

func MenuSinhvien() {
	reader := bufio.NewReader(os.Stdin)
	for {
		Controller.ClearScreen()

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
		line = Controller.TextInt(line)
		switch line {
		case "1":
			Controller.AddSinhVien()
		case "2":
			Controller.DeleteSinhVien()
		case "3":
			controller.EditSinhVien()
		case "4":
			Controller.ViewSinhVien()
		case "5":
			Controller.FindSinhVien()
		case "6":
			return
		default:
			fmt.Println("Hãy nhập lại dữ liệu")
		}
	}
}
