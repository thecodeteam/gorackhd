package monorail

import (
	"os"
	"testing"

	"github.com/go-openapi/runtime"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spiegela/gorackhd/client/nodes"
	"github.com/spiegela/gorackhd/models"
)

func TestMonorail(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoRackHD.Monorail Suite")
}

var _ = Describe("Monorail Client", func() {
	var client *Monorail
	// configure the host. include port with environment variable.
	// For instance the vagrant image would be localhost:9090
	var rackhdEndpoint = "http://localhost:9090"
	if os.Getenv("GORACKHD_ENDPOINT") != "" {
		rackhdEndpoint = os.Getenv("GORACKHD_ENDPOINT")
	}

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

var _ = Describe("Monorail Nodes Client", func() {
	var client *Monorail
	var auth runtime.ClientAuthInfoWriter

	BeforeEach(func() {
		username := "admin"
		password := "admin123"
		var rackhdEndpoint = "http://localhost:9090"
		if os.Getenv("GORACKHD_ENDPOINT") != "" {
			rackhdEndpoint = os.Getenv("GORACKHD_ENDPOINT")
		}
		client = New(rackhdEndpoint)
		auth, _ = client.Login(username, password)
	})

	It("should handle node CRUD operations", func() {
		getAllResp, err := client.Nodes().NodesGetAll(nil, auth)
		Ω(err).Should(BeNil())
		Ω(getAllResp.Payload).Should(BeEmpty())

		postParams := nodes.NewNodesPostParams().
			WithIdentifiers(
				&models.Node20PartialNode{
					AutoDiscover: "false",
					Identifiers:  []string{"08:00:27:A4:52:95"},
					Name:         "vbox1",
					Type:         "compute",
					Obms:         []*models.NodesPostObmByID{},
					Relations:    []*models.RelationsObj{},
					Tags:         []string{},
				},
			)
		nodesCreated, err := client.Nodes().NodesPost(postParams, auth)
		Ω(err).Should(BeNil())
		Ω(nodesCreated).ShouldNot(BeNil())

		getAllResp, err = client.Nodes().NodesGetAll(nil, auth)
		Ω(err).Should(BeNil())
		Ω(getAllResp.Payload).Should(HaveLen(1))

		delParams := nodes.NewNodesDelByIDParams().
			WithIdentifier(getAllResp.Payload[0].ID)
		nodesDeleted, err := client.Nodes().NodesDelByID(delParams, auth)
		Ω(err).Should(BeNil())
		Ω(nodesDeleted).ShouldNot(BeNil())

		getAllResp, err = client.Nodes().NodesGetAll(nil, auth)
		Ω(err).Should(BeNil())
		Ω(getAllResp.Payload).Should(BeEmpty())
	})
})
