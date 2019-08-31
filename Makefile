VERSION:=0.5

all: coredns-keygen

.PHONY: coredns-keygen
coredns-keygen:
	( cd coredns-keygen; $(MAKE) )

.PHONY: debian
debian:
	export MY_APP_VERSION=$(VERSION)
	nfpm -f .nfpm.yaml pkg -t coredns-utils.deb
