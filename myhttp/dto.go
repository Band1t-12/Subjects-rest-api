package myhttp

import (
	"encoding/json"
	"errors"
	"time"
)

type RequestDTO struct {
	Name    string
	Teacher string
}

func (r RequestDTO) Validation() error {
	if r.Name == "" {
		return errors.New("subject name is empty!")
	}
	if r.Teacher == "" {
		return errors.New("teacher name is empty!")
	}

	return nil
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

func CreateErrDTO(err error) ErrorDTO {
	errDTO := ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}

	return errDTO
}
