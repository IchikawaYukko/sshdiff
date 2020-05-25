# sshdiff

[日本語](READMEja.md)

Show differece of remote command output.

May be useful for server migration.

# Install
1. Download release executable
1. Place executable on /usr/local/bin or somewhere
1. Set permission to `rwxr-xr-x` (755)

# Usage
## Simple usage
```
$ sshdiff server1 server2 ls
```

run `ls` in server1 and server2 by ssh, then show difference of `ls` output.

## Specify ssh username
```
$ sshdiff yuriko@server1 yuriko@server2 ls
```

## Multiple command args
```
$ sshdiff yuriko@server1 yuriko@server2 'rpm -ql|sort'
````

If command has args, must quote it.

If ports(22) or keys(id_rsa, id_ecdsa, etc..) are different from default, please specify on your ~/.ssh/config like below

```
Host server1
HostName server1.example.com
Port 8022
IdentityFile ~/.ssh/id_ecdsa_server1

Host server2
HostName server2.example.com
Port 1022
IdentityFile ~/.ssh/id_ecdsa_server2
```