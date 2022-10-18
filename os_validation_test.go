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

func apkTests(t *testing.T) {
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

func aptTests(t *testing.T) {
	Convey("Apt should be able to update its cache files", t, func() {
		run_cmd_with_env("DEBIAN_FRONTEND=noninteractive", "apt-get", "update")
	})
	// This failed when there was an new kernel. That might be okay, we'll have to see how often it comes up. Maybe upgrade instead of dist-upgrade
	Convey("Apt should be able to upgrade installed packages", t, func() {
		run_cmd_with_env("DEBIAN_FRONTEND=noninteractive", "apt-get", "-o", "Dpkg::Options::=\"--force-confold\"", "dist-upgrade", "-q", "-y", "--allow-change-held-packages")
	})
	Convey("Apt should be able to install a new package", t, func() {
		run_cmd_with_env("DEBIAN_FRONTEND=noninteractive", "apt-get", "-o", "Dpkg::Options::=\"--force-confold\"", "install", "fortune-mod", "-q", "-y", "--allow-change-held-packages")
	})
}

func nixTests(t *testing.T) {
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

func yumTests(t *testing.T) {
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

func linuxNetworkingTests(t *testing.T) {
	run_cmd_without_check("ip", "address")
	run_cmd_without_check("ip", "link")
	run_cmd_without_check("ip", "route")
	run_cmd_without_check("lspci")
	run_cmd_without_check("lshw")
}

func TestOSValidation(t *testing.T) {
	log.Printf("Running tests for %s", runtime.GOOS)
	switch runtime.GOOS {
	case "linux":
		linuxNetworkingTests(t)
		var si sysinfo.SysInfo
		si.GetSysInfo()
		log.Printf("Running tests for %s", si.OS.Vendor)
		switch si.OS.Vendor {
		case "almalinux", "redhat":
			yumTests(t)
		case "debian", "ubuntu":
			aptTests(t)
		}
		log.Print(si.OS)
		log.Print(si.OS.Vendor)
	default:
		log.Printf("No tests implemented for os %s", runtime.GOOS)
	}
}
