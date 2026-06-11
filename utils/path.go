package utils

import (
	"net/url"
	"strconv"
	"strings"
)

type PathParam struct {
	Name  string
	Value string
}

func BuildPath(template string, params ...PathParam) string {
	out := template
	for _, p := range params {
		out = strings.ReplaceAll(out, "{"+p.Name+"}", p.Value)
	}
	return out
}

func PathInt(name string, v int64) PathParam {
	return PathParam{Name: name, Value: strconv.FormatInt(v, 10)}
}

func PathInt32(name string, v int32) PathParam {
	return PathParam{Name: name, Value: strconv.FormatInt(int64(v), 10)}
}

func PathStringParam(name, v string) PathParam {
	return PathParam{Name: name, Value: url.PathEscape(v)}
}
