// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	"errors"
	"os"

	pingonedavinci_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/pingonedavinci"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdPingOneDavinciImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pingonedavinci",
		Short: "Import current state to Terraform configuration from PingOneDavinci",
		Long:  "Import current state to Terraform configuration from PingOneDavinci",
		RunE: func(cmd *cobra.Command, args []string) error {
			username := os.Getenv("PINGONE_USERNAME")
			if len(username) == 0 {
				return errors.New("PingOne username for Davinci must be set through `PINGONE_USERNAME` env var")
			}
			password := os.Getenv("PINGONE_PASSWORD")
			if len(password) == 0 {
				return errors.New("PingOne password for Davinci must be set through `PINGONE_PASSWORD` env var")
			}
			region := os.Getenv("PINGONE_REGION")
			if len(region) == 0 {
				return errors.New("PingOne region for Davinci must be set through `PINGONE_REGION` env var")
			}
			environmentId := os.Getenv("PINGONE_ENVIRONMENT_ID")
			if len(environmentId) == 0 {
				return errors.New("PingOne environemnt id for Davinci must be set through `PINGONE_ENVIRONMENT_ID` env var")
			}
			targetEnvironmentId := options.Profile

			provider := newPingOneDavinciProvider()
			err := Import(provider, options, []string{username, password, region, environmentId, targetEnvironmentId})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newPingOneDavinciProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "", "")
	cmd.Flags().StringVarP(&options.Profile, "target-environment-id", "t", "", "dv0-abc-1234-xyz-5678")
	return cmd
}

func newPingOneDavinciProvider() terraformutils.ProviderGenerator {
	return &pingonedavinci_terraforming.PingOneDavinciProvider{}
}
