/*
Copyright (c) 2018 Red Hat, Inc.

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

// This example shows how to delete a cluster.

import (
	"fmt"
	"log"
	"os"

	"github.com/openshift-online/uhc-sdk-go/pkg/client"
)

func main() {
	// Create a logger that has the debug level enabled:
	logger, err := client.NewGoLoggerBuilder().
		Debug(true).
		Build()
	if err != nil {
		log.Fatalf("Can't build logger: %v", err)
	}

	// Create the connection, and remember to close it:
	token := os.Getenv("UHC_TOKEN")
	connection, err := client.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		Build()
	if err != nil {
		log.Fatalf("Can't create connection: %v", err)
	}
	defer connection.Close()

	// Send a request to delete a cluster:
	response, err := connection.Delete().
		Path("/api/clusters_mgmt/v1/clusters/1BDFg66jv2kDfBh6bBog3IsZWVH").
		Send()
	if err != nil {
		log.Fatalf("Can't delete cluster: %s", err)
	}

	// Print the result:
	fmt.Printf("%d\n", response.Status())
	fmt.Printf("%s\n", response.String())
}
