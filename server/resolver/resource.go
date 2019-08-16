package resolver

import (
	"context"
	"github.com/QN-Resources/server/entity"
)

func (r *ResourceQuery) List(ctx context.Context, args struct {
	Id string
}) []*ResourceResolver {
	list := []*ResourceResolver{}
	root := entity.Instance()
	for _, item := range root.FindNodesById("root") {
		list = append(list, &ResourceResolver{node: &item})
	}
	return list
}

func (_ *ResourceQuery) Hello() string { return "Hello, world!" }

type ResourceResolver struct {
	node *entity.TreeNode
}

func (r *ResourceResolver) Id() string {
	return r.node.Id
}
