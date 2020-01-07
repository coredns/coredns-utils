package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/miekg/dns"
)

var helpFlag = flag.Bool("h", false, "show short help message")
var zskFlag = flag.Bool("zsk", false, "generate zone signing key (zsk)")
var keyFlag uint16 = 257 // CSK/KSK

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s [OPTIONS] ZONE [ZONE]...\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Generate Keys for DNSSEC (default is CSK/KSK).\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if *helpFlag || len(flag.Args()) == 0 {
		flag.Usage()
		return
	}

	if *zskFlag {
		keyFlag = 256 // ZSK
	}

	for _, zone := range flag.Args() {
		key := &dns.DNSKEY{
			Hdr:       dns.RR_Header{Name: dns.Fqdn(zone), Class: dns.ClassINET, Ttl: 3600, Rrtype: dns.TypeDNSKEY},
			Algorithm: dns.ECDSAP256SHA256, Flags: keyFlag, Protocol: 3,
		}
		priv, err := key.Generate(256)
		if err != nil {
			log.Fatal(err)
		}

		ds := key.ToDS(dns.SHA256)

		base := fmt.Sprintf("K%s+%03d+%05d", key.Header().Name, key.Algorithm, key.KeyTag())
		if err := ioutil.WriteFile(base+".key", []byte(key.String()+"\n"), 0644); err != nil {
			log.Fatal(err)
		}
		if err := ioutil.WriteFile(base+".private", []byte(key.PrivateKeyString(priv)), 0600); err != nil {
			log.Fatal(err)
		}
		if err := ioutil.WriteFile(base+".ds", []byte(ds.String()+"\n"), 0644); err != nil {
			log.Fatal(err)
		}
		fmt.Println(base) // output keys generated to stdout to mimic dnssec-keygen
	}
}
