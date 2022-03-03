package models

type User struct {
	ID        uint   `json:"id"form:"id"`
	StudentId uint64 `json:"student_id"form:"student_id"validate:"required"` //学号

}

func CreateUserByStudentId(studentId uint64) error {
	u := User{StudentId: studentId}
	DB.Create(&u)
	return nil
}

func GetUserIdByStudentId(studentId uint64) (userId uint, err error) {
	var u User
	DB.Where("student_id =?", studentId).First(&u)
	if u.ID == 0 {
		_ = CreateUserByStudentId(studentId)
	}
	DB.Where("student_id =?", studentId).First(&u)
	return userId, nil

}
