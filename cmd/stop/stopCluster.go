/*
Copyright © 2020 The k3d Author(s)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package stop

import (
	"github.com/spf13/cobra"

	"github.com/rancher/k3d/v3/cmd/util"
	"github.com/rancher/k3d/v3/pkg/cluster"
	"github.com/rancher/k3d/v3/pkg/runtimes"
	k3d "github.com/rancher/k3d/v3/pkg/types"

	log "github.com/sirupsen/logrus"
)

// NewCmdStopCluster returns a new cobra command
func NewCmdStopCluster() *cobra.Command {

	// create new command
	cmd := &cobra.Command{
		Use:               "cluster  (NAME [NAME...] | --all)",
		Short:             "Stop an existing k3d cluster",
		Long:              `Stop an existing k3d cluster.`,
		ValidArgsFunction: util.ValidArgsAvailableClusters,
		Run: func(cmd *cobra.Command, args []string) {
			clusters := parseStopClusterCmd(cmd, args)
			if len(clusters) == 0 {
				log.Infoln("No clusters found")
			} else {
				for _, c := range clusters {
					if err := cluster.StopCluster(cmd.Context(), runtimes.SelectedRuntime, c); err != nil {
						log.Fatalln(err)
					}
				}
			}
		},
	}

	// add flags
	cmd.Flags().BoolP("all", "a", false, "Start all existing clusters")

	// add subcommands

	// done
	return cmd
}

// parseStopClusterCmd parses the command input into variables required to start clusters
func parseStopClusterCmd(cmd *cobra.Command, args []string) []*k3d.Cluster {
	// --all
	var clusters []*k3d.Cluster

	if all, err := cmd.Flags().GetBool("all"); err != nil {
		log.Fatalln(err)
	} else if all {
		clusters, err = cluster.GetClusters(cmd.Context(), runtimes.SelectedRuntime)
		if err != nil {
			log.Fatalln(err)
		}
		return clusters
	}

	if len(args) < 1 {
		log.Fatalln("Expecting at least one cluster name if `--all` is not set")
	}

	for _, name := range args {
		cluster, err := cluster.GetCluster(cmd.Context(), runtimes.SelectedRuntime, &k3d.Cluster{Name: name})
		if err != nil {
			log.Fatalln(err)
		}
		clusters = append(clusters, cluster)
	}

	return clusters
}
