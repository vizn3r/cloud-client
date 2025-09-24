module github.com/vizn3r/cloud/daemon

require (
	github.com/vizn3r/cloud/lib v0.0.0
	github.com/vizn3r/cloud/server v0.0.0
)

replace github.com/vizn3r/cloud/server => ../server

replace github.com/vizn3r/cloud/lib => ../lib

go 1.25.0
