package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getMostFrequentWord(c *gin.Context) {

}

func createAWord(c *gin.Context) {

	word := c.Param("word")

	// Ouvre le fichier JSON
	file, err := os.OpenFile("words.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer file.Close()

	// Décode le contenu JSON dans une variable
	var words []string
	err = json.NewDecoder(file).Decode(&words)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode file"})
		return
	}

	// Ajoute le nouveau mot à la variable
	words = append(words, word)

	// Réécrit le contenu JSON dans le fichier
	file.Seek(0, 0)
	err = json.NewEncoder(file).Encode(words)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Word added successfully"})
}

func main() {
	router := gin.Default()
	router.POST("/service/word=:word", createAWord)
	router.GET("/service/prefix=:prefix", getMostFrequentWord)
	router.Run("localhost:8080")
}
