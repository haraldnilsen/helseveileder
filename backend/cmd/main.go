package main

import (
	"fmt"
	"net/http"
	"strconv"

	"helseveileder/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type FormData struct {
    Age                string `json:"age"`
    Education          string `json:"education"`
    HealthcarePersonnel bool  `json:"healthcare_personnel"`
    Gender             string `json:"gender"`
}

func main() {
    router := gin.Default()
    router.Use(cors.Default())
    
    // Info about user
    router.POST("/submitform", func(c *gin.Context) {

        var requestBody FormData

        if err := c.BindJSON(&requestBody); err != nil {
            fmt.Print(err)
        }

        // Capture both the ID and error returned from InsertData
		respondentId, err := db.InsertUserData(requestBody.Age, requestBody.Education, requestBody.HealthcarePersonnel, requestBody.Gender)
		
        if err != nil {
			fmt.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to insert data"})
			return
		}

        // Respond with the ID of the newly inserted respondent
		c.JSON(http.StatusOK, gin.H{"respondentID": respondentId})
    })

    // Get questions & answers from database
    router.GET("/userquestions", func(c *gin.Context) {
        respondentID, err := strconv.Atoi(c.Query("respondentID"))

        if err != nil {
            fmt.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong respondentID-format (should be int)."})
			return
        }

        questions, err := db.GetUserQuestions(respondentID)

        if err != nil {
            fmt.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting questions for given user."})
			return
        }

        c.JSON(http.StatusOK, gin.H{"questions": questions})
    })

    // Run the server on port 8080
    router.Run(":8080")
    //db.SetupDb()
}