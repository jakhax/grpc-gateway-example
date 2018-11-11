# Makefile

- A makefile will enable us to manage our project easily, in this case we will use a make file to handle the following tasks.


```
api                            Auto-generate grpc go sources
clean                          Remove previous builds
client                         Build the binary file for client
dep                            Get the dependencies
help                           Display this help screen
server                         Build the binary file for server
```
- Run `make {action}` to execute the action, running `make` without args will execute the actions specified in the makefile `all:` section.
