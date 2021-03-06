// Copyright (c) 2017 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"github.com/urfave/cli"
)

var versionCLICommand = cli.Command{
	Name:  "version",
	Usage: "display version details",
	Action: func(context *cli.Context) error {
		ctx, err := cliContextToContext(context)
		if err != nil {
			return err
		}

		span, _ := trace(ctx, "version")
		defer span.Finish()

		cli.VersionPrinter(context)
		return nil
	},
}
