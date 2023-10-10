package v1Users

import (
	"capital-challenge-server/dbHelper"
	"capital-challenge-server/models"
	"capital-challenge-server/utils"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/u2takey/go-utils/uuid"
)
const (
	NEW_USER_STARTING_BALANCE = 1000
	NEW_USER_GAME_NUMBER = 1

)
// Register user godoc
// @Summary      Register user
// @Description  Register user
// @Tags         users
// @Param        request   body      UserRegistrationRequest  true  "request"
// @Success      200 {object} UserRegistrationResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /users/register [post]
func Registration(c *gin.Context) {
	var registrationDetails UserRegistrationRequest
	if err := c.BindJSON(&registrationDetails); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		utils.Log(c, http.StatusBadRequest)
		log.Println(err)
		return
	}

	err := validateUserRegistrationRequest(registrationDetails)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		utils.Log(c, http.StatusBadRequest)
		log.Println(err)
		return
	}

	emailExists, usernameExists, err := dbHelper.CheckIfUsernameOrEmailExists(registrationDetails.Email, registrationDetails.Username)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		utils.Log(c, http.StatusInternalServerError)
		return
	}
	if usernameExists > 0 && emailExists > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "username and email already exists"})
		utils.Log(c, http.StatusBadRequest)
		return
	} else if emailExists > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email already exists"})
		utils.Log(c, http.StatusBadRequest)
		return
	} else if usernameExists > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "username already exists"})
		utils.Log(c, http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(registrationDetails.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		utils.Log(c, http.StatusInternalServerError)
		return
	}

	var user models.Users
	user.ID = uuid.NewUUID()
	user.CurrentGameNumber = 1
	user.Email = registrationDetails.Email
	user.Username = registrationDetails.Username
	user.Password = hashedPassword

	err = dbHelper.CreateUserRecord(user)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		utils.Log(c, http.StatusInternalServerError)
		return
	}

	var userBalance models.UserBalance
	userBalance.ID = uuid.NewUUID()
	userBalance.StartingBalance = NEW_USER_STARTING_BALANCE
	userBalance.GameNumber = NEW_USER_GAME_NUMBER
	userBalance.CurrentBalance = NEW_USER_STARTING_BALANCE
	userBalance.UserID = user.ID

	err = dbHelper.CreateUserBalance(userBalance)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		utils.Log(c, http.StatusInternalServerError)
		return
	}

	var response UserRegistrationResponse
	response.CurrentBalance = userBalance.CurrentBalance
	response.CurrentGameNumber = userBalance.GameNumber
	response.Email = user.Email
	response.StartingBalance = userBalance.StartingBalance
	response.Username = user.Username

	c.JSON(http.StatusOK, response)
}

// Login user godoc
// @Summary      Login user
// @Description  Login user
// @Tags         users
// @Param        request   body      UserLoginRequest  true  "request"
// @Success      200 {object} string
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /users/ [post]
func Login(c *gin.Context) {
	var LoginDetails UserLoginRequest
	if err := c.BindJSON(&LoginDetails); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		utils.Log(c, http.StatusBadRequest)
		return
	}
	err := validateUserLoginRequest(LoginDetails)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		utils.Log(c, http.StatusBadRequest)
		return
	}
	user, err := dbHelper.GetUserByEmail(LoginDetails.Email)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		utils.Log(c, http.StatusBadRequest)
		return
	}

	err = utils.VerifyPassword(LoginDetails.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid password"})
		utils.Log(c, http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		utils.Log(c, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, token)
}

