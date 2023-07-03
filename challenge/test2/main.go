package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func getMostFrequentWord(c *gin.Context) {
	// Get the prefix from request
	prefix := c.Param("prefix")

	// Check if the prefix contains only letters using a regular expression
	match, _ := regexp.MatchString("^[a-zA-Z]+$", prefix)
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid prefix"})
		return
	}

	// Open JSON file
	file, err := os.OpenFile("words.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer file.Close()

	// Decode JSON content into a variable
	var words []string
	err = json.NewDecoder(file).Decode(&words)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode file"})
		return
	}

	// Count for every word containing the prefix in a map
	counts := make(map[string]int)
	for _, word := range words {
		if strings.HasPrefix(word, prefix) {
			counts[word]++
		}
	}

	// Check which word is the most frequent using counting map
	maxCount := 0
	mostFrequentWord := "null"
	for word, count := range counts {
		if count > maxCount {
			maxCount = count
			mostFrequentWord = word
		}
	}
	//Return the Status : OK and add the most frequent word and the number of occurence
	c.JSON(http.StatusOK, gin.H{"word": mostFrequentWord, "count": maxCount})
}

func createAWord(c *gin.Context) {

	// Get the word from request
	word := c.Param("word")

	// Open JSON file
	file, err := os.OpenFile("words.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer file.Close()

	// Decode JSON content into a variable
	var words []string
	err = json.NewDecoder(file).Decode(&words)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode file"})
		return
	}

	// Add new word
	words = append(words, word)

	// Write JSON content in words.json
	file.Seek(0, 0)
	err = json.NewEncoder(file).Encode(words)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode file"})
		return
	}

	// Return the Status : OK
	c.JSON(http.StatusOK, gin.H{"message": "Word added successfully"})
}

func main() {
	router := gin.Default()
	// Route POST : /service/word=:word
	router.POST("/service/word=:word", createAWord)
	// Route GET : /service/prefix=:prefix
	router.GET("/service/prefix=:prefix", getMostFrequentWord)
	router.Run("localhost:8080")
}
