package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "api/models"
    "api/database"

    "github.com/swaggo/swag/example/celler/httputil"
)



// @Summary Get list of users
// @Description Get list of users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} gin.H
// @Router /users [get]
func GetUsers(c *gin.Context) {
    var users []models.User
    database.DB.Find(&users)

    c.JSON(http.StatusOK, gin.H{"data": users})
}

// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} gin.H
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var user models.User
    database.DB.First(&user, id)

    c.JSON(http.StatusOK, gin.H{"data": user})
}

// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param input body models.User true "User input"
// @Success 200 {object} gin.H
// @Router /users [post]
func CreateUser(c *gin.Context) {
    var userInput models.User
    if err := c.ShouldBindJSON(&userInput); err != nil {
        httputil.NewError(c, http.StatusBadRequest, err)
        return
    }

    database.DB.Create(&userInput)
    c.JSON(http.StatusOK, gin.H{"data": userInput})
}

// @Summary Update a user by ID
// @Description Update a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param input body models.User true "User input"
// @Success 200 {object} gin.H
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var userInput models.User
    if err := c.ShouldBindJSON(&userInput); err != nil {
        httputil.NewError(c, http.StatusBadRequest, err)
        return
    }

    var user models.User
    database.DB.First(&user, id)
    user.Name = userInput.Name
    user.Email = userInput.Email
    database.DB.Save(&user)

    c.JSON(http.StatusOK, gin.H{"data": user})
}

// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} gin.H
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    database.DB.Delete(&models.User{}, id)

    c.JSON(http.StatusOK, gin.H{"data": "User deleted"})
}