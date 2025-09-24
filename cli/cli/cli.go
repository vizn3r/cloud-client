package cli

import (
	"context"
	_ "embed"

	"github.com/urfave/cli/v3"
	"github.com/vizn3r/client/cli/http"
	"github.com/vizn3r/go-lib/conf"
	"github.com/vizn3r/go-lib/logger"
)

/*
* Commands:
*    - File operations
*   cloud ls/list          - List files in the cloud
*   cloud push/upload      - Push a file to the cloud
*   cloud rm/delete        - Remove a file from the cloud
*   cloud pull/download    - Pull a file from the cloud
*   cloud share            - Share a file from the cloud
*   cloud info             - Show information about a file
*
*    - P2P operations
*   cloud share --p2p      - Share a file to another P2P client
*   cloud pull --p2p       - Pull a file from another P2P client
*   cloud host             - Host a file to another P2P client
*   cloud sync             - Synchronize a file with another P2P client/server
*
*    - System commands
*   cloud help             - Show help
*   cloud version          - Show the version
*   cloud staatus          - Show the status
*   cloud config           - Show the configuration
*
*
 */

//go:embed cli.default.yaml
var defConf []byte
var log = logger.New("CORE", logger.Cyan)

type GlobalConf struct {
	HTTP *http.Conf `yaml:"http"`
}

func Commands() *cli.Command {
	cmd := &cli.Command{
		Name:  "cloud",
		Usage: "cloud storage thingy CLI",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Configuration file",
			},
			&cli.BoolFlag{
				Name:    "use_def_conf",
				Aliases: []string{"dc"},
				Usage:   "Use default configuration",
				Value:   true,
			},
		},
		Before: func(ctx context.Context, c *cli.Command) (context.Context, error) {
			confPath := c.String("config")
			useDefault := c.Bool("use_def_conf") && confPath == ""

			if useDefault {
				log.Info("Using default configuration")
				err := conf.LoadFromBytes[GlobalConf](defConf, "yaml")
				if err != nil {
					return ctx, err
				}
				return ctx, nil
			}

			err := conf.FindAndLoadConfig[GlobalConf](confPath)
			if err != nil {
				return ctx, err
			}

			log.Info("Using provided configuration")

			cfg := conf.Get[GlobalConf]()
			if cfg == nil {
				log.Fatal("Config data was not loaded properly")
			}
			http.Init(cfg.HTTP)

			return ctx, nil
		},
		Commands: []*cli.Command{
			{
				Name:  "auth",
				Usage: "Authenticate with the cloud",
				Commands: []*cli.Command{
					{
						Name:  "login",
						Usage: "Login to the cloud",
						Action: func(ctx context.Context, c *cli.Command) error {
							return nil
						},
					},
					{
						Name:  "logout",
						Usage: "Logout from the cloud",
						Action: func(ctx context.Context, c *cli.Command) error {
							return nil
						},
					},
					{
						Name:  "register",
						Usage: "Register with the cloud",
						Action: func(ctx context.Context, c *cli.Command) error {
							return nil
						},
					},
					{
						Name:  "whoami",
						Usage: "Show who you are",
						Action: func(ctx context.Context, c *cli.Command) error {
							return nil
						},
					},
				},
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List files in the cloud",
				Action: func(ctx context.Context, c *cli.Command) error {
					return nil
				},
			},
			{
				Name:    "upload",
				Aliases: []string{"push"},
				Usage:   "Push a file to the cloud",
				Action: func(ctx context.Context, c *cli.Command) error {
					return nil
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "Remove a file from the cloud",
				ArgsUsage: "<file>",
				Action: func(ctx context.Context, c *cli.Command) error {
					return nil
				},
			},
			{
				Name:    "download",
				Aliases: []string{"pull"},
				Usage:   "Pull a file from the cloud",
				Action: func(ctx context.Context, c *cli.Command) error {
					return nil
				},
			},
			{
				Name: "share",
				// Aliases: []string{"s"},
				Usage: "Share a file from the cloud",
				Action: func(ctx context.Context, c *cli.Command) error {
					return nil
				},
			},
		},
	}
	return cmd
}
