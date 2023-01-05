package person

type teacher struct {
	Name string
	Role string
}

// func (teacher teacher) String() string {
// 	return "i am " + teacher.Name + " is a " + teacher.Role
// }

func Test2(t *teacher) {
	t.Name = "C"
}

func NewTeacher(name string) *teacher {
	return &teacher{Name: name, Role: "teacher"}
}
