# golang_selfstudy

##To fix problems involving _<package> is not in GOROOT_:

_GOPATH MUST BE OUTSIDE OF GOROOT directory!!!_
export GOPATH=/mnt/sda1/programming/gopath
export PATH=$PATH:$GOPATH/bin

export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin