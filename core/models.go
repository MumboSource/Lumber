package main

import "www.velocidex.com/golang/go-prefetch"

type AppInfo struct {
	name     string
	path     string
	icon     string
	prefetch prefetch.PrefetchInfo
}

type Export struct {
	Type    string
	Journal string
	Apps    map[string]AppInfo

	Hostname       string
	WindowsEdition string
	DeviceId       string
	Architechture  string
	WindowsVersion string
	ProductId      string

	RamTotal float64
	RamUsed  float64

	CpuUsage []float64
	CpuModel string
}

type Win32_OperatingSystem struct {
	Caption      string
	Version      string
	SerialNumber string
}

type Win32_ComputerSystemProduct struct {
	UUID string
}
