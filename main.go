package main

import (
	"fmt"
	"os"

	"kubeadm/app"
	"kubeadm/app/util"
)

func main() {
	if err := app.Run(); err != nil {
		fmt.Printf(util.AlphaWarningOnExit)
		os.Exit(1)
	}
	os.Exit(0)
}
