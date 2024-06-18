package curlcommon

import (
	"log"
	"os"
	"os/exec"
)

/*
reference:
 https://stackoverflow.com/questions/69311943/why-is-the-curl-command-doesnt-work-in-golang
*/

func ExcecuteCurl(curl_command []string) error {
	cmd := exec.Command(curl_command[0], curl_command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return err
	}
	return nil
}
