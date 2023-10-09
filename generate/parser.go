package generate

import (
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"strconv"
	"strings"
)

func applyGenerate(p *plugin.Plugin) (*AuthenticateResp, error) {
	authUnits := make([]*AuthenticateUnit, 0)
	srv, _ := strconv.Unquote(p.Api.Info.Properties["srv"])
	for _, group := range p.Api.Service.Groups {
		for _, route := range group.Routes {
			var path string
			if len(srv) > 0 {
				path = srv + group.GetAnnotation("prefix") + route.Path
			} else {
				path = group.GetAnnotation("prefix") + route.Path
			}
			if path[0] != '/' {
				path = "/" + path
			}
			path = "/api" + path
			authUnit := &AuthenticateUnit{
				Method: strings.ToUpper(route.Method),
				URL:    path,
			}
			roleIntArray := make([]int, 0)
			if route.AtDoc.Properties != nil {
				if _, ok := route.AtDoc.Properties["roles"]; ok {
					roles := strings.Trim(route.AtDoc.Properties["roles"], "\"")
					roleArray := strings.Split(roles, ",")
					for _, role := range roleArray {
						roleInt, _ := strconv.Atoi(strings.TrimSpace(role))
						roleIntArray = append(roleIntArray, roleInt)
					}
				}
			}
			if len(roleIntArray) == 0 {
				roleIntArray = []int{1, 2, 3, 4}
			}
			authUnit.Roles = roleIntArray
			authUnits = append(authUnits, authUnit)
		}
	}
	return &AuthenticateResp{Authenticate: authUnits}, nil
}
