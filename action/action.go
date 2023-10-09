package action

import (
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/goctl-yaml/generate"
)

func Generator(ctx *cli.Context) error {
	filename := ctx.String("filename")
	if len(filename) == 0 {
		filename = "auth_gateway.yaml"
	}

	p, err := plugin.NewPlugin()
	if err != nil {
		return err
	}

	return generate.Do(filename, p)
}
