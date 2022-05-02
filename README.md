# udm-local-dns

`udm-local-dns` runs on UDM-SE and generates DNS entries for client devices.


## How it works
`udm-local-dns` polls Network controller API and generates new dnsmasq configuration if network client list changes.


## Compatibility
Only `UniFi OS - Dream Machine SE v2` is supported, i don't have any other devices to test.


## How to install 
1. Crate local user on udm
2. Enable ssh on udm
4. ssh to udm
5. Add APT repo and install package:
```
# wget -O - https://apt.gregory.beer/conf/gregory.beer.gpg.key | apt-key add -
# echo "deb https://apt.gregory.beer stretch utils" > /etc/apt/sources.list.d/gregory.beer.list
# apt-get update
# apt-get install udm-local-dns
```

6. Edit configuration, in most cases you only need to change username and password:
```
# vim /etc/local-dns.toml
```

7. Start and enable daemon:
```
# systemctl start udm-local-dns # start daemon 
# systemctl status udm-local-dns # check if daemon runs without errors
# systemctl enable udm-local-dns # enable daemon (start it automatically on boot)

```
7. Persist package across firmware update (optional, you can just reinstall package after update if you like).<br>
Edit `/etc/default/ubnt-dpkg-cache` and add `udm-local-dns` to `DPKG_CACHE_UBNT_PKGS`
Your file should looks something like this:
```
DPKG_CACHE=yes
DPKG_CACHE_DEBUG=yes
DPKG_CACHE_APT_UNPACK=yes
DPKG_CACHE_APT_REMOVE_AS_PURGE=yes
DPKG_CACHE_UBNT_PKGS="unifi unifi-protect ulp-go unifi-access unifi-talk unifi-connect uid-agent udm-local-dns"
```

