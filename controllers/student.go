package controllers

import (
	"eduSense/helpers"
	"eduSense/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type StudentController struct {}


func (StudentController) ReadMany(w http.ResponseWriter, r *http.Request)  {
	db:= helpers.GetDBConnection()
	defer db.Close()

	students := []*models.Student{}

	db.Find(&students)

	studentListBytes, err := json.Marshal(students)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(studentListBytes)

}

func (StudentController) Create(w http.ResponseWriter, r *http.Request)  {
	db:= helpers.GetDBConnection()
	defer db.Close()

	student := models.Student{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	student.Name = r.Form.Get("name")
	student.Email = r.Form.Get("email")
	student.Address = r.Form.Get("address")

	db.FirstOrCreate(&student, models.Student{Email: student.Email})

	studentBytes, err := json.Marshal(student)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(studentBytes)

}



