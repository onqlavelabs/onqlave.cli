# Overview
Onqlave CLI is a cli executable package for developers to interact with Onqlave resources via command line interface.

# Release note and asset

- CLI release will contain a release note of the docker image repository.
- CLI packages for **Linux**, **MacOS**, and **Windows** are automatically created and uploaded as release assets by GitHub actions.

![release-note-and-asset](https://t36712295.p.clickup-attachments.com/t36712295/20cd37f4-56ac-447f-9edc-28101c44cd18/image.png)

# How to use CLI executable

## Linux
- Download Linux executable file `onqlave-linux` from the release assets
- Grant executable permission to `onqlave-linux` file
```
sudo chmod +x onqlave-linux
```
- Execute `onqlave-linux` using the following
```
./onqlave-linux
```
## MacOS
- Download Linux executable file `onqlave-darwin` from the release assets
- Grant executable permission to `onqlave-darwin` file
```
sudo chmod +x onqlave-darwin
```
- Execute `onqlave-darwin` using the following
```
./onqlave-darwin
```
## Windows
- Download Windows executable file `onqlave-windows.exe` from the release assets
- Execute `onqlave-windows.exe` using the following
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

