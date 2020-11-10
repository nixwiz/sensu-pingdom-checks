[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/nixwiz/sensu-pingdom-checks)
![Go Test](https://github.com/nixwiz/sensu-pingdom-checks/workflows/Go%20Test/badge.svg)
![goreleaser](https://github.com/nixwiz/sensu-pingdom-checks/workflows/goreleaser/badge.svg)

# Sensu Pingdom Checks

## Table of Contents
- [Overview](#overview)
  - [Attribution](#attribution)
  - [Checks](#checks)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Check definition](#check-definition)
  - [Environment variables](#environment-variables)
- [Usage examples](#usage-examples)
  - [check-pingdom-aggregates](#check-pingdom-aggregates)
  - [check-pingdom-credits](#check-pingdom-credits)
- [Installation from source](#installation-from-source)

## Overview

The Sensu Pingdom Checks is a collection of [Sensu Checks][1] that provide monitoring
of [Pingdom][2] services.

### Attribution

Portions of the pingdom package contained here are derived from the work of
[Russell Cardullo][3] in their [go-pingdom repository][4].

The logic and motivation for these checks are derived from the checks found in
the [sensu-plugins-pingdom][5] plugins.

### Checks

This collection contains the following checks:

* `check-pingdom-aggregates` - for checking if any Pingdom monitored sites have
a status of "down".
* `check-pingdom-credits` - for checking if available SMS and Checks credits
are of a sufficient value.

## Configuration

### Asset registration

[Sensu Assets][6] are the best way to make use of this plugin. If you're not
using an asset, please consider doing so! If you're using sensuctl 5.13 with
Sensu Backend 5.13 or later, you can use the following command to add the asset:

```
sensuctl asset add nixwiz/sensu-pingdom-checks
```

If you're using an earlier version of sensuctl, you can find the asset on the
[Bonsai Asset Index][7].

### Check definitions

#### check-pingdom-aggregates

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: check-pingdom-aggregates
  namespace: default
spec:
  command: check-pingdom-aggregates --warning 1 --critical 3
  subscriptions:
  - system
  runtime_assets:
  - nixwiz/sensu-pingdom-checks
```

#### check-pingdom-credits

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: check-pingdom-credits
  namespace: default
spec:
  command: >-
    check-pingdom-credits
    --critical-available-sms 5
    --warning-available-sms 10
    --critical-available-checks 5
    --warning-available-checks 10
  subscriptions:
  - system
  runtime_assets:
  - nixwiz/sensu-pingdom-checks
```

### Environment variables

The check definitions above assume the Pingdom API key is available via the
environment variable `PINGDOM_API_KEY`.  To keep from exposing it in the
check configuration, [you can set it on the agent(s)][8] that will be running
the check(s).

However, the preferable way to do this would be to use [secrets management][9]
with [mTLS agent authentication][10] to allow the agent(s) access to this value
as a secret.

## Usage examples

### check-pingdom-aggregates

#### Help output

```
Sensu check for number of down Pingdom checks

Usage:
  check-pingdom-aggregates [flags]
  check-pingdom-aggregates [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -k, --api-key string   API Key for connecting to Pingdom (PINGDOM_API_KEY env var)
  -c, --critical int     Critical threshold of down pingdom checks
  -w, --warning int      Warning threshold of down pingdom checks
  -h, --help             help for check-pingdom-aggregates

Use "check-pingdom-aggregates [command] --help" for more information about a command.
```

### check-pingdom-credits

#### Help output

```
Sensu check for available Pingdom credits

Usage:
  check-pingdom-credits [flags]
  check-pingdom-credits [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -k, --api-key string                  API Key for connecting to Pingdom (PINGDOM_API_KEY env var)
  -c, --critical-available-sms int      Critical threshold for available SMS messages (default -1)
  -w, --warning-available-sms int       Warning threshold for available SMS messages (default -1)
  -C, --critical-available-checks int   Critical threshold for available checks (default -1)
  -W, --warning-available-checks int    Warning threshold for available checks (default -1)
  -h, --help                            help for check-pingdom-credits

Use "check-pingdom-credits [command] --help" for more information about a command.
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an
Asset. If you would like to compile and install the plugin from source or
contribute to it, download the latest version or create an executable binary
from this source.

From the local path of the sensu-pingdom-checks repository:

```
go build
```

[1]: https://docs.sensu.io/sensu-go/latest/reference/checks/
[2]: https://www.pingdom.com/
[3]: https://github.com/russellcardullo
[4]: https://github.com/russellcardullo/go-pingdom
[5]: https://github.com/sensu-plugins/sensu-plugins-pingdom
[6]: https://docs.sensu.io/sensu-go/latest/reference/assets/
[7]: https://bonsai.sensu.io/nixwiz/sensu-pingdom-checks
[8]: https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-schedule/agent/#use-environment-variables-with-the-sensu-agent
[9]: https://docs.sensu.io/sensu-go/latest/guides/secrets-management/
[10]: https://docs.sensu.io/sensu-go/latest/operations/deploy-sensu/secure-sensu/#sensu-agent-mtls-authentication
