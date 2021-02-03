package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/whyeasy/atlasdb-performance-advisor/internal"
)

var (
	config internal.Config
)

func init() {
	flag.StringVar(&config.LogFormat, "logFormat", os.Getenv("LOG_FORMAT"), "Default is logfmt, can be set to JSON.")
	flag.StringVar(&config.LogLevel, "logLevel", os.Getenv("LOG_LEVEL"), "Set different log level, default is Info.")
	flag.StringVar(&config.GroupID, "groupId", os.Getenv("GROUP_ID"), "Provide the Group ID, which is the Project ID within AtlasDB.")
	flag.StringVar(&config.PublicKey, "publicKey", os.Getenv("PUBLIC_KEY"), "Provide the Public Key of the created API key within AtlasDB")
	flag.StringVar(&config.PrivateKey, "privateKey", os.Getenv("PRIVATE_KEY"), "Provide the Private Key of the created API key within AtlasDB")
	flag.IntVar(&config.Since, "since", 24, "Provide amount of hours in the past you want to retrieve data for.")
}

func main() {
	if err := parseConfig(); err != nil {
		log.Error(err)
		flag.Usage()
		os.Exit(2)
	}
	initLogger()

	log.Info("Running AtlasDB slow queries logger")

	internal.GetData(config.GroupID, config.PublicKey, config.PrivateKey, config.Since)
}

func parseConfig() error {
	flag.Parse()
	var err error
	required := []string{"groupId", "publicKey", "privateKey"}
	flag.VisitAll(func(f *flag.Flag) {
		for _, r := range required {
			if r == f.Name && (f.Value.String() == "" || f.Value.String() == "0") {
				err = fmt.Errorf("%v is empty", f.Usage)
			}
		}
		if f.Name == "logFormat" && (f.Value.String() == "" || f.Value.String() == "0") {
			err = f.Value.Set("logfmt")
			if err != nil {
				log.Error(err)
			}
		}
		if f.Name == "logLevel" && (f.Value.String() == "" || f.Value.String() == "0") {
			err = f.Value.Set("info")
			if err != nil {
				log.Error(err)
			}
		}
	})
	return err
}

func initLogger() {
	if strings.EqualFold(config.LogFormat, "json") {
		log.SetFormatter(&log.JSONFormatter{})
	}
	ll, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		ll = log.DebugLevel
	}
	log.SetLevel(ll)
}
