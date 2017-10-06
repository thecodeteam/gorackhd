package rackhd

import (
	"os"
	"testing"

	api "github.com/go-openapi/runtime"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spiegela/gorackhd/client/nodes"
	"github.com/spiegela/gorackhd/models"
)

func TestRackHD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoRackHD.RackHD Suite")
}

var _ = Describe("RackHD Client", func() {
	var client *RackHD
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
		if os.Getenv("GORACKHD_USERNAME") != "" {
			rackhdEndpoint = os.Getenv("GORACKHD_USERNAME")
		}
		password := "admin123"
		if os.Getenv("GORACKHD_PASSWORD") != "" {
			rackhdEndpoint = os.Getenv("GORACKHD_PASSWORD")
		}

		It("should log into RackHD", func() {
			auth, err := client.Login(username, password)
			Ω(err).Should(Succeed())
			Ω(auth).ShouldNot(BeNil())
		})
	})

})

var _ = Describe("RackHD Nodes Client", func() {
	var client *RackHD
	var auth api.ClientAuthInfoWriter
	var nodeID string

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
				&models.PartialNode{
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

	Context("with an existing node", func() {

		BeforeEach(func() {
			postParams := nodes.NewNodesPostParams().
				WithIdentifiers(
					&models.PartialNode{
						AutoDiscover: "false",
						Identifiers:  []string{"08:00:27:A4:52:95"},
						Name:         "vbox1",
						Type:         "compute",
						Obms:         []*models.NodesPostObmByID{},
						Relations:    []*models.RelationsObj{},
						Tags:         []string{},
					},
				)
			client.Nodes().NodesPost(postParams, auth)
			getAllResp, _ := client.Nodes().NodesGetAll(nil, auth)
			nodeID = getAllResp.Payload[0].ID
		})

		AfterEach(func() {
			delParams := nodes.NewNodesDelByIDParams().
				WithIdentifier(nodeID)
			client.Nodes().NodesDelByID(delParams, auth)
		})

		It("should handle node tag CRUD operations", func() {
			getParams := nodes.NewNodesGetTagsByIDParams().
				WithIdentifier(nodeID)
			getAllTagsResp, err := client.Nodes().NodesGetTagsByID(getParams, auth)
			Ω(err).Should(BeNil())
			Ω(getAllTagsResp.Payload).Should(BeEmpty())

			patchParams := nodes.NewNodesPatchTagByIDParams().
				WithIdentifier(nodeID).
				WithBody(&models.PostTags{Name: "testTag"})
			patchTagResp, err := client.Nodes().NodesPatchTagByID(patchParams, auth)
			Ω(err).Should(BeNil())
			Ω(patchTagResp.Payload).ShouldNot(BeNil())

			getAllTagsResp, err = client.Nodes().NodesGetTagsByID(getParams, auth)
			Ω(err).Should(BeNil())
			Ω(getAllTagsResp.Payload).Should(HaveLen(1))
			Ω(getAllTagsResp.Payload[0]).Should(Equal("testTag"))

			delParams := nodes.NewNodesDelTagByIDParams().
				WithIdentifier(nodeID).
				WithTagName("testTag")
			delTagResp, err := client.Nodes().NodesDelTagByID(delParams, auth)
			Ω(err).Should(BeNil())
			Ω(delTagResp).ShouldNot(BeNil())

			getAllTagsResp, err = client.Nodes().NodesGetTagsByID(getParams, auth)
			Ω(err).Should(BeNil())
			Ω(getAllTagsResp.Payload).Should(BeEmpty())
		})

	})

})
