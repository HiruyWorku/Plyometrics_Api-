package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samsmith00/plyometricAPI/go_backend/routes"
)

func main() {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get prot and directories
	port := os.Getenv("PORT")
	uploadsDir := os.Getenv("UPLOAD_DIR")
	path := os.Getenv("CURRENT_DIR")

	// create directory for videos
	if _, err := os.Stat(path + uploadsDir); os.IsNotExist(err) {
		// path/to/whatever does not exist
		if err := os.Mkdir(".uploads", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	router := gin.Default()

	router.LoadHTMLFiles("index.html", "successful_upload.html")

	fmt.Printf("path + uploadDir: %s\n", path+uploadsDir)
	router.Any("/", func(c *gin.Context) {
		routes.Index(c, uploadsDir, path)
	})

	router.Any("/test", func(c *gin.Context) {
		routes.Test(c)
	})

	router.Run(port)
}
