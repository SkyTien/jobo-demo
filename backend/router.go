package main

import (
	"Goal-Back-End/controller"
)

func initializeRoutes() {
	// Serve static files from the "static" directory
	router.Static("/web", "./web")

	userRoutes := router.Group("/api/usr")
	{
		userRoutes.POST("/login", controller.ReqLogin)
	}

	patientRoute := router.Group("/api/patient")
	{
		patientRoute.GET("/lists", controller.GetList)
		patientRoute.GET("/order/:id", controller.GetOrderById)
		patientRoute.PUT("/order/:id", controller.UpdateOrderById)
	}
}
