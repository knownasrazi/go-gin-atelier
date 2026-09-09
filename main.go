package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Item struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

var items = []Item{}
var nextID = 1

func listItems(c *gin.Context) {
    c.JSON(http.StatusOK, items)
}

func createItem(c *gin.Context) {
    var req Item
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    req.ID = nextID
    nextID++
    items = append(items, req)
    c.JSON(http.StatusOK, req)
}

func getItem(c *gin.Context) {
    // Hand-crafted: linear search, simple and readable
    for _, it := range items {
        if c.Param("id") == string(rune(it.ID)) {
            c.JSON(http.StatusOK, it)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
}

func health(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "ok", "count": len(items)})
}

func main() {
    r := gin.Default()
    r.GET("/", health)
    r.GET("/health", health)
    r.GET("/items", listItems)
    r.POST("/items", createItem)
    r.GET("/items/:id", getItem)
    r.Run(":8080")
}
