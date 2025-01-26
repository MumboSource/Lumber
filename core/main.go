package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Akumzy/ipc"
	"www.velocidex.com/golang/go-prefetch"
)

var comms *ipc.IPC

type Export struct {
	PrefetchData map[string]prefetch.PrefetchInfo
	Paths        map[string]string
}

func communicate_error(content string) {
	comms.Send("err", content)
	os.Exit(1)
}

func main() {
	comms = ipc.New()

	go func() {
		comms.Send("heartbeat", "It lives!") // Notify the frontend that the core is starting
		comms.Send("progress", "Core loaded, extracting Prefetch...")

		prefetch_list, err := os.ReadDir("C:\\Windows\\Prefetch")

		if err != nil {
			communicate_error("Couldn't access prefetch directory, Admin permissions are required.")
			os.Exit(1)
		}

		finalized_prefetch_list := make(map[string]prefetch.PrefetchInfo)

		for _, file_entry := range prefetch_list {
			file_name := file_entry.Name()

			if !strings.HasSuffix(file_name, ".pf") {
				continue
			}

			handle, err := os.Open("C:\\Windows\\Prefetch\\" + file_name)

			if err != nil {
				communicate_error(fmt.Sprintf("Couldn't load prefetch file (%s)", file_name))
			}

			pf, err := prefetch.LoadPrefetch(handle)

			if err != nil {
				communicate_error(fmt.Sprintf("Could parse prefetch file (%s)", file_name))
			}

			finalized_prefetch_list[strings.Split(file_name, ".")[0]] = *pf
		}

		if err != nil {
			communicate_error("Stringifying PrefetchInfo failed...")
		}

		dr, _ := os.Getwd()

		_, err = exec.Command(dr + "\\wpv\\totemp.bat").Output()

		if err != nil {
			fmt.Println(err)
			communicate_error("Couldn't retreive file paths for prefetch files")
		}

		wpv_out, err := os.ReadFile(dr + "\\wpv\\temp.txt")

		if err != nil {
			communicate_error("Couldn't retreive file paths for prefetch files")
		}

		line_split := strings.Split(string(wpv_out), "\n")
		paths := make(map[string]string)

		for _, line := range line_split {
			comma_split := strings.Split(line, ",")

			if len(comma_split) < 6 {
				continue
			}

			fmt.Println(comma_split)

			discovered_name := comma_split[5]
			discovered_path := comma_split[6]

			paths[discovered_name] = discovered_path
		}

		fmt.Println("\n\nDone!")

		export, _ := json.MarshalIndent(Export{finalized_prefetch_list, paths}, " ", " ")
		comms.Send("finalized", string(export))
	}()

	comms.Start()
	fmt.Println("Hello World!")
}
