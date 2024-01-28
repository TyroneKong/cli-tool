package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "tracing the ip",
	Long:  `tracing the ip`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, v := range args {
				showData(v)
			}
		} else {
			fmt.Println("Please provide ip to trace")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

type Ip struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"location"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
}

func showData(str string) {
	url := "http://ipinfo.io/" + str + "/geo"

	response := getData(url)

	var data Ip

	if err := json.Unmarshal(response, &data); err != nil {
		log.Println("unable to unmarshall response")
	}
	c := color.New(color.FgHiMagenta).Add(color.Underline)
	c.Println("data found")

	log.Printf("\n IP: %s\n City: %s\n Region: %s\n Country: %s\n Location: %s\n Timezone: %s\n Postal: %s\n", data.IP, data.City, data.Region, data.Country, data.Location, data.Timezone, data.Postal)
}

func getData(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		log.Println("unable to trace ip")
	}
	responseByte, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable to read response")
	}
	return responseByte
}
