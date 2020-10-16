# blcli

blcli is a command-line interface for BitLaunch.io

```
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

```sh
sudo mv ~/blcli /usr/local/bin
```

### Building from source
```sh
git clone https://github.com/bitlaunchio/blcli.git
cd blcli
go get .
go build .
```

## Authentication

To use `blcli` you'll need an API access token. You can generate one in your BitLaunch account [API page](https://app.bitlaunch.io/account/api). More information is available at the [developer hub](https://developers.bitlaunch.io).

Once you have your token, you can either:

1. Specify it with each request:

```sh
blcli --token TOKEN_HERE ...
```

2. Set it as an environment variable:

```sh
export BL_API_TOKEN=TOKEN_HERE
```

## Examples

Here are a few examples of using `blcli`. More help is available with `blcli [command] -h` and further documentation is available at the [developer hub](https://developers.bitlaunch.io/)

* View your account and balance:
```sh
blcli account
```
* View your account usage:
```sh
blcli account usage --period 2020-09
```
* View your account history/activity:
```sh
blcli account history
```
* List all servers on your account:
```sh
blcli server list
```
* Create a server:
```sh
blcli server create --host bitlaunch --name test --region lon1 --image 10002 --size nibble-1024 --password b1Tl4uNCH!
```
* Restart a server:
```sh
blcli server restart aaaaaaaaaaabbbbbbbbbbbbb
```
* Rebuild a server:
```sh
blcli server rebuild aaaaaaaaaaabbbbbbbbbbbbb --image 10000 --description "Ubuntu 18.04 LTS"
```
* Resize a server:
```sh
blcli server resize aaaaaaaaaaabbbbbbbbbbbbb --size nibble-2048
```
* Create a new Lightning Network transaction:
```sh
blcli transaction create 20 BTC --lightning
```

