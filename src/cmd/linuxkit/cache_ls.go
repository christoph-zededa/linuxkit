package main

import (
	"fmt"

	cachepkg "github.com/linuxkit/linuxkit/src/cmd/linuxkit/cache"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func cacheLsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "list images in the linuxkit cache",
		Long:  `List images in the linuxkit cache.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// list all of the images and content in the cache
			images, err := cachepkg.ListImages(cacheDir)
			if err != nil {
				log.Fatalf("error reading image names: %v", err)
			}
			fmt.Printf("%-80s %s\n", "image name", "root manifest hash")
			for name, hash := range images {
				fmt.Printf("%-80s %s\n", name, hash)
			}
			return nil
		},
	}
	return cmd
}
