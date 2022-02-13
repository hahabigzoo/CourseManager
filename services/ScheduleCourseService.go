package services

import (
	"Course/entity"
)

var pb, pa []int
var vis []int
var dfn int

func ScheduleCourseService(request *entity.ScheduleCourseRequest) (response entity.ScheduleCourseResponse) {
	Data := match(request.TeacherCourseRelationShip)
	if Data != nil {
		response = entity.ScheduleCourseResponse{Code: entity.OK, Data: Data}
	} else {
		response = entity.ScheduleCourseResponse{Code: entity.UnknownError, Data: Data}
	}
	return
}

func match(relationship map[string][]string) map[string]string {
	//增广路算法进行匹配，由于图是老师到课程的，所以从每一个课程开始寻找增广路，扩展老师与课程的匹配关系
	course := make(map[string]string)
	cnt := 0
	for v, _ := range relationship {
		if dfs(v, relationship, make(map[string]bool, 0), course) {
			cnt++
		}
	}
	if cnt < len(relationship) {
		//不是所有老师都匹配成功
		return nil
	}
	schedule := make(map[string]string, 0)
	for k, v := range course {
		schedule[v] = k
	}
	return schedule
}

func dfs(v string, graph map[string][]string, vis map[string]bool, course map[string]string) bool {
	vis[v] = true
	for _, u := range graph[v] {
		if _, ok := vis[u]; !ok {
			vis[u] = true
			if _, ok := course[u]; (!ok) || dfs(course[u], graph, vis, course) {
				course[u] = v
				return true
			}
		}
	}
	return false
}
