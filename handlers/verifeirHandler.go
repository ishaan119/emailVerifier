package handlers

import (
	"github.com/Verifier/emailVerifier/entity"
	"github.com/Verifier/emailVerifier/utils"
	"github.com/Verifier/emailVerifier/verifier"
	"github.com/gin-gonic/gin"
	"net/http"
)

//VerifyEmails takes a list of emails and returns in email is verified
func VerifyEmails(c *gin.Context) {

	emailToVerify := c.PostForm("email")

	emailList := make([]string, 0)
	emailList = append(emailList, emailToVerify)
	var emailListObj entity.EmailList
	emailListObj.Emails = emailList
	err := verifier.VerifyEmailLists(emailListObj)
	if err != nil {
		utils.LogErr("Error verifying list", err)
	}

}

//GetIndexPage returns the index page for the website
func GetIndexPage(c *gin.Context)  {

	c.HTML(http.StatusOK, "index.html", nil)
}
