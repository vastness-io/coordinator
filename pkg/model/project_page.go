package model

type ProjectPage struct {
	Meta struct {
		CurrentPage int
		LastPage    int
		PerPage     int
		TotalCount  int
	}
	Projects []*Project
}
