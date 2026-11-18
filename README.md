# Pigeon_post
Common repo for shared server/client modules: protocol, types, config &amp; provisioning. 

## Setup
The config folder contains an example config.json file.

To use the pigeon_post module, copy config.json into your application's working directory and update the configuration values as required.

The protocol, host, and port settings must match between the client and server instances; otherwise a connection cannot be established.

Example:
```
{
    "protocol": "tcp",
    "host": "127.0.0.1",
    "port": "8080"
}
```
