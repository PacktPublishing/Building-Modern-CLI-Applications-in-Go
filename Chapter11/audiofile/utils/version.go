package utils

var VERSION string = "1.0.0"

func Version() string {
	if VERSION == "dev" {
		return "dev"
	}

	return VERSION
}
