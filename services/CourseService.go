package services

import (
	"Course/configs"
	"Course/entity"
	"Course/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

func CreateCourse(course *models.Course) (err error, courseInter *models.Course) {
	err = configs.DB.Create(course).Error
	return err, course
}

func GetCourseById(id int64) (course models.Course, err error) {
	//configs.DB.Find(course, if).Error
	err = configs.DB.First(&course, id).Error
	return
}

func GetCourse(request entity.GetCourseRequest) (response entity.GetCourseResponse) {
	id_str := request.CourseID
	fmt.Println(id_str)
	id, err1 := strconv.ParseInt(id_str, 10, 64)
	if err1 != nil {
		response.Code = entity.ParamInvalid
		return
	}
	course, err := GetCourseById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Code = entity.CourseNotExisted
		return
	}
	if err != nil {
		response.Code = entity.UnknownError
		fmt.Println(err)
		return
	}
	response.Code = entity.OK
	response.Data = entity.TCourse{
		CourseID:  strconv.FormatInt(course.CourseID, 10),
		Name:      course.CourseName,
		TeacherID: course.TeacherID,
	}
	return
}

func BindCourseTeacher(request entity.BindCourseRequest) (response entity.BindCourseResponse) {
	id_str := request.CourseID
	id, _ := strconv.ParseInt(id_str, 10, 64)
	course, err := GetCourseById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Code = entity.CourseNotExisted
		return
	}
	if err != nil {
		response.Code = entity.UnknownError
		fmt.Println(err)
		return
	}
	if course.TeacherID != "" {
		response.Code = entity.CourseHasBound
		return
	}
	course.TeacherID = request.TeacherID
	err1 := configs.DB.Model(&course).Update("TeacherID", course.TeacherID)
	fmt.Println(err1)
	//if err1 != nil {
	//	response.Code = entity.UnknownError
	//	fmt.Println(err1)
	//	return
	//}
	response.Code = entity.OK
	return
}
func UnbindCourseTeacher(request entity.UnbindCourseRequest) (response entity.UnbindCourseResponse) {
	id_str := request.CourseID
	id, _ := strconv.ParseInt(id_str, 10, 64)
	course, err := GetCourseById(id)
	fmt.Println(course)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Code = entity.CourseNotExisted
		return
	}
	if err != nil {
		response.Code = entity.UnknownError
		return
	}
	if course.TeacherID == "" {
		response.Code = entity.CourseNotBind
		return
	}
	course.TeacherID = ""
	configs.DB.Model(&course).Update("TeacherID", "")
	//if err1 != nil {
	//	response.Code = entity.UnknownError
	//	return
	//}
	response.Code = entity.OK
	return
}

func GetCourseListByTeacherId(id string) (courselist []models.Course, err error) {
	err = configs.DB.Where(map[string]interface{}{"TeacherID": id}).Find(&courselist).Error
	return
}

func GetTeacherCourse(request entity.GetTeacherCourseRequest) (response entity.GetTeacherCourseResponse) {
	id := request.TeacherID
	courselist, _ := GetCourseListByTeacherId(id)
	courses := make([]*entity.TCourse, 0)
	for _, course := range courselist {
		courses = append(courses, &entity.TCourse{
			CourseID:  strconv.FormatInt(course.CourseID, 10),
			Name:      course.CourseName,
			TeacherID: course.TeacherID,
		})
	}
	response.Code = entity.OK
	response.Data = struct{ CourseList []*entity.TCourse }{CourseList: courses}
	return
}
