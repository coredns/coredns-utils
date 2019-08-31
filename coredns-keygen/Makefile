ifeq (, $(shell which mmark))
    $(warning "No mmark in $$PATH, exiting, see github.com/mmarkdown/mmark")
all:
	@echo "noop"
else
all: coredns-keygen.8
endif

coredns-keygen.8: README.md
	@/bin/echo -e '%%%\n title = "coredns-keygen 8"\n' \
		'area = "CoreDNS"\n workgroup = "CoreDNS"\n%%%\n\n' > $@.header
	@cat $@.header $< > $@.md && rm $@.header
	mmark -man $@.md > $@ && rm $@.md
