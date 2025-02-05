/*

This is just to give some familiarity about env variables.
This appConfig module is NOT used anywhere else in the app

*/

package appConfig

import (
	"os"
	"log"
)

ServerHost := os.GetEnv("SERVER_HOST")

func init() {
	if ServerHost == "" {
		log.Fatal("Config variable not set SERVER_HOST")
	}
}