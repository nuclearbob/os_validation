package os_validation

import (
	"log"
	"os/exec"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func run_cmd(cmd string, args ...string) error {

	command := exec.Command(cmd, args...)

	err := command.Run()

	if err != nil {
		log.Fatal(err)
	}

	So(err, ShouldEqual, nil)

	return err
}

func TestYum(t *testing.T) {
	Convey("Yum should be able to update its cache files", t, func() {
		run_cmd("yum", "-y", "makecache")
	})
	Convey("Yum should be able to upgrade installed packages", t, func() {
		run_cmd("yum", "-y", "upgrade")
	})
	Convey("Yum should be able to install a new package", t, func() {
		run_cmd("yum", "-y", "install", "httpd")
	})
}
