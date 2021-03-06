// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package cli

import (
	"github.com/spf13/cobra"
	"github.com/sylabs/singularity/src/docs"
	"github.com/sylabs/singularity/src/pkg/libexec"
	"github.com/sylabs/singularity/src/pkg/sylog"
)

var (
	// PushLibraryURI holds the base URI to a Sylabs library API instance
	PushLibraryURI string
)

func init() {
	PushCmd.Flags().SetInterspersed(false)

	PushCmd.Flags().StringVar(&PushLibraryURI, "library", "https://library.sylabs.io", "")
	PullCmd.Flags().SetAnnotation("library", "envkey", []string{"LIBRARY"})

	SingularityCmd.AddCommand(PushCmd)
}

// PushCmd singularity push
var PushCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(2),
	PreRun:                sylabsToken,
	Run: func(cmd *cobra.Command, args []string) {
		// Push to library requires a valid authToken
		if authToken != "" {
			libexec.PushImage(args[0], args[1], PushLibraryURI, authToken)
		} else {
			sylog.Fatalf("Couldn't push image to library: %v", authWarning)
		}
	},

	Use:     docs.PushUse,
	Short:   docs.PushShort,
	Long:    docs.PushLong,
	Example: docs.PushExample,
}
