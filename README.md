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
sudo sh -c "echo export PATH=$PATH:/usr/local/go/bin >> /etc/profile"
```

```console
# system-wide Go install
export PATH=$PATH:/usr/local/go/bin
```



[Go Documentation](https://go.dev/doc/)

[Go book](https://www.openmymind.net/assets/go/go.pdf)
