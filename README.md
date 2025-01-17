# go-tutorial-full
This is an example of how to use Go

### Installation

[Download and Install](https://go.dev/doc/install)

Once downloaded, cd to the location of that file and enter the following, replacing go1.22.3 with whatever version you downloaded:

```console 
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz
```

Now add the following to your /etc/profile for a system-wide install:

```console
sudo sh -c "echo 'export PATH=\$PATH:/usr/local/go/bin' >> /etc/profile"
```

Open /etc/profile file and this should be appended

```console
export PATH=$PATH:/usr/local/go/bin
```

if you make a mistake you can always use nano:

```console
sudo nano /etc/profile
```

The next step is  in ~/.profile

```bash 
# set PATH so it includes GOPATH/bin if it exists
if [ -x "$(command -v go)" ] && [ -d "$(go env GOPATH)/bin" ]; then
    PATH="$(go env GOPATH)/bin:$PATH"
fi
```

Reload files:

reloading your profile files with. /etc/profile and . ~/.profile, or just restart.

Run this to confirm installation:

```console
go version
```

if you see the version of go then all is ok. 

To test if all is working follow this tutorial [Go Example tutorial](https://go.dev/doc/tutorial/getting-started)

#### Example recap

create a folder with some name (example hello) and cd into that folder:

run:

```console
go mod init example/hello
 ```

this will create go.mod file. When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a go.mod file that tracks the modules that provide those packages. That go.mod file stays with your code, including in your source code repository. 


run:

```console
touch hello.go
```

Open file and paste:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

run:

```console 
go run .
// or
go run hello.go
```

### Resource

[Go Documentation](https://go.dev/doc/)

[Go book](https://www.openmymind.net/assets/go/go.pdf)
