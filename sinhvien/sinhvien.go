package sinhvien

type SinhVien struct {
	ID   int
	Name string
}

//	:= sinhVien{
//		ID:1,
//		Name:"Vinh"
//	}
func (sv SinhVien) PrintValue() string {
	return sv.Name
}
