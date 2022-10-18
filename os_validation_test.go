package os_validation

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/zcalusic/sysinfo"
)

func run_cmd(cmd string, args ...string) error {

	command := exec.Command(cmd, args...)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
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

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()

	if err != nil {
		log.Fatal(err)
	}

	So(err, ShouldEqual, nil)

	return err
}

func run_cmd_without_check(cmd string, args ...string) error {

	command := exec.Command(cmd, args...)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()

	if err != nil {
		log.Print(err)
	}

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
	Convey("Apt should be able to update its cache files", t, func() {
		run_cmd_with_env("DEBIAN_FRONTEND=noninteractive", "apt", "update")
	})
	// This failed when there was an new kernel. That might be okay, we'll have to see how often it comes up. Maybe upgrade instead of dist-upgrade
	Convey("Apt should be able to upgrade installed packages", t, func() {
		run_cmd_with_env("DEBIAN_FRONTEND=noninteractive", "apt", "-o", "Dpkg::Options::=\"--force-confold\"", "dist-upgrade", "-q", "-y", "--force-yes")
	})
	Convey("Apt should be able to install a new package", t, func() {
		run_cmd_with_env("DEBIAN_FRONTEND=noninteractive", "apt", "-o", "Dpkg::Options::=\"--force-confold\"", "install", "fortune-mod", "-q", "-y", "--force-yes")
	})
}

func TestApk(t *testing.T) {
	Convey("apk should be able to update its cache files", t, func() {
		run_cmd("apk", "update")
	})
	// This failed when there was an new kernel. That might be okay, we'll have to see how often it comes up. Maybe upgrade instead of dist-upgrade
	Convey("apk should be able to upgrade installed packages", t, func() {
		run_cmd("apk", "upgrade")
	})
	Convey("apk should be able to install a new package", t, func() {
		run_cmd("apk", "add", "fortune")
	})
}

func TestNix(t *testing.T) {
	Convey("nix should be able to update its cache files", t, func() {
		run_cmd("nix-channel", "--update")
	})
	// This failed when there was an new kernel. That might be okay, we'll have to see how often it comes up. Maybe upgrade instead of dist-upgrade
	Convey("nix should be able to upgrade installed packages", t, func() {
		run_cmd("nix-env", "--upgrade")
	})
	Convey("nix should be able to install a new package", t, func() {
		run_cmd("nix-env", "-i", "fortune-mod")
	})
}

func linuxNetworkingTests(t *testing.T) {
	log.Print(runtime.GOOS)
	run_cmd_without_check("ip", "address")
	run_cmd_without_check("ip", "link")
	run_cmd_without_check("ip", "route")
	run_cmd_without_check("lspci")
	run_cmd_without_check("lshw")
}

func TestOSValidation(t *testing.T) {
	var si sysinfo.SysInfo
	si.GetSysInfo()
	log.Print(si.OS)
	log.Print(si.OS.Vendor)
	if runtime.GOOS == "linux" {
		linuxNetworkingTests(t)
	}

}
