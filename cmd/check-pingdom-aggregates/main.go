package main

import (
	"fmt"

	"github.com/nixwiz/sensu-pingdom-checks/pingdom"
	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
)

// Config represents the check plugin config.
type Config struct {
	sensu.PluginConfig
	APIKey string
	Critical int
	Warning int
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "check-pingdom-aggregates",
			Short:    "Sensu check for number of down Pingdom checks",
			Keyspace: "sensu.io/plugins/check-pingdom-aggregates/config",
		},
	}

	options = []*sensu.PluginConfigOption{
		&sensu.PluginConfigOption{
			Path:      "api-key",
			Env:       "PINGDOM_API_KEY",
			Argument:  "api-key",
			Shorthand: "k",
			Default:   "",
			Secret:    true,
			Usage:     "API Key for connecting to Pingdom (PINGDOM_API_KEY env var)",
			Value:     &plugin.APIKey,
		},
		&sensu.PluginConfigOption{
			Path:      "critical",
			Env:       "",
			Argument:  "critical",
			Shorthand: "c",
			Default:   0,
			Usage:     "Critical threshold of down pingdom checks",
			Value:     &plugin.Critical,
		},
		&sensu.PluginConfigOption{
			Path:      "warning",
			Env:       "",
			Argument:  "warning",
			Shorthand: "w",
			Default:   0,
			Usage:     "Warning threshold of down pingdom checks",
			Value:     &plugin.Warning,
		},
	}
)

func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *types.Event) (int, error) {
	if plugin.Critical == 0 {
		return sensu.CheckStateUnknown, fmt.Errorf("No critical threshold supplied")
	}
	if plugin.Warning == 0 {
		return sensu.CheckStateUnknown, fmt.Errorf("No warning threshold supplied")
	}
	if plugin.Warning > plugin.Critical {
		return sensu.CheckStateUnknown, fmt.Errorf("Warning threshold cannot be greater than critical threshold")
	}
	return sensu.CheckStateOK, nil
}

func executeCheck(event *types.Event) (int, error) {
	client, err := pingdom.NewClientWithConfig(pingdom.ClientConfig{
		APIToken: plugin.APIKey,
	})
	if err != nil {
		return sensu.CheckStateCritical, err
        }

	checks, err := client.Checks.List()
	if err != nil {
		return sensu.CheckStateCritical, err
        }

	down := 0
	for _, check := range checks {
		if check.Status == "down" {
			down += 1
		}
	}

	if down >= plugin.Critical {
		fmt.Printf("%s CRITICAL - %d pingdom checks down\n", plugin.PluginConfig.Name, down)
		return sensu.CheckStateCritical, nil
	} else if down >= plugin.Warning {
		fmt.Printf("%s WARNING - %d pingdom checks down\n", plugin.PluginConfig.Name, down)
		return sensu.CheckStateWarning, nil
	}

	fmt.Printf("%s OK - %d pingdom checks down\n", plugin.PluginConfig.Name, down)
	return sensu.CheckStateOK, nil
}
