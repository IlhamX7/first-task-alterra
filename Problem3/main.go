package main

import (
	"fmt"
)

type Student struct {
	name  string
	score int
}

type StudentList struct {
	students []Student
}

func (s *StudentList) AddStudent(name string, score int) {
	s.students = append(s.students, Student{name: name, score: score})
}

func (s StudentList) Average() float64 {
	totalScore := 0
	for _, student := range s.students {
		totalScore += student.score
	}
	return float64(totalScore) / float64(len(s.students))
}

func (s StudentList) Min() (min int, name string) {
	minStudent := s.students[0]
	for _, student := range s.students {
		if student.score < minStudent.score {
			minStudent = student
		}
	}
	return minStudent.score, minStudent.name
}

func (s StudentList) Max() (max int, name string) {
	maxStudent := s.students[0]
	for _, student := range s.students {
		if student.score > maxStudent.score {
			maxStudent = student
		}
	}
	return maxStudent.score, maxStudent.name
}

func main() {
	var students StudentList

	for i := 1; i <= 5; i++ {
		var name string
		var score int

		fmt.Printf("Input %d Student's Name: ", i)
		fmt.Scan(&name)
		fmt.Printf("Input %d Student's Score: ", i)
		fmt.Scan(&score)

		students.AddStudent(name, score)
	}

	averageScore := students.Average()

	minScore, minName := students.Min()
	maxScore, maxName := students.Max()

	fmt.Printf("\nAverage Score: %.2f\n", averageScore)
	fmt.Printf("Min Score of Students: %s(%d)\n", minName, minScore)
	fmt.Printf("Max Score of Students: %s(%d)\n", maxName, maxScore)
}
