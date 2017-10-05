package mock

//go:generate mockgen -package mock -destination mock_rackhd.go github.com/spiegela/gorackhd/rackhd Iface
//go:generate mockgen -package mock -destination mock_nodes.go github.com/spiegela/gorackhd/rackhd NodeIface
//go:generate mockgen -package mock -destination mock_skus.go github.com/spiegela/gorackhd/rackhd SkuIface
//go:generate mockgen -package mock -destination mock_tags.go github.com/spiegela/gorackhd/rackhd TagIface
