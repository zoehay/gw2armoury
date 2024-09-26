package usersessiontest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Session", func() {
	Context("when the api key is not in the database", func() {
		When("the user provides a valid api key", func() {
			It("should attach the sessionid cookie to the response", func() {

				userJson := `{"AccountName":"Name forAccount", "APIKey":"stringthatisapikey", "Password":"stringthatispassword"}`
				req, _ := http.NewRequest("POST", "/addkey", strings.NewReader(userJson))

				gin.SetMode(gin.TestMode)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)

				c.Request = req
				AccountHandler.CreateGuest(c)

				cookie := w.Result().Cookies()
				fmt.Println(cookie)

				Expect(cookie[0].Name).To(Equal("sessionID"))

				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).NotTo(HaveOccurred())

				Expect(response["SessionID"].(string)).To(Equal(cookie[0].Value))
				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})

		// When("the user provides an invalid api key", func() {
		// 	It("should reject", func() {
		// 		// Test failed
		// 	})
		// })
	})

	// Context("when the api key is already in the database", func() {
	// 	When("the user provides a valid api key", func() {
	// 		It("should attach a cookie to the response, and the session should be renewed", func() {
	// 			// Test successful login
	// 		})
	// 	})
	// })

})
