# coredns-keygen

## Name

*coredns-keygen* - generate a key pair suitable for signing DNS zones.

## Description

*coredns-keygen* generates a Common Signing Key for the purpose of signing zones. It has no options
and will generate a key with the ECDSAP256SHA256 algorithm (elliptic curve) and the KSK bit set.

## Syntax

~~~
coredns-keygen ZONES...
~~~

* **ZONES** zones it should generate keys for.

For each key pair the following files are created:

* `K<zone>.+<algorithm>+<keytag>.key` for the DNSKEY RR,
* `K<zone>.+<algorithm>+<keytag>.ds` for the DS RR, and,
* `K<zone>.+<algorithm>+<keytag>.private` for the private one.

For each generated key the base name of these file is printed to standard output once.

## Examples

Generate keys for example.org and example.net:

~~~
$ coredns-keygen example.org example.net
Kexample.org.+013+09787
Kexample.net.+013+00440
~~~

## Also See

dnssec-keygen(8) can also used to generate keys and supports more options. ldns-keygen(1) and
ldns-key2ds(1) or similar utilities.

See RFC 4033, 4034, 4035 for the DNSSEC specification.
