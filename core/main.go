package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"www.velocidex.com/golang/go-prefetch"
)

func communicate_error(content string) {
	fmt.Println("{\"error\": true, \"err\": \"" + content + "\"}")
	os.Exit(1)
}

func communicate_progress(content string) {
	fmt.Println("{\"Type\": \"progress\", \"message\": \"" + content + "\"}|!|/")
}

func main() {
	communicate_progress("Some logs need chopping")

	prefetch_list, err := os.ReadDir("C:\\Windows\\Prefetch")

	if err != nil {
		communicate_error("Couldn't access prefetch directory, Admin permissions are required.")
		os.Exit(1)
	}

	finalized_prefetch_list := make(map[string]prefetch.PrefetchInfo)

	communicate_progress("Finding a forest")
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
		communicate_error("Stringifying PrefetchInfo failed")
	}

	dr, _ := os.Executable()

	dr = filepath.Dir(dr)

	communicate_progress("Sharpening our axe")
	_, err = exec.Command(dr + "\\bin\\totemp.bat").Output()

	if err != nil {
		fmt.Println(err)
		communicate_error("Couldn't retreive file paths for prefetch files")
	}

	bin_out, err := os.ReadFile(dr + "\\bin\\temp.txt")

	if err != nil {
		communicate_error("Couldn't retreive file paths for prefetch files")
	}

	line_split := strings.Split(string(bin_out), "\n")
	paths := make(map[string]string)

	communicate_progress("Scouting for trees")
	for _, line := range line_split {
		comma_split := strings.Split(line, ",")

		if len(comma_split) < 6 {
			continue
		}

		discovered_name := comma_split[5]
		discovered_path := comma_split[6]

		paths[discovered_name] = discovered_path
	}

	communicate_progress("Chopping away")
	journal_output, _ := exec.Command("fsutil", "usn", "readjournal", "c:", "csv").Output()

	apps := make(map[string]AppInfo)

	for identifier, prefetch_item := range finalized_prefetch_list {
		icon_base, _ := exec.Command(dr+"\\bin\\icon_extractor.exe", paths[strings.ToUpper(identifier)]).Output()
		apps[identifier] = AppInfo{
			paths[strings.ToUpper(identifier)],
			string(icon_base),
			prefetch_item,
		}
	}

	export, _ := json.MarshalIndent(Export{"bundle", string(journal_output), apps}, " ", " ")

	communicate_progress("Timber!")

	fmt.Println(string(export) + "|!|/")
}
