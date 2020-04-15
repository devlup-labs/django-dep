package handler

import (
	"github.com/devlup-labs/django-dep/config"
	"github.com/devlup-labs/django-dep/types"
	"github.com/emicklei/go-restful"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var (
	BUF_LEN = 1024
	running = false
)

var Deploy = func(request *restful.Request, response *restful.Response) {
	spec := new(types.RequestPayload)
	if err := request.ReadEntity(spec); err != nil {
		log.Print(err.Error())
		_ = response.WriteHeaderAndEntity(http.StatusBadRequest, map[string]string{
			"error": "request is not valid",
		})
		return
	}

	if running {
		response.WriteHeaderAndEntity(
			http.StatusUnprocessableEntity,
			map[string]string{"reason": "Another deployment is already running"},
		)
	}

	if spec.Token == config.Token() {
		running = true
		cmd := exec.Command(config.ScriptPath(), spec.Reference, strings.Join(spec.Features, ""), strings.Join(spec.Restart, ""))
		pipeReader, pipeWriter := io.Pipe()
		cmd.Stdout = pipeWriter
		cmd.Stderr = pipeWriter
		go writeCmdOutput(response, pipeReader)
		_ = cmd.Run()
		_ = pipeWriter.Close()
		running = false
	} else {
		_ = response.WriteHeaderAndEntity(http.StatusForbidden, map[string]string{
			"reason": "Token is not valid",
		})
		return
	}
}

func writeCmdOutput(res http.ResponseWriter, pipeReader *io.PipeReader) {
	buffer := make([]byte, BUF_LEN)
	for {
		n, err := pipeReader.Read(buffer)
		if err != nil {
			_ = pipeReader.Close()
			break
		}

		data := buffer[0:n]
		_, _ = res.Write(data)
		if f, ok := res.(http.Flusher); ok {
			f.Flush()
		}
		//reset buffer
		for i := 0; i < n; i++ {
			buffer[i] = 0
		}
	}
}
