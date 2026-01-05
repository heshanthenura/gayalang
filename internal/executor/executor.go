package executor

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/heshanthenura/gayalang/internal/ast"
)

type Context map[string]string

func ExecuteProgram(prog *ast.Program) Context {
	ctx := Context{}

	for _, req := range prog.Requests {
		status, body := executeRequest(req, ctx)
		if req.Expect.Status != 0 && status != req.Expect.Status {
			fmt.Printf("Request %s failed: expected %d but got %d\n", req.Name, req.Expect.Status, status)
		} else {
			fmt.Printf("Request %s succeeded: %d\n", req.Name, status)
		}

		if req.SaveVar != "" {
			ctx[req.SaveVar] = body
			fmt.Printf("Saved %s = %s\n", req.SaveVar, body)
		}
	}

	return ctx
}

func executeRequest(req ast.RequestNode, ctx Context) (int, string) {
	var resp *http.Response
	var err error

	client := &http.Client{}

	switch req.Method {
	case "GET":
		resp, err = client.Get(req.URL)
	case "POST":
		resp, err = client.Post(req.URL, "application/json", bytes.NewBuffer([]byte{}))
	default:
		fmt.Printf("Unsupported method %s\n", req.Method)
		return 0, ""
	}

	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return 0, ""
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := string(bodyBytes)

	return resp.StatusCode, body
}
