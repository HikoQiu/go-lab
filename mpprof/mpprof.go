package mpprof

import (
	"flag"
	"os"
	"strings"
	"log"
	"runtime/pprof"
	"fmt"
)

var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")

func BeginCPUProfile() {
	flag.Parse()

	if *cpuProfile == "" {
		return
	}

	if !strings.Contains(*cpuProfile, ".prof") {
		log.Println("profile file ext invalid. (.prof)")
	}

	f, err := os.Create(*cpuProfile)
	if err != nil {
		log.Fatalln(err.Error())
	}

	pprof.StartCPUProfile(f)
}

func EndCPUProfile() {
	if *cpuProfile == "" {
		return
	}
	fmt.Println("Stop cpu profile...")

	pprof.StopCPUProfile()
}
