package server

import (
	"battery-analysis-platform/app/main/consts"
	"battery-analysis-platform/app/main/controller/api"
	"battery-analysis-platform/app/main/controller/auth"
	"battery-analysis-platform/app/main/controller/file"
	"battery-analysis-platform/app/main/controller/websocket"
	"battery-analysis-platform/app/main/middleware"
	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {
	r.GET("/login", auth.Login)
	r.POST("/login", auth.Login)

	rootPath := r.Group("/")
	rootPath.Use(middleware.PermissionRequired(consts.UserTypeCommonUser))
	{
		rootPath.POST("/logout", auth.Logout)
	}

	apiV1Path := r.Group("/api/v1")
	apiV1Path.Use(middleware.PermissionRequired(consts.UserTypeCommonUser))
	{
		apiV1Path.POST("/self/change-password", api.UpdateSelfPassword)

		apiV1Path.GET("/sys-info", api.GetSysInfo)

		apiV1Path.GET("/mining/base", api.GetBatteryList)

		apiV1Path.POST("/mining/tasks", api.CreateMiningTask)
		apiV1Path.GET("/mining/tasks", api.GetMiningTaskList)
		apiV1Path.GET("/mining/tasks/:taskId/data", api.GetMiningTaskData)
		apiV1Path.DELETE("/mining/tasks/:taskId", api.DeleteMiningTask)

		apiV1Path.POST("/dl/tasks", api.CreateDlTask)
		apiV1Path.GET("/dl/tasks", api.GetDlTaskList)
		apiV1Path.GET("/dl/tasks/:taskId/training-history", api.GetDlTaskTraningHistory)
		apiV1Path.GET("/dl/tasks/:taskId/eval-result", api.GetDlEvalResultHistory)
		apiV1Path.DELETE("/dl/tasks/:taskId", api.DeleteDlTask)
	}

	apiV1NeedPermissionPath := r.Group("/api/v1")
	apiV1NeedPermissionPath.Use(middleware.PermissionRequired(consts.UserTypeSuperUser))
	{
		apiV1NeedPermissionPath.POST("/users", api.CreateUser)
		apiV1NeedPermissionPath.GET("/users", api.GetUserList)
		apiV1NeedPermissionPath.PUT("/users/:name", api.UpdateUserInfo)
	}

	wsV1Path := r.Group("/websocket/v1")
	wsV1Path.Use(middleware.PermissionRequired(consts.UserTypeCommonUser))
	{
		//wsV1Path.GET("/sys-info", websocket.GetSysInfo)
		wsV1Path.GET("/mining/tasks", websocket.GetMiningTaskList)
		wsV1Path.GET("/dl/tasks", websocket.GetDlTaskList)
		wsV1Path.GET("/dl/tasks/:taskId/training-history", websocket.GetDlTaskTraningHistory)
	}

	filePath := r.Group("/file")
	filePath.Use(middleware.PermissionRequired(consts.UserTypeCommonUser))
	{
		filePath.GET("/dl/model/:taskId", file.DownloadDlModel)
	}
}
