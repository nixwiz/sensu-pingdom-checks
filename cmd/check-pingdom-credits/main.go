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
	CriticalAvailableSMS int
	WarningAvailableSMS int
	CriticalAvailableChecks int
	WarningAvailableChecks int
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "check-pingdom-credits",
			Short:    "Sensu check for available Pingdom credits",
			Keyspace: "sensu.io/plugins/check-pingdom-credits/config",
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
			Path:      "critical-available-sms",
			Env:       "",
			Argument:  "critical-available-sms",
			Shorthand: "c",
			Default:   -1,
			Usage:     "Critical threshold for available SMS messages",
			Value:     &plugin.CriticalAvailableSMS,
		},
		&sensu.PluginConfigOption{
			Path:      "warning-available-sms",
			Env:       "",
			Argument:  "warning-available-sms",
			Shorthand: "w",
			Default:   -1,
			Usage:     "Warning threshold for available SMS messages",
			Value:     &plugin.WarningAvailableSMS,
		},
		&sensu.PluginConfigOption{
			Path:      "critical-available-checks",
			Env:       "",
			Argument:  "critical-available-checks",
			Shorthand: "C",
			Default:   -1,
			Usage:     "Critical threshold for available checks",
			Value:     &plugin.CriticalAvailableChecks,
		},
		&sensu.PluginConfigOption{
			Path:      "warning-available-checks",
			Env:       "",
			Argument:  "warning-available-checks",
			Shorthand: "W",
			Default:   -1,
			Usage:     "Warning threshold for available checks",
			Value:     &plugin.WarningAvailableChecks,
		},
	}
)

func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *types.Event) (int, error) {
	if plugin.CriticalAvailableSMS == -1 {
		return sensu.CheckStateUnknown, fmt.Errorf("No critical threshold for available SMS messages")
	}
	if plugin.WarningAvailableSMS == -1 {
		return sensu.CheckStateUnknown, fmt.Errorf("No warning threshold for available SMS messages")
	}
	if plugin.WarningAvailableSMS < plugin.CriticalAvailableSMS {
		return sensu.CheckStateUnknown, fmt.Errorf("Warning threshold for available SMS messages must be greater than critical threshold")
	}
	if plugin.CriticalAvailableChecks == -1 {
		return sensu.CheckStateUnknown, fmt.Errorf("No critical threshold for available checks")
	}
	if plugin.WarningAvailableChecks == -1 {
		return sensu.CheckStateUnknown, fmt.Errorf("No warning threshold for available checks")
	}
	if plugin.WarningAvailableChecks < plugin.CriticalAvailableChecks {
		return sensu.CheckStateUnknown, fmt.Errorf("Warning threshold for available checks must be greater than critical threshold")
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

	credits, _ := client.Credits.Info()

	crit := false
	warn := false

	if credits.AvailableSMS <= plugin.CriticalAvailableSMS {
		crit = true
	} else if credits.AvailableSMS <= plugin.WarningAvailableSMS {
		warn = true
	}

	if credits.AvailableChecks <= plugin.CriticalAvailableChecks {
		crit = true
	} else if credits.AvailableChecks <= plugin.WarningAvailableChecks && !crit {
		warn = true
	}

	if crit {
		fmt.Printf("%s Critical - Available SMS=%d, Available Checks=%d\n", plugin.PluginConfig.Name, credits.AvailableSMS, credits.AvailableChecks)
		return sensu.CheckStateCritical, nil
	} else if warn {
		fmt.Printf("%s Warning - Available SMS=%d, Available Checks=%d\n", plugin.PluginConfig.Name, credits.AvailableSMS, credits.AvailableChecks)
		return sensu.CheckStateWarning, nil
	}

	fmt.Printf("%s OK - Available SMS=%d, Available Checks=%d\n", plugin.PluginConfig.Name, credits.AvailableSMS, credits.AvailableChecks)
	return sensu.CheckStateOK, nil

}
