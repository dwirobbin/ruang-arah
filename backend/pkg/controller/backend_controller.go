package controller

import (
	"encoding/json"
	"net/http"
	"ruang-arah/backend/helper"
	"ruang-arah/backend/model/web"
	"ruang-arah/backend/pkg/service"
	"strconv"
)

type BackendControllerImpl struct {
	backendService service.BackendService
}

func NewBackendController(backendService service.BackendService) *BackendControllerImpl {
	return &BackendControllerImpl{
		backendService: backendService,
	}
}

func (c *BackendControllerImpl) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var questionRequest web.QuestionRequest
	err := json.NewDecoder(r.Body).Decode(&questionRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})
		return
	}

	err = questionRequest.ValidateQuestion("create")
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: http.StatusText(http.StatusUnprocessableEntity),
			Data:    err.Error(),
		})
		return
	}

	questionResponse, err := c.backendService.CreateQuestion(questionRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "Successfully created",
		Data:    questionResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *BackendControllerImpl) GetQuestions(w http.ResponseWriter, r *http.Request) {
	quizResponse, err := c.backendService.GetQuestions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "Successfully retrieved",
		Data:    quizResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *BackendControllerImpl) GetQuestionById(w http.ResponseWriter, r *http.Request) {
	questionId, _ := strconv.Atoi(r.URL.Query().Get("questionId"))

	quizResponse, err := c.backendService.GetQuestionById(int32(questionId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "Successfully retrieved",
		Data:    quizResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *BackendControllerImpl) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	var questionRequest web.QuestionRequest
	err := json.NewDecoder(r.Body).Decode(&questionRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})
		return
	}

	err = questionRequest.ValidateQuestion("update")
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: http.StatusText(http.StatusUnprocessableEntity),
			Data:    err.Error(),
		})
		return
	}

	questionId, _ := strconv.Atoi(r.URL.Query().Get("questionId"))
	questionResponse, err := c.backendService.UpdateQuestion(
		int32(questionId), questionRequest,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "Successfully updated",
		Data:    questionResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *BackendControllerImpl) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	questionId, _ := strconv.Atoi(r.URL.Query().Get("questionId"))

	_, err := c.backendService.DeleteQuestion(int32(questionId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "Successfully deleted",
		Data:    nil,
	}

	helper.WriteToResponseBody(w, webResponse)
}
