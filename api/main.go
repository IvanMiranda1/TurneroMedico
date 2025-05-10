package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error  loading .env file")
	}

	ginEngine := gin.Default()
	ginEngine.POST("/medico", func(c *gin.Context) {
		var medico domain.Medico

		if err := c.BindJSON(&medico); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		//setear datos automaticos como fecha de creacion, etc

		//Set a timeout to allow the connection process to abort if it takes too long
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		//client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

		conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatal(err)
		}

		err = conn.Ping(ctx)
		if err != nil {
			conn.Close(ctx)
			return nil, err
		}

	})

	log.Fatalln(ginEngine.Run(":8001"))
}
