package handlers

import (
	"github.com/Verifier/emailVerifier/utils"
	"github.com/Verifier/emailVerifier/verifier"
	"github.com/gin-gonic/gin"
	"net/http"
)

//VerifyEmails takes a list of emails and returns in email is verified
func VerifyEmails(c *gin.Context) {
	emailToVerify := c.Query("email")
	if emailToVerify == ""{
		utils.LogInfo("No email found in params")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "No email found in params",
		})
		return
	}


	err := verifier.ValidateEmail(emailToVerify)
	if err != nil {
		utils.LogErr("Error verifying email", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{})

}

//GetIndexPage returns the index page for the website
func GetIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
