package person

type student struct {
	Name string
	Role string
}

// func (student student) String() string {
// 	return "i am " + student.Name + " is a " + student.Role
// }

func Test1(p *student) {
	p.Name = "C"
}

func NewStudent(name string) student {
	return student{Name: name, Role: "student"}
}
