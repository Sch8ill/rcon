# rcon

[![Go Report Card](https://goreportcard.com/badge/github.com/sch8ill/rcon)](https://goreportcard.com/report/github.com/sch8ill/rcon)

 A RCON client for executing commands on a remote RCON server.

 This Go-based RCON client allows Minecraft server administrators to send commands to their server over a remote RCON connection. With this client, you can easily send commands and view output on the console without having to log in directly to the server.

 This client implements the RCON protocol used by Minecraft servers and has the potential to work with other game servers as well, though it hasn't been extensively tested with them.

 To use the client, you will need to have access to the server's RCON password and IP address/hostname. Once connected, you can send commands through the client's terminal interface and receive output on the console like you would with the standard Minecraft server console.

## Installation

### Prebuild

 Download the latest build for your platform from the [latest release](https://github.com/Sch8ill/rcon/releases/latest "latest release")

### Build

 You can also build your own executable.
 This requires:

 ```txt
 git
 go v1.18 or higher
 make
 ```

 Run these commands to build the executable:

 ```bash
 git clone https://github.com/Sch8ill/rcon
 make -C rcon
 mv rcon/build/rcon-cli rcon-cli
 rm -rf rcon
 ```

## Usage

### CLI

 ```txt
 USAGE:
    rcon [global options] command [command options] [arguments...]
 
 COMMANDS:
    help, h  Shows a list of commands or help for one command
 
 GLOBAL OPTIONS:
    --address value, -a value   address of the server you want to connect to (localhost:25575 for example) (default: "localhost")
    --password value, -p value  password of the RCON server you want to connect to (default: "minecraft")
    --command value, -c value   a single command to be executed
    --timeout value, -t value   timeout for the connection to the server (default: 7s)
    --no-colors, --no-colours   if the cli should not output colors (default: false)
    --help, -h                  show help
    --version, -v               print the version
  ```

#### Example commands

 Move into the directory of the executable and open up a terminal.

 Open up an interactive RCON terminal:

 ```bash
 ./rcon-cli -a <the-servers-address> -p <the-servers-password>
 ```

 Run a single command on the server:

 ```bash
 ./rcon-cli -a <the-servers-address> -p <the-servers-password> -c <the-command-that-shall-be-executed>
 ```

 Take a look at the other [command flags](#usage) for more features.

### Libary

 Install the module using:

 ```bash
 go get github.com/sch8ill/rcon
 ```

 Import it:

 ```go
 import "github.com/sch8ill/rcon"
 ```

 Create a new RCON client:

 ```go
 client, err := rcon.Dial("localhost", "password", 5) // address, password, timeout
 if err != nil {
   panic(err)
 }
 ```

 Execute a command over RCON:

 ```go
 output, err := client.ExecuteCmd("<your-command>")
 if err != nil {
   panic(err)
 }
 
 fmt.Println(output)
 ```
