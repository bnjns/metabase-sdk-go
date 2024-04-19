// Package metabase provides the functionality to create a new client.
//
//	import (
//		"github.com/bnjns/metabase-sdk-go/metabase"
//	)
//	func main() {
//		authenticator, err := metabase.NewApiKeyAuthenticator("<api key>")
//		if err != nil {
//			panic(err)
//		}
//
//		client, err := metabase.NewClient("<host>", authenticator)
//		if err != nil {
//			panic(err)
//		}
//	}
package metabase
