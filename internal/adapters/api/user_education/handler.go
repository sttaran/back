package user_education

import (
	"back/internal/adapters/middlewares"
	"back/internal/domain/education_institution"
	"back/internal/domain/user_education"
	"back/pkg/logger"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
)

const (
	createPath = "/userEducation"
)

type Handler struct {
	userEducationService userEducationService
	logger               logger.ILogger
}

func NewUserEducationHandler(storage *mysqlClient.MySQLClient, logger logger.ILogger) *Handler {
	eduStorage := education_institution.NewEducationInstitutionStorage(storage)
	userEduStorage := user_education.NewUserEducationStorage(storage)
	userEducationService := user_education.NewUserEducationService(userEduStorage, eduStorage)
	return &Handler{
		userEducationService: userEducationService,
		logger:               logger,
	}
}

func (h *Handler) Register(e *gin.Engine) {
	authorized := e.Group("/")
	authorized.Use(middlewares.AuthMiddleware)
	authorized.POST(createPath, h.createUserEducationExperience())
}
