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
}
