package handlers

import (
	"github.com/Verifier/emailVerifier/entity"
	"github.com/Verifier/emailVerifier/utils"
	"github.com/Verifier/emailVerifier/verifier"
	"github.com/gin-gonic/gin"
)

//VerifyEmails takes a list of emails and returns in email is verified
func VerifyEmails(c *gin.Context) {
	var emailListObj entity.EmailList
	bindErr := c.Bind(&emailListObj)

	if bindErr != nil {
		utils.LogErr("Error while binding email lists", bindErr)
	}

	err := verifier.VerifyEmailLists(emailListObj)
	if err != nil {
		utils.LogErr("Error verifying list", err)
	}

}
