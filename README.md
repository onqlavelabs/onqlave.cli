# Overview

Onqlave CLI is a cli executable package for developers to interact with Onqlave resources via command line interface.

# Release asset

- CLI release will contain a release note of the docker
  image [repository](https://github.com/onqlavelabs/onqlave.cli/pkgs/container/onqlavelabs%2Fonqlave.cli).
- CLI packages for **Linux**, **MacOS**, and **Windows** are automatically created and uploaded as assets in each
  release.

# Download Onqlave CLI

## Linux

- Download Linux executable file `onqlave-linux` from the release assets
- Grant executable permission to `onqlave-linux` file

```
sudo chmod +x onqlave-linux
```

- Make sure `onqlave-linux` is executable:

```
./onqlave-linux
```

## MacOS

- Download Linux executable file `onqlave-darwin` from the release assets
- Grant executable permission to `onqlave-darwin` file

```
sudo chmod +x onqlave-darwin
```

- Make sure `onqlave-darwin` is executable:

```
./onqlave-darwin
```

## Windows

- Download Windows executable file `onqlave-windows.exe` from the release assets
- Make sure `onqlave-windows.exe` is executable:

```
.\onqlave-windows.exe
```

## Docker

- Download the Docker image from the CLI release note

```
docker pull ghcr.io/onqlavelabs/onqlavelabs/onqlave.cli:{$version}
```

- Verify the Docker image

```
docker images
```

The Docker image for the CLI package should be visible in the Docker images list:

```
REPOSITORY                                   TAG         IMAGE ID       CREATED         
ghcr.io/onqlavelabs/onqlavelabs/onqlave.cli  {$version}  17a828917e85   45 hours ago
```

- Run the Docker image in interactive mode

```
docker run -it 17a828917e85
```

- Make sure `onqlave` CLI package inside the docker image is executable:

```
onqlave
```

## Installation Script

- An installation shell script is provided to download any specific Onqlave CLI version
- Download and execute the installation script:

```shell
curl -s "https://raw.githubusercontent.com/onqlavelabs/onqlave.cli/scripts/install.sh" | bash ${cli-version}
```

- For Windows users, [WSL](https://learn.microsoft.com/en-us/windows/wsl/install) is recommended before using this
  installation script

# How to use Onqlave CLI

## Configure environment

- Make sure you have a working Onqlave CLI executable
- Configure the environment by execute the following command:

```
onqlave config init
```

- Make sure the environment is configured by execute the following command:

```
onqlave config current
```

- A configured environment should be as follows:

```
┌───────────────────────────────────────────────────┐
│ Key          Value                                     │
│───────────────────────────────────────────────────│
│ ApiBaseUrl   https://api.onqlave.com                   │
│ ConfigPath   /home/$user/.config/onqlave/config        │
│ Env          prod                                      │
│ TenantId                                               │
│ TenantName                                             │
└───────────────────────────────────────────────────┘
```

## Start using the CLI

Full documentation of using the CLI can be found [here](https://docs.onqlave.com/guides/cli-guide/overview-cli/)
