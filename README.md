# Honeybits-win
A simple tool to create breadcrumbs and honeytokens, to lead the attackers to your honeypots!

The Linux version of this project: [honeybits](https://github.com/0x4D31/honeybits)

_Author: Adel "0x4D31" Karimi._

## Features:
* Creating fake credentials in Windows Credential Manager
* Reading config from a remote Key/Value Store such as Consule or etcd

## Requirements:
* [Go Lang 1.7+](https://golang.org/dl/)
* Viper (```go get github.com/spf13/viper```)
* crypt (```go get github.com/xordataexchange/crypt/config```)

## Usage:
```
> go run honeybits-win.go

  /\  /\___  _ __   ___ _   _| |__ (_) |_ ___
 / /_/ / _ \| '_ \ / _ \ | | | '_ \| | __/ __|
/ __  / (_) | | | |  __/ |_| | |_) | | |_\__ \
\/ /_/ \___/|_| |_|\___|\__, |_.__/|_|\__|___/
========================|___/=================

Failed reading remote config. Reading the local config file...
Local config file loaded.

[+] Generic credential created (192.168.1.66)
[+] Generic credential created (realco-AWS_SECRET_ACCESS_KEY-david)
[+] Domain credential created (domain01)
[+] Domain credential created (winsrv)
```

## TODO:
* Honeyfiles
  + Type 1 - honeytoken (monitored)
  + Type 2 - breadcrumb (containing false information)
  + Type 3 - beacon docs
* Content generator module for honeyfiles
* More traps, including:
  + AWS credentials file
  + Fake entries in CMD/PowerShell commands history
  + Fake browser history, bookmarks and saved passwords
  + Database files/backups: SQLite, MySQL
  + Confoguration, backup, and connection files such as RDP and VPN
  + MS Outlook Data file (.ost/.pst)
  + Hosts files (hosts, lmhosts)
  + Fake ARP entries
  + KeePass file with fake entries (.kdbx)
  + Registery keys (WinSCP, PuTTY, etc.)
  + Injected fake credentials in LSASS
* Documentation