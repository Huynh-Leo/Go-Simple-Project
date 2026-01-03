package sinhvien

type SinhVien struct {
	ID    int
	Name  string
	Year  int
	Grade string
}

func (sv SinhVien) PrintValue() string {
	return sv.Name
}
