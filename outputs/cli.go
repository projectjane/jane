package outputs

import (
	"fmt"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/mmcquillan/hex/models"
	"github.com/mmcquillan/hex/parse"
)

type Cli struct {
}

func (x Cli) Write(message models.Message, config models.Config) {
	fmt.Print("\n")
	for _, output := range message.Outputs {
		if message.Debug && parse.EitherMember(config.Admins, message.Attributes["hex.user"], message.Attributes["hex.channel"]) {
			output.Response = output.Response + "\n\n[ Debug: " + output.Command + " ]"
		}
		if message.Attributes["hex.rule.format"] == "true" {
			if output.Success {
				color.Set(color.FgGreen)
			} else {
				color.Set(color.FgRed)
			}
			fmt.Print(output.Response, "\n")
			color.Unset()
		} else {
			fmt.Print(output.Response, "\n")
		}
	}
	if message.Debug {
		keys := make([]string, 0, len(message.Attributes))
		for key, _ := range message.Attributes {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		fmt.Printf("MESSAGE DEBUG (%d sec to complete)\n", message.EndTime-message.StartTime)
		for _, key := range keys {
			if strings.HasPrefix(key, "hex.var.") {
				fmt.Printf("  %s: '%s'\n", key, "********")
			} else {
				fmt.Printf("  %s: '%s'\n", key, message.Attributes[key])
			}
		}
	}
	fmt.Print("\n> ")
}
