package studentInfo

import "fmt"

// StudentInfo
type StudentInfo struct {
	name      string
	studentID string
	status    int
}

// NewStudent creates studentInfo
func NewStudent(name string) *StudentInfo {
	student := StudentInfo{name: name, studentID: "", status: 0}
	return &student
}

// The first step to create student's own id
func (s *StudentInfo) IdNumber(year string) {
	s.studentID += year
}

// Check the id
func (s StudentInfo) IdCheck() string {
	return s.studentID
}

// return name
func (s StudentInfo) Student() string {
	return s.name
}

// if status==1 : 000
func (s *StudentInfo) Status(status int) {
	if status == 1 {
		s.studentID += "000"
	} else {
		s.studentID += "110"
	}
}

// return the id
func (s StudentInfo) StatusCheck() string {
	return s.studentID
}

func (s StudentInfo) String() string {
	return fmt.Sprint(s.Student(), "'s studentID is ", s.IdCheck())
}
