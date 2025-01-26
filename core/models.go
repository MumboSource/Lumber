package main

import "www.velocidex.com/golang/go-prefetch"

type AppInfo struct {
	path string
	icon string
}

type Export struct {
	Type         string
	PrefetchData map[string]prefetch.PrefetchInfo
	Journal      string
	Paths        map[string]string
}
