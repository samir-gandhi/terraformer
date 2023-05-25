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
	"strconv"

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
			password := os.Getenv("PINGONE_PASSWORD")
			region := os.Getenv("PINGONE_REGION")
			environmentId := os.Getenv("PINGONE_ENVIRONMENT_ID")
			accessToken := os.Getenv("PINGONE_DAVINCI_ACCESS_TOKEN")
			targetEnvironmentId := options.Profile

			abstract := options.Abstract

			providerConfigVars := map[string]*string{
				"PINGONE_USERNAME":             &username,
				"PINGONE_PASSWORD":             &password,
				"PINGONE_DAVINCI_ACCESS_TOKEN": &accessToken,
				"PINGONE_REGION":               &region,
				"PINGONE_ENVIRONMENT_ID":       &environmentId,
			}
			pingonedavinci_terraforming.ReadPingOneConfig(options.Zone, providerConfigVars)

			if (len(username) == 0 || len(password) == 0) && len(accessToken) == 0 {
				return errors.New("PingOne Davinci credentials must be set through `PINGONE_USERNAME` and `PINGONE_PASSWORD` env vars or `PINGONE_DAVINCI_ACCESS_TOKEN` env var")
			}
			if len(region) == 0 {
				return errors.New("PingOne region for Davinci must be set through `PINGONE_REGION` env var")
			}
			if len(environmentId) == 0 {
				return errors.New("PingOne environemnt id for Davinci must be set through `PINGONE_ENVIRONMENT_ID` env var")
			}

			provider := newPingOneDavinciProvider()
			err := Import(provider, options, []string{username, password, accessToken, region, environmentId, targetEnvironmentId, strconv.FormatBool(abstract)})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newPingOneDavinciProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "", "")
	cmd.Flags().StringVarP(&options.Zone, "pingone-config-file", "F", "", "/full/path/to/pingone-config.env")
	cmd.Flags().StringVarP(&options.Profile, "target-environment-id", "t", "", "dv0-abc-1234-xyz-5678")
	return cmd
}

func newPingOneDavinciProvider() terraformutils.ProviderGenerator {
	return &pingonedavinci_terraforming.PingOneDavinciProvider{}
}
