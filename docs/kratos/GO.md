# Ory Kratos Go Snippets

These are for when you want to interact with Ory Kratos on a Go Client.

## Architecture

1. Allow auth req go to Ory Kratos
   1. Next.js uses the API integrations to help rewrite the routes to Kratos properly.
2. Protected req
   1. Make Go middleware
      1. Calls kratos to check the validity of the session (?)
   2. Protected endpoint happens

## Logout

```go
// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ory/client-go"
)

var ory *client.APIClient

func init() {
	cfg := client.NewConfiguration()
	cfg.Servers = client.ServerConfigurations{
		{URL: fmt.Sprintf("https://%s.projects.oryapis.com", os.Getenv("ORY_PROJECT_SLUG"))},
	}

	ory = client.NewAPIClient(cfg)
}

func logout(ctx context.Context, sessionToken string) error {
	_, err := ory.FrontendApi.PerformNativeLogout(ctx).
		PerformNativeLogoutBody(*client.NewPerformNativeLogoutBody(sessionToken)).
		Execute()
	if err != nil {
		return err
	}
	// Logout was successful

	return nil
}
```

## Email and Phone Verification

```go
// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"

	"github.com/ory/kratos/examples/go/pkg"

	ory "github.com/ory/client-go"
)

// If you use Open Source this would be:
//
// var client = pkg.NewSDKForSelfHosted("http://127.0.0.1:4433")
var client = pkg.NewSDK("playground")

func performVerification(email string) *ory.VerificationFlow {
	ctx := context.Background()

	// Initialize the flow
	flow, res, err := client.FrontendApi.CreateNativeVerificationFlow(ctx).Execute()
	pkg.SDKExitOnError(err, res)

	// If you want, print the flow here:
	//
	//	pkg.PrintJSONPretty(flow)

	// Submit the form
	afterSubmit, res, err := client.FrontendApi.UpdateVerificationFlow(ctx).Flow(flow.Id).
		UpdateVerificationFlowBody(ory.UpdateVerificationFlowWithLinkMethodAsUpdateVerificationFlowBody(&ory.UpdateVerificationFlowWithLinkMethod{
			Email:  email,
			Method: "link",
		})).Execute()
	pkg.SDKExitOnError(err, res)

	return afterSubmit
}

func main() {
	pkg.PrintJSONPretty(
		performVerification("someone@foobar.com"),
	)
}
```

## Account Recovery

```go
// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"

	"github.com/ory/kratos/examples/go/pkg"

	ory "github.com/ory/client-go"
)

// If you use Open Source this would be:
//
// var client = pkg.NewSDKForSelfHosted("http://127.0.0.1:4433")
var client = pkg.NewSDK("playground")

func performRecovery(email string) *ory.RecoveryFlow {
	ctx := context.Background()

	// Initialize the flow
	flow, res, err := client.FrontendApi.CreateNativeRecoveryFlow(ctx).Execute()
	pkg.SDKExitOnError(err, res)

	// If you want, print the flow here:
	//
	//	pkg.PrintJSONPretty(flow)

	// Submit the form
	afterSubmit, res, err := client.FrontendApi.UpdateRecoveryFlow(ctx).Flow(flow.Id).
		UpdateRecoveryFlowBody(ory.UpdateRecoveryFlowWithLinkMethodAsUpdateRecoveryFlowBody(&ory.UpdateRecoveryFlowWithLinkMethod{
			Email:  email,
			Method: "link",
		})).Execute()
	pkg.SDKExitOnError(err, res)

	return afterSubmit
}

func main() {
	pkg.PrintJSONPretty(
		performRecovery("someone@foobar.com"),
	)
}
```
