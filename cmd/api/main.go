package main

import (
	"log"
	"todo-go-restapi/internal/config"
	"todo-go-restapi/internal/database"
	"todo-go-restapi/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// migrate create -ext sql -dir migrations -seq create_todos_api_table


func main() {
	var cfg *config.Config 
	var err error 
	cfg, err = config.Load()

	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	var pool *pgxpool.Pool 
	pool, err = database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer pool.Close()

	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context){
		 c.JSON(200, gin.H{
			 "message": "Todo API is running good!",
			 "status": "success",
			 "database": "connected",
		 })
	})

	router.POST("/todos", handlers.CreateTodoHandler(pool))
router.GET("/todos", handlers.GetAllTodosHandler(pool))
	router.GET("/todos/:id", handlers.GetTodoByIDHandler(pool))
	router.PUT("/todos/:id", handlers.UpdateTodoHandler(pool))
	
	router.Run(":"+ cfg.Port)
}