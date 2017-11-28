package stringsvc

import (
	"context"

	"github.com/devimteam/microgen/example/svc/entity"
)

// @microgen middleware, logging, grpc, http, recover
// @grpc-addr devim.string.team
// @protobuf github.com/devimteam/protobuf/stringsvc
type StringService interface {
	// @logs-ignore ans, err
	Uppercase(ctx context.Context, str ...map[string]interface{}) (ans string, err error)
	Count(ctx context.Context, text string, symbol string) (count int, positions []int, err error)
	// @logs-len comments
	TestCase(ctx context.Context, comments []*entity.Comment) (tree map[string]int, err error)
}
