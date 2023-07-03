#!/bin/sh

set -e

main() {
  # input version
  version=$1
  os_type=$(uname -sm)

  # declare variable
  linux_x86="Linux x86_64"
  linux_arch64="Linux arch64"
  darwin_x86="Darwin x86_64"
  darwin_arm64="Darwin arm64"

  cli_package=""
  cli_installed="onqlave"
  download_url="https://github.com/onqlavelabs/onqlave.cli/releases/download"

  # validate version detail, version detail must be provided
  if [ -z "$version" ]; then
    echo "Error: CLI version is required" 1>&2
    exit 1
  fi

  if [ "$os_type" = "MINGW64*" ]; then
    cli_package="onqlave-windows.exe"
  else
    echo "Error: Unknown os architecture $os_type" 1>&2
    exit 1
  fi

  # get cli package target based on os architecture
  case $os_type in
  "$linux_x86")
    cli_package="onqlave-linux"
    ;;
  "$linux_arch64")
    echo "Error: Official onqlave builds for Linux arch64 are not available" 1>&2
    exit 1
    ;;
  "$darwin_x86" | "$darwin_arm64")
    cli_package="onqlave-darwin"
    ;;
  *)
    if [ "$OSTYPE" == "msys" ]; then
      cli_package="onqlave-windows.exe"
    else
      echo "Error: Unknown os architecture" 1>&2
      exit 1
    fi
    ;;
  esac

  # prepare download link
  echo "Asset URL: ${download_url}/${version}/${cli_package}"

  # download CLI package using curl
  if curl --output /dev/null --silent --head --fail "${download_url}/${version}/${cli_package}"; then
    curl -LJ "${download_url}/${version}/${cli_package}" >${cli_package}
  else
    echo "Error: CLI version ${version} is not valid"
    exit 1
  fi

  if [ "$cli_package" != "onqlave-windows.exe" ]; then
    # copy onqlave executable file to PATH, applied only for Linux and macOS
    sudo mv $cli_package /usr/local/bin/onqlave
    sudo chmod +x /usr/local/bin/onqlave
  else
    cli_installed="./onqlave-windows.exe"
  fi

  echo "onqlave executable was installed successfully"
  echo "Run '$cli_installed --help' to get started"
}

main "$@"
