package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/yusufpapurcu/wmi"
	"www.velocidex.com/golang/go-prefetch"
)

var osInfo []Win32_OperatingSystem
var osProduct []Win32_ComputerSystemProduct

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

	err = wmi.Query("SELECT Caption, Version, SerialNumber FROM Win32_OperatingSystem", &osInfo)

	if err != nil {
		communicate_error("Failed to query WMI: " + err.Error())
	}

	err = wmi.Query("SELECT UUID FROM Win32_ComputerSystemProduct", &osProduct)
	if err != nil {
		communicate_error("Failed to query WMI: " + err.Error())
	}

	ram, _ := mem.VirtualMemory()
	ram_gb := float64(ram.Total) / 1e9
	ram_used := float64(ram.Used) / 1e9

	cpuInfo, _ := cpu.Info()
	cpuUsage, _ := cpu.Percent(0, false)
	cpuModel := cpuInfo[0].ModelName

	communicate_progress("Chopping away")
	journal_output, _ := exec.Command("fsutil", "usn", "readjournal", "c:", "csv").Output()

	apps := make(map[string]AppInfo)

	for identifier, prefetch_item := range finalized_prefetch_list {
		identifier = strings.ToUpper(identifier)
		icon_base, _ := exec.Command(dr+"\\bin\\icon_extractor.exe", paths[identifier]).Output()
		apps[identifier] = AppInfo{
			identifier,
			paths[identifier],
			string(icon_base),
			prefetch_item,
		}
	}

	hostname, _ := os.Hostname()

	export, _ := json.MarshalIndent(Export{"bundle", string(journal_output), apps, hostname, osInfo[0].Caption, osProduct[0].UUID, runtime.GOARCH, osInfo[0].Version, osInfo[0].SerialNumber, ram_gb, ram_used, cpuUsage, cpuModel}, " ", " ")

	communicate_progress("Timber!")

	fmt.Println(string(export) + "|!|/")
}
