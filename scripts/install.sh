#!/bin/sh

set -e

main() {
  # input version
  version=$1

  # declare variable
  linux_x86="Linux x86_64"
  linux_arch64="Linux arch64"
  darwin_x86="Darwin x86_64"
  darwin_arm64="Darwin arm64"

  # onqlave.all repository is in private mode so this download is not yet working
  cli_package=""
  cli_prefix="cli"
  download_url="https://github.com/onqlavelabs/onqlave.cli/releases/download"

  # validate version detail, version detail must be provided
  if [ -z "$version" ]; then
    echo "Error: CLI version is required" 1>&2
    exit 1
  fi

  # check operating system version
  if [ "$OS" = "Windows_NT" ]; then
    echo "Error: Official onqlave builds for Windows NT are not available" 1>&2
    exit 1
  fi

  # get cli package target based on os architecture
  case $(uname -sm) in
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
    echo "Error: Unknown os architecture" 1>&2
    exit 1
    ;;
  esac

  # prepare download link
  echo "${download_url}/${cli_prefix}${version}/${cli_package}"

  # download CLI package using curl
  if curl --output /dev/null --silent --head --fail "${download_url}/${cli_prefix}${version}/${cli_package}"; then
    curl -LJ "${download_url}/${cli_prefix}${version}/${cli_package}" >${cli_package}
  else
    echo "Error: CLI version ${version} is not valid"
    exit 1
  fi

  sudo cp $cli_package /usr/local/bin/onqlave
  sudo chmod +x /usr/local/bin/onqlave

  # finalize result
  echo "onqlave was installed successfully"
  echo "Run 'onqlave --help' to get started"
}

main "$@"
