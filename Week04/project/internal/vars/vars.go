package vars

import "fmt"

type Vars struct {
	Message string
	Params  map[string]interface{}
	params  []string
}

func NewVars(msg string) *Vars {
	vs := new(Vars)
	vs.Message = msg
	vs.Params = make(map[string]interface{})
	return vs
}

func (vs *Vars) Set(key string, value interface{}) {
	vs.Params[key] = value
	vs.params = append(vs.params, key)
}

func (vs *Vars) Get(key string) string {
	v, ok := vs.Params[key]
	if ok {
		return v.(string)
	}
	return ""
}

func (vs *Vars) String() string {
	msg := vs.Message
	for _, v := range vs.params {
		msg += " " + fmt.Sprintf("%s=\"%v\"", v, vs.Params[v])
	}
	return msg
}
