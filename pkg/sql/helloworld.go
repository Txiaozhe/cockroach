package sql

import (
	"context"
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

func (p *planner) Helloworld(ctx context.Context, stmt *tree.Helloworld) (planNode, error) {
	switch stmt.Mode {
	case tree.HelloworldModeSay:
		fmt.Println("hello!!")
	case tree.HelloworldModeSmile:
		fmt.Println("haha!!")
	default:
		return nil, fmt.Errorf("Unhandled Helloworld mode %v!", stmt.Mode)
	}

	return &zeroNode{}, nil
}
