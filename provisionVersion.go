package main

const (
	version = "0"
	release = "0"
	hotfix  = "0"
	feature = "1"
)

/**
Version get the complete Library's version
*/
func Version() string {
	return version + "." + release + "." + hotfix + "." + feature
}
