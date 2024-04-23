<div align="center">
<h1>Update Install</h1>
Easy way to install deb package for LInux Distributions.<br>
Never struggle with download links again.<br>
<h3><code>ui install</code></h3>
<br>
</div>

# Installation
## Build from source
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

## Install Pre-build file with script
```bash
sudo wget https://github.com/Update-Install/CLI/releases/latest/download/ui-cli_linux_amd64 -q --show-progress --progress=bar:force -O /usr/local/bin/ui
```

# Usage
## Commands
`ui --help` show command information  
`ui source` show the config file of ui  
`ui source -n {name} -u {URL}` set a source for a package  
`ui install` download the package file in config file's list then install it.  
`ui install {name}` download and install the specific package

## Example (Ubuntu)
This is example had been tested on Ubuntu 23.10, 22.04

**Step 1.** Add the link to config using `ui config` command

```bash
ui source -n discord -u https://discord.com/api/download?platform=linux&format=deb
```

**Step 2.** Install the all packages in config

```bash
ui install
```

**Step 3.** Install the specific package with package name in config
```bash
ui install discord
``` 
