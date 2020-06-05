package rest

type Param struct {
	Query  map[string]interface{}
	Var    map[string]interface{}
	Offset int
	Limit  int
}

func NewParam() Param {
	return Param{
		Query:  map[string]interface{}{},
		Var:    map[string]interface{}{},
		Offset: 1,
		Limit:  10,
	}
}

func (it *Param) GetInclude() string {
	if it.Var["include"] != nil {
		return it.Var["include"].(string)
	}
	return ""
}
