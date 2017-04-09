package monorail

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMonorail(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoRackHD.Monorail Suite")
}

var _ = Describe("Monorail Client", func() {
	var client *Monorail
	// configure the host. include port with environment variable. For instance the vagrant image would be localhost:9090
	var rackhdEndpoint = "http://localhost:9090"
	if os.Getenv("GORACKHD_ENDPOINT") != "" {
		rackhdEndpoint = os.Getenv("GORACKHD_ENDPOINT")
	}

	Context("With a client", func() {
		BeforeEach(func() {
			client = New(rackhdEndpoint)
		})

		It("should instantiate a client", func() {
			Ω(client.Endpoint.String()).Should(Equal(rackhdEndpoint))
		})

		Context("With valid user", func() {
			username := "admin"
			password := "admin123"

			It("should log into RackHD", func() {
				auth, err := client.Login(username, password)
				Ω(err).Should(Succeed())
				Ω(auth).ShouldNot(BeNil())
			})
		})

	})
})
