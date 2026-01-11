package view

import (
	"bufio"
	"fmt"
	"os"

	Controller "gapp.go/Controller"
	controller "gapp.go/Controller"
)

func MenuGiangvien() {
	reader := bufio.NewReader(os.Stdin)
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
			Controller.AddGiangVien()
		case "2":
			Controller.DeleteGiangVien()
		case "3":
			controller.EditGiangVien()
		case "4":
			Controller.ViewGiangVien()
		case "5":
			Controller.FindGiangVien()
		case "6":
			return
		default:
			fmt.Println("Hãy nhập lại dữ liệu")
		}
	}

}
