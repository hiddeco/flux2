/*
Copyright 2020 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/types"

	"github.com/fluxcd/flux2/internal/utils"
	notificationv1 "github.com/fluxcd/notification-controller/api/v1beta1"
)

var suspendAlertCmd = &cobra.Command{
	Use:   "alert [name]",
	Short: "Suspend reconciliation of Alert",
	Long:  "The suspend command disables the reconciliation of a Alert resource.",
	Example: `  # Suspend reconciliation for an existing Alert
  flux suspend alert main`,
	RunE: suspendAlertCmdRun,
}

func init() {
	suspendCmd.AddCommand(suspendAlertCmd)
}

func suspendAlertCmdRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Alert name is required")
	}
	name := args[0]

	ctx, cancel := context.WithTimeout(context.Background(), rootArgs.timeout)
	defer cancel()

	kubeClient, err := utils.KubeClient(rootArgs.kubeconfig, rootArgs.kubecontext)
	if err != nil {
		return err
	}

	namespacedName := types.NamespacedName{
		Namespace: rootArgs.namespace,
		Name:      name,
	}
	var alert notificationv1.Alert
	err = kubeClient.Get(ctx, namespacedName, &alert)
	if err != nil {
		return err
	}

	logger.Actionf("suspending Alert %s in %s namespace", name, rootArgs.namespace)
	alert.Spec.Suspend = true
	if err := kubeClient.Update(ctx, &alert); err != nil {
		return err
	}
	logger.Successf("Alert suspended")

	return nil
}
