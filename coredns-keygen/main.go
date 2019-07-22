package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/miekg/dns"
)

func main() {
	for _, zone := range os.Args[1:] {
		key := &dns.DNSKEY{
			Hdr: dns.RR_Header{
				Name:   dns.Fqdn(zone),
				Class:  dns.ClassINET,
				Ttl:    3600,
				Rrtype: dns.TypeDNSKEY,
			},
			Algorithm: dns.ECDSAP256SHA256,
			Flags:     257,
			Protocol:  3,
		}
		priv, err := key.Generate(256)
		if err != nil {
			log.Fatal(err)
		}

		base := fmt.Sprintf("K%s+%03d.+%05d", key.Header().Name, key.Algorithm, key.KeyTag())
		if key.Header().Name == "." {
			base = fmt.Sprintf("K%s.+%03d.+%05d", key.Header().Name, key.Algorithm, key.KeyTag()) // have .. for th root zone
		}
		if err := ioutil.WriteFile(base+".key", []byte(key.String()+"\n"), 0644); err != nil {
			log.Fatal(err)
		}
		if err := ioutil.WriteFile(base+".private", []byte(key.PrivateKeyString(priv)), 0600); err != nil {
			log.Fatal(err)
		}
		fmt.Println(base) // output keys generate to stdout to mimic dnssec-keygen
	}
}
