package services

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestScheduleCourseService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用
}
