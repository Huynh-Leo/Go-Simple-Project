package sinhvien

type SinhVien struct {
	ID    int    `json:"id"`
	Name  string `json:"ten"`
	Year  int    `json:"nam"`
	Grade string `json:"lop"`
}

func (sv SinhVien) PrintValue() string {
	return sv.Name
}
