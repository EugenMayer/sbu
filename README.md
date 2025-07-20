# sbu - shelly-bulk-update

Automatically updates the firmware of all your [Shelly](https://shelly.cloud/) devices at once.


## Installation

Download the latest binary for your platform following the instructions on the [Releases](https://github.com/eugenmayer/shelly-bulk-update/releases) page.


## Usage

General flags:

- `--password`: (default: none) if provided, use as the authentication password. You can also use `~/.sbu.yml` instead
- `--gen`: 1,2,3 or "all" (default: all) - if provided, updates only shellys of the given generation 
- `--channel`: stable or beta (default: stable) - if provided, updates to the given update channel
- `--version`: show the current version

If any (or all) of your devices have authentication enabled, use the `--password` flags to define your
credentials (or use the configuration file in `~/.sbu.yml`)

```bash
./sbu --password MyPa$$w0rd
```

To update your Shelly devices to the latest beta version, use `--channel=beta`.

If you only want to update all Shelly devices of a specific device generation, use either `--gen=1` for
[generation 1](https://shelly-api-docs.shelly.cloud/gen1/#shelly-family-overview) or `--gen=2` for
[generation 2](https://shelly-api-docs.shelly.cloud/gen2/). For example, this can be used to update all second
generation devices to the latest beta version but keep first generation devices on the stable channel.

### Specific hosts

This will update both given hosts to the beta channel

```bash
./sbu --host 10.0.0.2 --host --10.0.0.3 --channel beta
```

### Autodiscovery

```bash
./sbu
```

It will automatically discover all your Shelly devices using mDNS and attempt to update them to the latest stable
version if possible.

Please note:
* The initial discovery can take up to 1 minute.
* While updates are in progress and devices are restarting, you might see connection errors. Sometimes it takes a few
  minutes, please be patient :-)

## Configuration

Create a file in your home folder called `~/.sbu.yml`.

### Authentication

You can provide authentication via `--password`

Create a file in your home folder called `~/.sbu.yml` and put this into our file

```yaml
default:
  credentials:
    username: admin # this should be admin - always. See https://shelly-api-docs.shelly.cloud/gen2/General/Authentication/#setting-authentication-credentials-for-a-device
    password: verysecret
```

# Credits

Credits to [shelly-bulk-update](https://github.com/fermayo/shelly-bulk-update) which this version was initially based on.
