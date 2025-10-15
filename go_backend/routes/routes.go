package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type test struct {
	Message  string `json:"test_message"`
	Returned bool   `json:"returned"`
}

var testStruct = test{
	Message:  "this is the test message",
	Returned: true,
}

func Index(c *gin.Context, uploadDir, path string) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "index.html", nil)
		return
	}

	file, err := c.FormFile("video") // "video" matches name="video" from html template
	if err != nil {
		fmt.Printf("Error getting file: %s", err)
		return
	}
	log.Println(file.Filename)

	c.SaveUploadedFile(file, path+uploadDir+"/"+file.Filename)
	c.HTML(http.StatusAccepted, "successful_upload.html", gin.H{
		"filename": file.Filename,
	})
	// c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

// testing successful upload page
func Test(c *gin.Context) {
	c.HTML(http.StatusAccepted, "successful_upload.html", gin.H{
		"filename": "FILENAMEHERE",
	})
}
