package generate

type AuthenticateUnit struct {
	Method string `yaml:"Method"`     // Method: GET/POST/PUT/DELETE ...
	URL    string `yaml:"Url"`        // Url: example /api/vmp_leak/v1/leak
	Roles  []int  `yaml:"Roles,flow"` // Roles: example [ 1,2,3,4 ]
}

type AuthenticateResp struct {
	Authenticate []*AuthenticateUnit `yaml:"Authenticate"`
}
