package cmd

import (
	"encoding/json"
	"github.com/amitizle/go-luas"
	"github.com/amitizle/luas-cli/internal/output"
	"github.com/spf13/cobra"
	"os"
)

var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get real time information, forecasting the next trams",
	Run:   forecast,
}

func init() {
	luasCmd.AddCommand(forecastCmd)
	forecastCmd.Flags().StringP("stop", "s", "", "stop identifier")
	forecastCmd.MarkFlagRequired("stop")
}

func forecast(cmd *cobra.Command, args []string) {
	stopName := cmd.Flag("stop").Value.String()
	stop, err := luas.GetStop(stopName)
	if err != nil {
		output.Errorf("error getting stop %s, %v", stopName, err)
		os.Exit(1)
	}
	stopInfo, err := stop.Forecast()
	if err != nil {
		output.Errorf("error getting forecast for stop %s, %v", stop.Name, err)
		os.Exit(1)
	}
	printOutput(cmd.Flag("format").Value.String(), stop, stopInfo)
}

func printOutput(format string, stop *luas.Stop, stopInfo *luas.StopInfo) {
	switch format {
	case "table":
		output.Infof("Forecast for stop %s from %s\n\n%s\n",
			stop.Name,
			stopInfo.Created,
			stopInfo.Message)
		for _, direction := range stopInfo.Directions {
			output.Infof("Direction: %s", direction.Name)
			for _, tram := range direction.Trams {
				if tram.Destination == "No trams forecast" {
					output.Warnf("No trams")
				} else if tram.DueMins == "DUE" {
					output.Infof("Destination: %s, due to arrive", tram.Destination, tram.DueMins)
				} else {
					output.Infof("Destination: %s, due in %s minutes", tram.Destination, tram.DueMins)
				}
			}
			output.Infof("")
		}
	case "json":
		jsonOutput, err := json.Marshal(stopInfo)
		if err != nil {
			output.Errorf("while generating json output: %v", err)
			os.Exit(1)
		}
		output.NoFormat(string(jsonOutput))
	default:
	}
}
