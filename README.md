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

### Resource

[Go Documentation](https://go.dev/doc/)

[Go book](https://www.openmymind.net/assets/go/go.pdf)
