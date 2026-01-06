package main

import (
	"bufio"
	"fmt"
	"os"

	Controller "gapp.go/Controller"
	View "gapp.go/View"
)

func main() {

MenuBandau:
	for {
		Controller.ClearScreen()
		// Tạo bộ đọc dữ liệu nhập
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("CHƯƠNG TRÌNH QUẢN LÝ")
		fmt.Println("1. Quản lý sinh viên")
		fmt.Println("2. Quản lý giảng viên")
		fmt.Println("3. Thoát")
		fmt.Println("============== Hãy chọn chức năng ====================")
		// bufio.NewReader nhằm đọc hết cả dòng cho đến khi kết thúc bằng Enter
		fmt.Print("Hãy nhập chức năng:")
		line, _ := reader.ReadString('\n')
		line = Controller.TextInt(line)
		switch line {
		case "1":
			// sinhvien.MenuSinhvien()
			View.MenuSinhvien()
		case "2":
			View.MenuGiangvien()
		case "3":
			break MenuBandau
		}
	}
	fmt.Println("Đã thoát vòng lặp")

}
