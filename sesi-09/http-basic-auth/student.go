package main

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func SelectedStudent(id string) *Student {
	for _, student := range students {
		if student.Id == id {
			return student
		}
	}
	return nil
}

func init() {
	students = append(students, &Student{Id: "s001", Name: "M. Ilham Syaputra", Grade: 6})
	students = append(students, &Student{Id: "s002", Name: "Rizky Wardana", Grade: 4})
	students = append(students, &Student{Id: "s003", Name: "Vionita Laura Assyifaa", Grade: 2})
	students = append(students, &Student{Id: "s004", Name: "Nadine Alyssaa Azzahra", Grade: 1})
}
