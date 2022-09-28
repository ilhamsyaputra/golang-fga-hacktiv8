// Identifying Method Information
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Grade int
}

func (s *Student) SetName(name string, grade int) {
	s.Name = name
	s.Grade = grade
}

func main() {
	var student1 = &Student{Name: "Ilham Syaputra", Grade: 3}
	fmt.Println(student1.Name)

	reflectValue := reflect.ValueOf(student1)
	method := reflectValue.MethodByName("SetName")

	/*
		merubah property dari student dapat dilakukan dengan merefleksi method untuk merubah property (SetName) Line 14
		kemudian di refleksikan dan di simpan ke variable `method` Line24
		Selanjutnya jalankan method .Call dari variable `method`, argumen harus berupa reflect array dan berurutan sesuai deklarasi parameternya Line 32-35
	*/

	method.Call([]reflect.Value{
		reflect.ValueOf("Ilham"),
		reflect.ValueOf(1),
	})
	fmt.Println(student1.Name)
	fmt.Println(student1.Grade)
}
