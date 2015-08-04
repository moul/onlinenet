package api

import (
	"strconv"
	"strings"
)

type ServerPath string

type ListServersResp []ServerPath

func (r *ServerPath) Identifier() int {
	idStr := strings.Split(string(*r), "/")[4]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	return idInt
}
