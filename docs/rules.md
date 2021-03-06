# Rules

Each rule is a seperate json file which consists of some rule options and a series of actions to execute if the rule matches. A rule will generally be either a match from slack input, a schedule or a web url.

json:
```
{
  "rule": "Say Hello",
  "match": "hello",
  "actions": [
    {
      "type": "hex-response",
      "command": "Hello ${hex.user}!"
    }
  ]
}
```

yaml:
```
---
rule: "Say Hello"
match: "hello"
actions:
- type: "hex-response"
  command: "Hello ${hex.user}!"
```

### Rule
- Config: `rule`
- Description: Name of the rule, also used as the title for formatted output
- Default:
- Type: string

### Match
- Config: `match`
- Description: String to match with * as wild card or /../ as regular expression
- Default:
- Type: string

### Schedule
- Config: `schedule`
- Description: Cron style schedule with seconds
- Default:
- Type: string

### URL
- Config: `url`
- Description: URL to match for incoming webhooks
- Default:
- Type: string

### ACL
- Config: `acl`
- Description: A comma delimited list of users and channels allowed to execute the rule
- Default: *
- Type: string

### Channel
- Config: `channel`
- Description: A channel to send output to if a schedule or if you want a place for all output to be copied to
- Default:
- Type: string

### Format
- Config: `format`
- Description: Flag to format output when displaying in slack
- Default: false
- Type: bool

### Threaded
- Config: `threaded`
- Description: Flag to set if the messages for each action should be threaded
- Default: false
- Type: bool

### OutputFailOnly
- Config: `output_fail_only`
- Description: Flag to only output if a failure
- Default: false
- Type: bool

### OutputOnChange
- Config: `output_on_change`
- Description: Flag to output on any status change
- Default: false
- Type: bool

### GroupOutput
- Config: `group_output`
- Description: Flag to output all actions as one message
- Default: false
- Type: bool

### Lock
- Config: `lock`
- Description: Flag to lock running to one at a time
- Default: false
- Type: bool

### Help
- Config: `help`
- Description: Custom help to display for the rule
- Default:
- Type: string

### Hide
- Config: `hide`
- Description: Flag for displaying help or not when user lists commands
- Default: false
- Type: bool

### Active
- Config: `active`
- Description: Flag for if the rule is run or not
- Default: true
- Type: bool

### Debug
- Config: `debug`
- Description: Flag for extra debug output in the logs
- Default: false
- Type: bool

### Actions.Type
- Config: `type`
- Description: The type of action, also known as the plugin to execute
- Default:
- Type: string

### Actions.Command
- Config: `command`
- Description: The command to give the plugin to resolve
- Default:
- Type: string

### Actions.HideOutput
- Config: `hide_output`
- Description: An option to not show the output of an action
- Default: false
- Type: bool

### Actions.OutputToVar
- Config: `output_to_var`
- Description: An option to take the output of the action and save it to the `hex.outputs.<action number>.response` var
- Default: false
- Type: bool

### Actions.OutputFailOnly
- Config: `output_fail_only`
- Description: Flag to only output action if a failure
- Default: false
- Type: bool

### Actions.RunOnFail
- Config: `run_on_fail`
- Description: A flag to let this action run if previous steps have failed
- Default: false
- Type: bool

### Actions.LastConfig
- Config: `last_config`
- Description: A flag to reuse the previous actions configuration (as a time saver)
- Default: false
- Type: bool

### Actions.Config
- Config: `config`
- Description: A set of key/value string pairs that can be plugin specific
- Default:
- Type: string key/value

