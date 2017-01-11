package exec

import (
	"io/ioutil"
	"os/exec"

	"github.com/flyaways/storage/util/log/log"
)

func Exec(name string, arg ...string) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("[E]", r)
		}
	}()

	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err.Error())
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err.Error())
	}

	if err := cmd.Start(); err != nil {
		panic(err.Error())
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		panic(err.Error())
	}

	if len(bytesErr) != 0 {
		panic(err.Error())
	}

	_, err = ioutil.ReadAll(stdout)
	if err != nil {
		panic(err.Error())
	}

	if err := cmd.Wait(); err != nil {
		panic(err.Error())
	}
	return
}
