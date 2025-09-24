module github.com/vizn3r/cloud/cli

go 1.25.0

require github.com/vizn3r/cloud/server v0.0.0

require github.com/vizn3r/cloud/lib v0.0.0

require github.com/urfave/cli/v3 v3.4.1 // indirect

replace github.com/vizn3r/cloud/server => ../server

replace github.com/vizn3r/cloud/lib => ../lib
