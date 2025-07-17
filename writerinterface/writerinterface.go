package writerinterface

import (
	"io"
	"os"
)

func WriteToConsole(data string) {
	io.WriteString(os.Stdout, data)
}
