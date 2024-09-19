package models

import (
	"time"

	"github.com/google/uuid"
)

// Teacher models

type Teacher struct {
	TeacherID uuid.UUID `json:"teacher_id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

// Course models

type Course struct {
	CourseID   uuid.UUID `json:"course_id"`
	CourseName string    `json:"course_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"update_at"`
}

// Group models

type Group struct {
	GroupID   uuid.UUID `json:"group_id"`
	GroupName string    `json:"group_name"`
	CourseID  uuid.UUID `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Subject models

type Subject struct {
	SubjectID   uuid.UUID `json:"subject_id"`
	SubjectName string    `json:"subject_name"`
	TeacherID   uuid.UUID `json:"teacher_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Student models

type Student struct {
	StudentID uuid.UUID `json:"student_id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	GroupID   uuid.UUID `json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Grade models

type Grade struct {
	GradeID    uuid.UUID `json:"grade_id"`
	GradeName  string    `json:"grade_name"`
	GradeValue string    `json:"grade_value"`
	GradeDate  time.Time `json:"grade_date"`
	SubjectID  uuid.UUID `json:"subject_id"`
	GroupID    uuid.UUID `json:"group_id"`
	StudentID  uuid.UUID `json:"student_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// GetList Req models

type GetListReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// GetTeacherList

type GetTeacherList struct {
	Teachers []*Teacher `json:"name"`
	Count    int        `json:"count"`
}

// TeacherCreateReq models

type TeacherCreateReq struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TeacherUpdateReq models

type TeacherUpdateReq struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// GetCoursesList models

type GetCoursesList struct {
	Courses []*Course `json:"courses"`
	Count   int       `json:"count"`
}

// GetGroupsList models

type GetGroupsList struct {
	Groups []*Group `json:"groups"`
	Count  int      `json:"count"`
}

// GetSubjectsList models

type GetSubjectsList struct {
	Subjects []*Subject `json:"subjects"`
	Count    int        `json:"count"`
}

// GetStudentsList models

type GetStudentsList struct {
	Students []*Student `json:"students"`
	Count    int        `json:"count"`
}

// GetGradesList models

type GetGradesList struct {
	Grades []*Grade `json:"grades"`
	Count  int      `json:"count"`
}
