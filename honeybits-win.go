package main

import (
    "fmt"
    "os"
    "strings"
    "github.com/danieljoos/wincred"
    "github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)


var ASCII = `
  /\  /\___  _ __   ___ _   _| |__ (_) |_ ___ 
 / /_/ / _ \| '_ \ / _ \ | | | '_ \| | __/ __|
/ __  / (_) | | | |  __/ |_| | |_) | | |_\__ \
\/ /_/ \___/|_| |_|\___|\__, |_.__/|_|\__|___/
========================|___/=================
                                         	
`

func check(e error) {
	if e != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", e.Error()))
	}
}

// Load the local or remore config file. For remote config: 
// Read the config values from environment variables and then load
// the remote config (remote Key/Value store such as etcd or Consul)
// e.g. $ export HBITS_KVSPROVIDER="consul"
// 		$ export HBITS_KVSADDR="127.0.0.1:32775"
//		$ export HBITS_KVSDIR="/config/hbconf.yaml"
// 		$ export HBITS_KVSKEY="/etc/secrets/mykeyring.gpg"
func loadCon() (*viper.Viper, error) {
	conf := viper.New()
	conf.SetEnvPrefix("hbits")
	conf.AutomaticEnv()

	conf.SetDefault("kvsprovider", "consul")
	conf.SetDefault("kvsdir", "/config/hbconf.yaml")
	//conf.SetDefault("path.bashhistory", "~/.bash_history")

	kvsaddr := conf.GetString("kvsaddr")
	kvsprovider := conf.GetString("kvsprovider")
	kvsdir := conf.GetString("kvsdir")

	// If HBITS_KVSKEY is set, use encryption for the remote Key/Value Store
	if conf.IsSet("kvskey") {
		kvskey := conf.GetString("kvskey")
		conf.AddSecureRemoteProvider(kvsprovider, kvsaddr, kvsdir, kvskey)
	} else {
		conf.AddRemoteProvider(kvsprovider, kvsaddr, kvsdir)
	}
	conf.SetConfigType("yaml")

	if err := conf.ReadRemoteConfig(); err != nil {
		// Reading local config file
		fmt.Print("Failed reading remote config. Reading the local config file...\n")
		conf.SetConfigName("hbconf")
		conf.AddConfigPath(".")
		if err := conf.ReadInConfig(); err != nil {
			return nil, err
		}
		fmt.Print("Local config file loaded.\n\n")
		return conf, nil
	}
	fmt.Print("Remote config file loaded\n\n")
	return conf, nil
}

func cred_check(ctype string, ctarget string, cuser string) bool {
	var res bool
	if ctype == "generic" {
        if cred, err := wincred.GetGenericCredential(ctarget); err == nil && cred.UserName == cuser {
		    res = true
		}
	} else if ctype == "domain" {
        if cred, err := wincred.GetDomainPassword(ctarget); err == nil && cred.UserName == cuser {
            res = true
		}
	}
	return res
}

func cred_create(ctype string, ctarget string, cuser string, cpass string) {
	if ctype == "generic" {
		cred := wincred.NewGenericCredential(ctarget)
		cred.CredentialBlob = []byte(cpass)
		cred.UserName = cuser
		err := cred.Write()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("[+] Generic credential created (%s)\n", ctarget)
		}
	} else if ctype == "domain" {
		cred := wincred.NewDomainPassword(ctarget)
		// cred.CredentialBlob = []byte(cpass)
		//   error: The parameter is incorrect.  
		//   Ref: https://msdn.microsoft.com/en-us/library/windows/desktop/aa380517(v=vs.85).aspx#domain_credentials
		cred.UserName = cuser
		err := cred.Write()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("[+] Domain credential created (%s)\n", ctarget)
		}
	}
}

func main() {
    fmt.Print(ASCII)

    conf, err := loadCon()
    check(err)

    //Windows Credential Manager
    if conf.GetString("wincreds.enabled") == "true" {
    	if gcreds := conf.GetStringSlice("wincreds.generic-creds"); len(gcreds) != 0 {
				for _, g := range gcreds {
					gconf := strings.Split(g, ":")
					if cred_check("generic", gconf[0], gconf[1]) {
                        fmt.Printf("Error: generic credential exists (%s)\n", gconf[0])
					} else {
						cred_create("generic", gconf[0], gconf[1], gconf[2])
					}
				}
		}
		if dcreds := conf.GetStringSlice("wincreds.domain-creds"); len(dcreds) != 0 {
				for _, d := range dcreds {
					dconf := strings.Split(d, ":")
					if cred_check("domain", dconf[0], dconf[1]) {
                        fmt.Printf("Error: domain credential exists (%s)\n", dconf[0])
					} else {
					cred_create("domain", dconf[0], dconf[1], dconf[2])
				    }
				}
		}
    }
}