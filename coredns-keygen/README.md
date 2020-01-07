# coredns-keygen

## Name

*coredns-keygen* - generate a key pair suitable for signing DNS zones.

## Description

*coredns-keygen* generates keys for the purpose of signing DNS zones. It has the option to
generate Zone Signing Key's (ZSK) however by default keys are generated with the KSK bit set.
All keys are generated with the ECDSAP256SHA256 algorithm (elliptic curve).

## Syntax

~~~sh
coredns-keygen [-zsk] ZONES...
~~~

* **-zsk**  generate ZSK instead of CSK/KSK
* **ZONES** zones it should generate keys for.

For each key pair the following files are created:

* `K<zone>.+<algorithm>+<keytag>.key` for the DNSKEY RR,
* `K<zone>.+<algorithm>+<keytag>.ds` for the DS RR, and,
* `K<zone>.+<algorithm>+<keytag>.private` for the private one.

For each generated key the base name of these file is printed to standard output once.

## Examples

Generate CSK/KSK keys for example.org and example.net:

~~~sh
$ coredns-keygen example.org example.net
Kexample.org.+013+09787
Kexample.net.+013+00440
~~~

Generate ZSK keys for example.org and example.net:

~~~sh
$ coredns-keygen -zsk example.org example.net
Kexample.org.+013+00234
Kexample.net.+013+08728
~~~

## Also See

dnssec-keygen(8) can also used to generate keys and supports more options. ldns-keygen(1) and
ldns-key2ds(1) or similar utilities.

See RFC 4033, 4034, 4035 for the DNSSEC specification.
