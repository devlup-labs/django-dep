package handler

import (
	"io"
	"net/http"
	"os/exec"
)

var (
	BUF_LEN = 1024
)

var Deploy = func(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("ls", "-al")
	pipeReader, pipeWriter := io.Pipe()
	cmd.Stdout = pipeWriter
	cmd.Stderr = pipeWriter
	go writeCmdOutput(w, pipeReader)
	cmd.Run()
	pipeWriter.Close()
}

func writeCmdOutput(res http.ResponseWriter, pipeReader *io.PipeReader) {
	buffer := make([]byte, BUF_LEN)
	for {
		n, err := pipeReader.Read(buffer)
		if err != nil {
			pipeReader.Close()
			break
		}

		data := buffer[0:n]
		res.Write(data)
		if f, ok := res.(http.Flusher); ok {
			f.Flush()
		}
		//reset buffer
		for i := 0; i < n; i++ {
			buffer[i] = 0
		}
	}
}
