package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	Model "gapp.go/Model"
)

var danhSachGV []Model.GiangVien

func AddGiangVien() {
	reader := bufio.NewReader(os.Stdin)

	var lineId string
	for {
		fmt.Print("Hãy nhập Id của giảng viên:")
		lineId, _ = reader.ReadString('\n')
		if !TextIntIdGV(danhSachGV, NormalizeInput(lineId)) {
			fmt.Println("ID không được trùng với ID có trong dữ liệu")
			continue
		} else if NormalizeInput(lineId) == "" {
			fmt.Println("ID không được để trống")
			continue
		}
		break
	}
	lineId = NormalizeInput(lineId)
	id, _ := strconv.Atoi(lineId)
	var lineName string
	for {
		fmt.Print("Hãy nhập Tên của giảng viên: ")
		lineName, _ = reader.ReadString('\n')
		if TextStringName(lineName) {
			break
		}
		fmt.Println("❌ Tên không hợp lệ, vui lòng nhập lại!")
	}

	var lineGrade string
	for {
		fmt.Print("Hãy nhập lớp đảm nhiệm của giảng viên:")
		lineGrade, _ = reader.ReadString('\n')
		if TextStringClass(NormalizeInput(lineGrade)) {
			break
		}
		fmt.Println("❌ Lớp không hợp lệ, vui lòng nhập lại!")
		continue
	}

	gv := Model.GiangVien{
		ID:    id,
		Name:  lineName,
		Grade: lineGrade,
	}

	danhSachGV = append(danhSachGV, gv)
}
func DeleteGiangVien() {
	var valId string
	reader := bufio.NewReader(os.Stdin)
	idx := 0
	for {
		fmt.Println("Hãy nhập ID cần xóa thông tin")
		valId, _ = reader.ReadString('\n')
		if TextIntIdGV(danhSachGV, NormalizeInput(valId)) {
			fmt.Println("ID không tồn tại ")
			continue
		}
		id, _ := strconv.Atoi(NormalizeInput(valId))
		for _, user := range danhSachGV {
			if user.ID != id {
				danhSachGV[idx] = user
				idx++
			}
		}
		danhSachGV = danhSachGV[:idx]
		break
	}
}
func EditGiangVien() {
	var valId string
	reader := bufio.NewReader(os.Stdin)
	ViewGiangVien()
	for {
		fmt.Println("Hãy nhập ID cần chỉnh sửa thông tin")
		valId, _ = reader.ReadString('\n')
		if TextIntIdGV(danhSachGV, NormalizeInput(valId)) {
			fmt.Println("ID không tồn tại ")
			continue
		}
		AddGiangVien()
		break
	}
}
func ViewGiangVien() {
	reader := bufio.NewReader(os.Stdin)
	for _, gv := range danhSachGV {
		fmt.Printf(
			"*ID: %d\n *Name: %s\n  *Grade: %s\n",
			gv.ID, gv.Name, gv.Grade,
		)
	}
	reader.ReadString('\n')
}
func FindGiangVien() {
	var valId string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Hãy nhập ID để tìm kiếm thông tin")
		valId, _ = reader.ReadString('\n')
		if TextIntIdGV(danhSachGV, NormalizeInput(valId)) {
			fmt.Println("ID không tồn tại ")
			continue
		}
		id, _ := strconv.Atoi(NormalizeInput(valId))
		for _, sv := range danhSachGV {
			if sv.ID == id {
				fmt.Println("ID:", sv.ID)
				fmt.Println("Name:", sv.Name)
				fmt.Println("Grade:", sv.Grade)
			}
		}
		reader.ReadString('\n')
		break
	}

}
