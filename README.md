# udm-local-dns

`udm-local-dns` runs on UDM-SE and generates DNS entries for client devices.


## How it works
`udm-local-dns` polls Network controller API and generates new dnsmasq configuration if network client list changes.


## Compatibility
Only `UniFi OS - Dream Machine SE v2` is supported, i don't have any other devices to test.


## How to install 
1. Crate local user on udm
2. Enable ssh on udm
3. Copy link from releases page
4. ssh to udm
5. Download deb package:
```
# wget https://deb-url-from/relase/page/
```

6. Install package:
```
# dpkg -i udm-local-dns_your-version_arm64.deb  
```

7. Edit configuration, in most cases you only need to change username and password:
```
# vim /etc/local-dns.toml
```

8. Start and enable daemon:
```
# systemctl start udm-local-dns # start daemon 
# systemctl status udm-local-dns # check if daemon runs without errors
# systemctl enable udm-local-dns # enable daemon (start it automatically on boot)

```
