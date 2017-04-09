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
	var client Monorail
	// configure the host. include port with environment variable. For instance the vagrant image would be localhost:9090
	var rackhdEndpoint = "https://localhost:9093"
	if os.Getenv("GORACKHD_ENDPOINT") != "" {
		rackhdEndpoint = os.Getenv("GORACKHD_ENDPOINT")
	}

	Context("With a client", func() {
		BeforeEach(func() {
			client = New(rackhdEndpoint)
		})

		It("should instantiate a client", func() {
			client.Endpoint.Should(Equal(rackhdEndpoint))
		})

	})
})
