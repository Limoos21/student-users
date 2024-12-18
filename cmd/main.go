package main

import (
	"github.com/gin-gonic/gin"
	"log"
	_ "os"
	"stud-trener/config"
	"stud-trener/internal/application"
	"stud-trener/internal/infra/db"
	"stud-trener/internal/infra/db/repository"
	logger2 "stud-trener/internal/infra/logger"
	"stud-trener/internal/infra/middleware"
	"stud-trener/internal/infra/router"
	"stud-trener/internal/interfaces/http"
)

func main() {
	// Загружаем конфигурацию
	logger, err := logger2.NewLogger("debug")
	logger.Info("Initializing logger")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Не удалось загрузить конфигурацию:", err)
	}
	logger.Info("Initializing config")

	dsn := "host=" + cfg.DBHost + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " port=" + cfg.DBPort + " sslmode=disable"
	dataBase, err := db.NewDatabase(dsn, logger)
	if err != nil {
		logger.Error("Connecting to database")
		panic(err)
	}
	logger.Info("Connecting to database")

	studentRepository := repository.NewStudentRepository(dataBase.DbPostgres, logger)
	teamRepository := repository.NewTeamRepository(dataBase.DbPostgres, logger)
	teamTrainerRepository := repository.NewTeamTrainerRepository(dataBase.DbPostgres, logger)
	trainerRepository := repository.NewTrainerRepository(dataBase.DbPostgres, logger)
	trainRepository := repository.NewTrainRepository(dataBase.DbPostgres, logger)
	tournamentRepository := repository.NewTournamentRepository(dataBase.DbPostgres, logger)
	reportRepository := repository.NewReportRepository(dataBase.DbPostgres, logger)
	userRepository := repository.NewUserRepository(dataBase.DbPostgres, logger)
	logger.Info("Initializing repository")

	studentUseCases := application.NewStudentUseCase(studentRepository, logger)
	teamTrainerUseCase := application.NewTeamTrainerUseCase(teamTrainerRepository, logger)
	trainerUseCase := application.NewTrainerUseCase(trainerRepository, logger)
	teamUsecases := application.NewTeamUseCase(teamRepository, logger)
	trainUseCase := application.NewTrainUseCase(trainRepository, logger)
	tournamentUseCase := application.NewTournamentUseCase(tournamentRepository, logger)
	reportUseCase := application.NewReportUseCase(reportRepository, logger)
	userUseCase := application.NewUserUseCase(userRepository, logger)
	logger.Info("Initializing useCases")

	//TODO: изменить при релизе gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	noCors := middleware.NoCorsMiddelware()
	server.Use(gin.Recovery())
	server.Use(noCors)
	server.StaticFile("/favicon.ico", "static/favicon.ico")

	studentGroup := router.NewRouterGroup("api/v1/student", []gin.HandlerFunc{middleware.NoOpMiddleware()}, server)
	teamGroup := router.NewRouterGroup("api/v1/team", []gin.HandlerFunc{middleware.NoOpMiddleware()}, server)
	teamTrainerGroup := router.NewRouterGroup("api/v1/team-trainer", []gin.HandlerFunc{middleware.NoOpMiddleware()}, server)
	tournamentGroup := router.NewRouterGroup("api/v1/tournament", []gin.HandlerFunc{middleware.NoOpMiddleware()}, server)
	trainerGroup := router.NewRouterGroup("api/v1/trainer", []gin.HandlerFunc{middleware.NoOpMiddleware()}, server)
	trainGroup := router.NewRouterGroup("api/v1/train", []gin.HandlerFunc{middleware.NoOpMiddleware()}, server)
	reportGroup := router.NewRouterGroup("api/v1/report", []gin.HandlerFunc{middleware.NoOpMiddleware()}, server)
	userGroup := router.NewRouterGroup("api/v1/user", []gin.HandlerFunc{middleware.NoOpMiddleware()}, server)
	http.NewStudentHandler(studentGroup, studentUseCases, logger)
	http.NewTeamHandler(teamGroup, teamUsecases, logger)
	http.NewTeamTrainerHandler(teamTrainerGroup, teamTrainerUseCase, logger)
	http.NewTournamentHandler(tournamentGroup, tournamentUseCase, logger)
	http.NewTrainerHandler(trainerGroup, trainerUseCase, logger)
	http.NewTrainHandler(trainGroup, trainUseCase, logger)
	http.NewReportHandler(reportGroup, reportUseCase, logger)
	http.NewUserHandler(userGroup, userUseCase, logger)
	logger.Info("Initializing handlers")
	err = server.Run(":8080")
	if err != nil {
		return
	}
	logger.Info("Starting server")
}
