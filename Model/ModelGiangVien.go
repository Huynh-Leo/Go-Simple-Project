package model

type GiangVien struct {
	ID    int
	Name  string
	Grade string
}

func (sv GiangVien) PrintValue() string {
	return sv.Name
}
