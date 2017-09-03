package verifier

import (
	"fmt"
	"github.com/Verifier/emailVerifier/entity"
	"github.com/Verifier/emailVerifier/utils"
)

//VerifyEmailLists .
func VerifyEmailLists(emailListObj entity.EmailList) (err error) {

	for _, email := range emailListObj.Emails {
		err = ValidateEmail(email)
		if err != nil {
			utils.LogErr("Error validating email", err)
			fmt.Println("InValid" + email)
		}

		fmt.Println("Valid" + email)
	}

	return err
}
