package controller

import (
	"encoding/json"
	"net/http"
	"ruang-arah/backend/helper"
	"ruang-arah/backend/middleware"
	"ruang-arah/backend/model/web"
	"ruang-arah/backend/pkg/service"
	"strconv"
)

type FrontendControllerImpl struct {
	frontendService service.FrontendService
}

func NewFrontendController(frontendService service.FrontendService) *FrontendControllerImpl {
	return &FrontendControllerImpl{
		frontendService: frontendService,
	}
}

func (c *FrontendControllerImpl) GetProgrammingLanguages(w http.ResponseWriter, r *http.Request) {
	proglangResponses, err := c.frontendService.GetProgrammingLanguages()
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
		Data:    proglangResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *FrontendControllerImpl) GetQuestionByProgrammingLanguageIdWithPagination(w http.ResponseWriter, r *http.Request) {
	programmingLanguageId, _ := strconv.Atoi(r.URL.Query().Get("programmingLanguageId"))

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	questionResponses, err := c.frontendService.GetQuestionByProgrammingLanguageIdWithPagination(
		int32(programmingLanguageId), int32(page), int32(limit),
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
		Message: "Successfully retrieved",
		Data:    questionResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *FrontendControllerImpl) SubmitAnswersAttempts(w http.ResponseWriter, r *http.Request) {
	userId, _ := middleware.GetUserId()
	var answerAttemptReq web.AnswerAttemptRequest

	err := json.NewDecoder(r.Body).Decode(&answerAttemptReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})
		return
	}

	err = answerAttemptReq.ValidateAnswerAttempt()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})
		return
	}

	answersAttemptResp, err := c.frontendService.SubmitAnswersAttempts(
		userId, answerAttemptReq,
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
		Message: "Success",
		Data:    answersAttemptResp,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *FrontendControllerImpl) GetRecommendationByLevelId(w http.ResponseWriter, r *http.Request) {
	levelId, _ := strconv.Atoi(r.URL.Query().Get("level_id"))

	recommendationResponse, err := c.frontendService.GetRecommendationByLevelId(
		int32(levelId),
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
		Message: "Success retrieved",
		Data:    recommendationResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
