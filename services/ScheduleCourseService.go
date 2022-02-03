package services

import (
	"Course/serializer"
)

var pb, pa []int
var vis []int
var dfn int

func ScheduleCourseService(request serializer.ScheduleCourseRequest) (response serializer.ScheduleCourseResponse) {
	Data := match(request.TeacherCourseRelationShip)
	if Data != nil {
		response = serializer.ScheduleCourseResponse{Code: serializer.OK, Data: Data}
	} else {
		response = serializer.ScheduleCourseResponse{Code: serializer.UnknownError, Data: Data}
	}
	return
}

func match(relationship map[string][]string) map[string]string {
	course := make(map[string]string)
	cnt := 0
	for v, _ := range relationship {
		if dfs(v, relationship, make(map[string]bool, 0), course) {
			cnt++
		}
	}
	if cnt < len(relationship) {
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
