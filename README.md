<div align="center">
<h1>Update-Install CLI</h1>
Easy way to install package for LInux Distributions<br>
<code>ui install</code>
</div>

# Commands
`ui --help` show command information  
`ui config` show the config file of ui  
`ui config -n {name} -u {URL}` set a source for a package  
`ui install` download the package file in config file's list then install it.  
`ui install {name}` download and install the specific package

# Installation
## Build from source
### Requirements
- Go (1.22)

### Installation Steps
Download the source files
```bash
git clone https://github.com/Update-Install/CLI.git
```
Install packages and build
```
go install
go build -o ui
```
Then just add the executable file to PATH

## Install Pre-build file with script
```bash
sudo wget https://github.com/Update-Install/CLI/releases/latest/download/ui-cli_linux_amd64 -q --show-progress --progress=bar:force -O /usr/local/bin/ui
```

# Example
## Ubuntu
This is example had been tested on Ubuntu 23.10

**Step 1.** Add the link to config using `ui config` command

```bash
ui config -n discord -u https://discord.com/api/download?platform=linux&format=deb
```

**Step 2.** Install the all packages in config

```bash
ui install
```

**Step 3.** Install the specific package with package name in config
```bash
ui install discord
``` 
