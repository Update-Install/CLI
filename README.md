# Installation
## Install Pre-built Version
Please refers to the [installation guide](https://github.com/Update-Install?view_as=public#installation) to install pre-built version of Update Install.

## Build From Source
### Requirements
- [Go (1.22)](https://go.dev/doc/install)
- [upx](https://github.com/upx/upx/releases/latest)

### Installation Steps
Download the source files.
```bash
git clone https://github.com/Update-Install/CLI.git
```

Run the build script.
```bash
./build.sh
```

Then copy the executable file to `/usr/local/bin`
```bash
sudo cp ./dist/ui-cli_linux_amd64 /usr/local/bin/ui
```

# Usage

### Step 1
Add the link to config using `ui config` command. (Ex: discord)
- `-n`: Specify the name of the app you want to install.
- `-u`: Provide the url for downloading the deb file of your app.
```bash
ui source -n discord -u https://discord.com/api/download?platform=linux&format=deb
```

### Step 2.
Install the specific package with package name you specified. (Ex: discord)
```bash
ui install discord
```

## More Commands
`ui --help` show command information  
`ui source` show the config file of ui  
`ui source -n {name} -u {URL}` set a source for a package  
`ui install` download the package file in config file's list then install it.  
`ui install {name}` download and install the specific package
