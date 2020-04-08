# blcli

```
blcli is a command-line interface for BitLaunch.io

Usage:
  blcli [command]

Available Commands:
  account        Retrieve account information
  create-options View images, sizes, and options available for a host when creating a new server.
  help           Help about any command
  server         Manage your virtual machines
  sshkey         Manage SSH Keys
  transaction    Manage transactions
  version        blcli version

Flags:
      --config string   config file (default is $HOME/.blcli.yaml)
  -h, --help            help for blcli
      --token string    API authentication token

Use "blcli [command] --help" for more information about a command.
```

## Installing `blcli`

### Downloading a Release

Visit the [Releases
page](https://github.com/bitlaunchio/blcli/releases) for the
[`blcli` GitHub project](https://github.com/bitlaunchio/blcli).

You can optionally move the `blcli` binary to your path. For example:

```
sudo mv ~/blcli /usr/local/bin
```

## Authentication

To use `blcli` you'll need an API access token. You can generate one in your BitLaunch account [API page](https://app.bitlaunch.io/account/api). More information is available at the [developer hub](https://developers.bitlaunch.io).

Once you have your token, you can either:

1. Specify it with each request:

`blcli --token TOKEN_HERE ...`

2. Set it as an environment variable:

`export BL_API_TOKEN=TOKEN_HERE`

## Examples

Here are a few examples of using `blcli`. More help is available with `blcli [command] -h` and further documentation is available at the [developer hub](https://developers.bitlaunch.io/)

* List all servers on your account:
```
blcli server list
```
* Create a server:
```
blcli server create --host bitlaunch --name test --region <region-id> --image <image-id> --size <size-id> --password b1Tl4uNCH!
```
* Create a new Lightning Network transaction:
```
blcli transaction create 20 BTC --lightning
```
* View your account and balance
```
blcli account
```
