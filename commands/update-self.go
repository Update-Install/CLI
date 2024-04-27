package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"

	"ui/cli/module"
)

func UpdateSelf(c *cli.Context) error {
	var cmd = exec.Command("sudo", "-v")
	cmd.Stdout = nil
	err := cmd.Run()
	if err != nil {
		return err
	}

	color.Yellowln("Updating ui-cli")
	fmt.Println()

	filePath := module.DownloadFileToCache("https://github.com/Update-Install/CLI/releases/latest/download/ui-cli_linux_amd64")

	module.FullWidthMessage("installation log start", color.Gray)
	cmd = exec.Command("sudo", "mv", filePath, "/usr/local/bin/ui")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	color.Greenln("Moved", filePath, "to /usr/local/bin")
	cmd = exec.Command("sudo", "chmod", "+x", "/usr/local/bin/ui")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	color.Greenln("Changed permission for /usr/local/bin/ui")
	module.FullWidthMessage("installation log end", color.Gray)

	color.Greenf("Successfully updated ui-cli to ")
	cmd = exec.Command("/usr/local/bin/ui", "version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
