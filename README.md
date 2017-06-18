# honeybits-win
The Windows version of [honeybits](https://github.com/0x4D31/honeybits), a simple tool to create breadcrumbs and honeytokens to lead the attackers to your honeypots!

### Features:
* Creating fake credentials in Windows Credential Manager
* Reading config from a remote Key/Value Store such as Consule or etcd

### TODO:
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