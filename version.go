package strmut

const (
	major = 0
	minor = 1
	patch = 0
)

var Version = Format("v%d.%d.%d", major, minor, patch)
