package myhttp

import (
	objects "StudyApi/objects/subject"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	Database *objects.Database
}

func NewHTTPHandlers(db *objects.Database) *HTTPHandlers {
	return &HTTPHandlers{
		Database: db,
	}
}

func (h *HTTPHandlers) HandleAddTeacher(w http.ResponseWriter, r *http.Request) {
	var request RequestTeacherDTO

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := request.Validation(); err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := h.Database.AddTeacher(request.Name); err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *HTTPHandlers) HandleUpdateTeacher(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	var request RequestTeacherDTO

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := request.Validation(); err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := h.Database.UpdateTacher(id, request.Name); err != nil {
		errDTO := CreateErrDTO(objects.ErrorTeacherNotFound)
		http.Error(w, errDTO.ToString(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *HTTPHandlers) HandleGetAllTeachers(w http.ResponseWriter, r *http.Request) {
	teachers, err := h.Database.GetAllSubjects()

	if err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(teachers, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(b)
	if err != nil {
		fmt.Println("Failed to write HTTP response!: ", err)
		return
	}
}

func (h *HTTPHandlers) HandleAddSubject(w http.ResponseWriter, r *http.Request) {
	var request RequestSubjectDTO

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := request.Validation(); err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	err := h.Database.AddSubject(request.Name, request.Teacher)
	if err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *HTTPHandlers) HandleGetSubject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	subject, err := h.Database.GetSubject(id)

	if err != nil {
		errDTO := CreateErrDTO(objects.ErrorSubjectNotFound)
		http.Error(w, errDTO.ToString(), http.StatusNotFound)
		return
	}

	b, err := json.MarshalIndent(subject, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(b)
	if err != nil {
		fmt.Println("Failed to write HTTP response!: ", err)
		return
	}
}

func (h *HTTPHandlers) HandleGetAllSubjects(w http.ResponseWriter, r *http.Request) {
	subjects, err := h.Database.GetAllSubjects()
	if err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(subjects, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(b)
	if err != nil {
		fmt.Println("Failed to write HTTP response!: ", err)
		return
	}
}

func (h *HTTPHandlers) HandleUpdateSubject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	var request RequestSubjectDTO

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := request.Validation(); err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err = h.Database.UpdateSubject(id, request.Name, request.Teacher); err != nil {
		errDTO := CreateErrDTO(objects.ErrorSubjectNotFound)
		http.Error(w, errDTO.ToString(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HTTPHandlers) HandleDeleteSubject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}

	if err := h.Database.DeleteSubject(id); err != nil {
		errDTO := CreateErrDTO(objects.ErrorSubjectNotFound)
		http.Error(w, errDTO.ToString(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
