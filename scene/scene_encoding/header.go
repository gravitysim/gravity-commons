package scene_encoding

const headerName = "gravityScene"
const version = "1.0.0"

type header struct {
	name    string
	version string
}

func createHeader() *header {
	return &header{name: headerName, version: version}
}
