package main

import "www.velocidex.com/golang/go-prefetch"

type AppInfo struct {
	path     string
	icon     string
	Prefetch prefetch.PrefetchInfo
}

type Export struct {
	Type    string
	Journal string
	Apps    map[string]AppInfo
}
