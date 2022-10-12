package os_validation

import (
	"log"
	"os"
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

func run_cmd_with_env(env string, cmd string, args ...string) error {

	command := exec.Command(cmd, args...)
	command.Env = os.Environ()
	command.Env = append(command.Env, env)

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

func TestApt(t *testing.T) {
	// Doesn't work
	Convey("Apt should be able to update its cache files", t, func() {
		run_cmd_with_env("DEBIAN_FRONTEND=noninteractive", "apt", "update")
	})
	Convey("Apt should be able to upgrade installed packages", t, func() {
		run_cmd("apt", "-o", "Dpkg::Options::=\"--force-confold\"", "dist-upgrade", "-q", "-y", "--force-yes")
	})
	Convey("Apt should be able to install a new package", t, func() {
		run_cmd("apt", "-o", "Dpkg::Options::=\"--force-confold\"", "install", "fortune-mod", "-q", "-y", "--force-yes")
	})
}
