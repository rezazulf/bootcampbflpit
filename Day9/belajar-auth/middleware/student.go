package middleware

import (
	"encoding/json"
	"net/http"
)

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGET(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}
	return nil
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}

func init() {
	students = append(students, &Student{Id: "s001", Name: "Bourne", Grade: 2})
	students = append(students, &Student{Id: "s002", Name: "Ethan", Grade: 2})
	students = append(students, &Student{Id: "s003", Name: "Wick", Grade: 3})
}
