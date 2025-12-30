package main

import (
	"fmt"
)

func main() {

	for {
		var bienVao int
		fmt.Println("CHƯƠNG TRÌNH QUẢN LÝ")
		fmt.Println("1. Quản lý sinh viên")
		fmt.Println("2. Quản lý giảng viên")
		fmt.Println("3. Thoát")
		fmt.Println("Hãy chọn chức năng")
		// Kiểm tra giá trị nhập vào
		val, err := fmt.Scan(&bienVao)
		// Nếu giá trị là số hoặc bé hơn 1 và lớn hơn 3 thì chạy
		if err != nil || val < 1 || val > 3 {
			fmt.Println("Hãy nhập lại")
		}
		switch bienVao {
		case 1:
			var bienSinhvien int
			fmt.Println("**** Quản lý sinh viên ****")
			fmt.Println("1. Thêm sinh viên")
			fmt.Println("2. Xóa sinh viên")
			fmt.Println("3. Sửa sinh viên")
			fmt.Println("4. Danh sách sinh viên")
			fmt.Println("5. Tìm kiếm sinh viên")
			fmt.Println("6. Quay lại")
			val, err := fmt.Scan(&bienSinhvien)
			if err != nil || val < 1 || val > 6 {
				fmt.Println("Hãy nhập lại")
			} else {
				switch bienSinhvien {
				case 1:
					fmt.Println("xong bước 1")
				case 2:
				case 3:
				case 4:
				case 5:
				case 6:
				}
			}
			// sv := sinhvien.SinhVien{
			// 	ID:   1,
			// 	Name: "Vinh",
			// }
			// fmt.Println(sv.PrintValue())
		case 2:
			fmt.Println("1")
		case 3:
			fmt.Println("1")
		default:
			fmt.Println("Hãy nhập lại")
		}
	}

}
