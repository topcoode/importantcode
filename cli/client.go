package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/urfave/cli"
)

type AMF struct{}

var amfCLi = []cli.Flag{
	cli.StringFlag{
		Name:  "free5gccfg",
		Usage: "common config file",
	},
	cli.StringFlag{
		Name:  "amfcfg",
		Usage: "amf config file",
	},
}
var (
	VERSION     string
	BUILD_TIME  string
	COMMIT_HASH string
	COMMIT_TIME string
)

func GetVersion() string {
	if VERSION != "" {
		return fmt.Sprintf(
			"\n\tfree5GC version: %s"+
				"\n\tbuild time:      %s"+
				"\n\tcommit hash:     %s"+
				"\n\tcommit time:     %s"+
				"\n\tgo version:      %s %s/%s",
			VERSION,
			BUILD_TIME,
			COMMIT_HASH,
			COMMIT_TIME,
			runtime.Version(),
			runtime.GOOS,
			runtime.GOARCH,
		)
	} else {
		return fmt.Sprintf(
			"\n\tNot specify ldflags (which link version) during go build\n\tgo version: %s %s/%s",
			runtime.Version(),
			runtime.GOOS,
			runtime.GOARCH,
		)
	}
}
func (*AMF) getCliCmd() (flags []cli.Flag) {
	return amfCLi
}

func main() {
	// guti := "20893cafe00000003e8"
	// fmt.Printf( "AMFId=%s\n", guti[len(guti)-12:])
	// fmt.Printf( "TMSI=%s\n", guti[len(guti)-8:])

	fmt.Printf("NumCPU=%d\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	app := cli.NewApp()
	fmt.Println("app----------------->", app)
	app.Name = "amf"
	fmt.Println("app.name----------->", app.Name)
	fmt.Println("AMF version: ", GetVersion())
	app.Usage = "-free5gccfg common configuration file -amfcfg amf configuration file"
	fmt.Println("app usage------->", app.Usage)
	//app.Flags = getCliCmd()
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("AMF Run error: %v", err)
		return
	}
}
