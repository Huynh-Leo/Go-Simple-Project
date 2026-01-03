package giangvien

type giangvien struct {
	ID    int
	Name  string
	Grade string
}

func (sv giangvien) PrintValue() string {
	return sv.Name
}
