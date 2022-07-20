# golang_selfstudy
README.md

To fix problems involving _<package> is not in GOROOT_:

#GOPATH MUST BE OUTSIDE OF GOROOT directory!!!
export GOPATH=/mnt/sda1/programming/gopath
export PATH=$PATH:$GOPATH/bin

export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin