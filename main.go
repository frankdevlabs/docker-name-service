package main

import (
	"docker-name-service/internal"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// response is a struct for response.
type response struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// getNewDockerName returns a name which similar to name of the docker container.
func getNewDockerName(c *gin.Context) {
	id := uuid.New()
	n := internal.GetRandomName(3)

	r := &response{
		ID:   id,
		Name: n,
	}

	c.IndentedJSON(http.StatusOK, r)
}

func main() {
	router := gin.Default()
	router.GET("/name", getNewDockerName)

	router.Run("localhost:8080")
}
