package api_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	api "github.com/automacoin/v1-database/api"
	"github.com/automacoin/v1-database/common/testers"
)

func TestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "api Suite")
}

var _ = Describe("api library", func() {

	var tester *testers.ApiTester

	BeforeEach(func() {
		router := api.Setup()
		tester = testers.NewApiTester(router)
	})

	AfterEach(func() {
		tester.Close()
	})

	Describe("Middlewares", func() {
		// TODO
		// Recovery
		// Logger
		// Authentication
	})

	Describe("Routes", func() {
		// TODO
		// Test the 404

		Describe("GET /ping", func() {
			Context("from anywhere", func() {
				It("should respond 'pong'", func() {
					response, err := tester.Get("/ping")
					Expect(err).ShouldNot(HaveOccurred())
					Expect(response.BodyStr).To(Equal(`{"message":"pong"}`))
				})

				It("should log the request", func() {

				})
			})
		})
	})
})
