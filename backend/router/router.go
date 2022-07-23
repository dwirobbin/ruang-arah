package router

import (
	"fmt"
	"net/http"
	"ruang-arah/backend/config"
	"ruang-arah/backend/exception"
	"ruang-arah/backend/middleware"
	"ruang-arah/backend/pkg/controller"
)

func NewRouter(ctrl *controller.Controller) *http.ServeMux {
	router := http.NewServeMux()

	fmt.Println("\t============================")
	fmt.Println("\t[INFO] Running in Port: " + config.API_PORT)
	fmt.Println("\t============================")

	router.Handle("/api/auth/login",
		exception.POST(
			http.HandlerFunc(ctrl.Login),
		),
	)

	router.Handle("/api/auth/register",
		exception.POST(
			http.HandlerFunc(ctrl.Register),
		),
	)

	router.Handle("/api/admin/questions/create",
		exception.POST(
			middleware.AuthMiddleWare(
				middleware.AdminMiddleWare(
					http.HandlerFunc(ctrl.CreateQuestion),
				),
			),
		),
	)

	router.Handle("/api/admin/questions",
		exception.GET(
			middleware.AuthMiddleWare(
				middleware.AdminMiddleWare(
					http.HandlerFunc(ctrl.GetQuestions),
				),
			),
		),
	)

	router.Handle("/api/admin/questions/detail",
		exception.GET(
			middleware.AuthMiddleWare(
				middleware.AdminMiddleWare(
					http.HandlerFunc(ctrl.GetQuestionById),
				),
			),
		),
	)

	router.Handle("/api/admin/questions/update",
		exception.PUT(
			middleware.AuthMiddleWare(
				middleware.AdminMiddleWare(
					http.HandlerFunc(ctrl.UpdateQuestion),
				),
			),
		),
	)

	router.Handle("/api/admin/questions/delete",
		exception.DELETE(
			middleware.AuthMiddleWare(
				middleware.AdminMiddleWare(
					http.HandlerFunc(ctrl.DeleteQuestion),
				),
			),
		),
	)

	router.Handle("/api/home/languages",
		exception.GET(
			middleware.AuthMiddleWare(
				http.HandlerFunc(ctrl.GetProgrammingLanguages),
			),
		),
	)

	router.Handle("/api/home/questions",
		exception.GET(
			middleware.AuthMiddleWare(
				http.HandlerFunc(ctrl.GetQuestionByProgrammingLanguageIdWithPagination),
			),
		),
	)

	router.Handle("/api/home/process-and-result",
		exception.POST(
			middleware.AuthMiddleWare(
				http.HandlerFunc(ctrl.SubmitAnswersAttempts),
			),
		),
	)

	router.Handle("/api/home/recommendation",
		exception.GET(
			middleware.AuthMiddleWare(
				http.HandlerFunc(ctrl.GetRecommendationByLevelId),
			),
		),
	)

	return router
}
