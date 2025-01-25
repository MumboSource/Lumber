package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Akumzy/ipc"
	"www.velocidex.com/golang/go-prefetch"
)

var comms *ipc.IPC

/*

progress
request

*/

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

		export, err := json.MarshalIndent(finalized_prefetch_list, " ", " ")

		if err != nil {
			communicate_error("Stringifying PrefetchInfo failed...")
		}

		dr, _ := os.Getwd()
		out, _ := os.Create(dr + "\\yap.json")

		out.WriteString(string(export))

		fmt.Println("\n\nDone!")
		//comms.Send("finalized", string(export))

	}()

	comms.Start()
	fmt.Println("Hello World!")
}
