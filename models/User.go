package models

type User struct {
	ID        uint   `json:"id"form:"id"`
	StudentId string `json:"student_id"form:"student_id"validate:"required"` //学号

}

//通过学号创建用户,返回新Id
func CreateUserByStudentId(studentId interface{}) (userId uint, err error) {
	u := User{StudentId: studentId.(string)}
	DB.Create(&u)
	return u.ID, nil
}

//通过学号获取用户Id,不存在的话则会创建一条用户数据
func MustGetUserIdByStudentId(studentId string) (userId uint, err error) {
	var u User
	DB.Where("student_id =?", studentId).First(&u)
	//此学号的用户不存在,则创建用户
	if u.ID == 0 {
		userId, _ = CreateUserByStudentId(studentId)
		return userId, nil
	}
	return u.ID, nil
}

//通过学号获取用户Id,可能为0
func GetUserIdByStudentId(studentId string) (userId uint, err error) {
	var u User
	DB.Where("student_id =?", studentId).First(&u)

	return u.ID, nil
}
