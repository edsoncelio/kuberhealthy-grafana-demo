package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kuberhealthy/kuberhealthy/v2/pkg/checks/external/checkclient"
)

func checkURL(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}

func main() {
	checkclient.Debug = true
	grafanaUrl := os.Getenv("GRAFANA_URL")
	if len(grafanaUrl) != 0 {
		ok := checkURL(fmt.Sprintf("%s/api/health", grafanaUrl))
		if !ok {
			checkclient.ReportFailure([]string{"Failed to GET the Grafana API endpoint!"})
			return
		}
		checkclient.ReportSuccess()
	} else {
		log.Fatalln("Failed to parse GRAFANA_INTERNAL_URL env var")
	}

	log.Println("Check finished!")

}
