package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mcuadros/go-defaults"
	"github.com/server/config"
	"github.com/server/pkg/file"
	"github.com/server/pkg/token"
	"gopkg.in/yaml.v3"
)

func main() {
	 if len(os.Args) < 2 {
		fmt.Println("Usage: <command> [options]")
		os.Exit(1)
	}	
	
	cmd := os.Args[1]

	switch cmd {
	 	case "gen-config":	
			genConfig()
		case "gen-jwks":
			genJWKs()
		case "-h", "--help", "help":
			fmt.Println("Commands: gen-jwks, gen-config")
			os.Exit(0)
		default:
			fmt.Printf("Unknown command: %q\n", cmd)
			os.Exit(1)
	 }

	os.Exit(1)
}

func genJWKs() {
	numOfKeys := flag.Int("gen-JWKs", 2, "Create JWKs keys")
	outPath := flag.String("o", "", "output location")

	path := envPath(*outPath)

	f, err := file.CreateFile(path + "jwk.txt", 0755)

	if (err != nil) {
		log.Fatal("failed to create file: %v", err)
	}

	defer f.Close()

	keyset, err := token.CreateJWKs(*numOfKeys)

	if (err != nil) {
		log.Fatal("failed to create keyset: %v", err)
	}

	err = json.NewEncoder(f).Encode(keyset)

	if (err != nil) {
		log.Fatal("failed to write to file: %v", err)
	}

	os.Exit(0)
}

func genConfig() {
	fileName := flag.String("gen-config", "config.yaml", "Create config file")
	outPath := flag.String("o", "", "output location")

	path := envPath(*outPath) + *fileName

	cfg := &config.Config{}
	defaults.SetDefaults(cfg)

	data, err := yaml.Marshal(cfg)

	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if !file.DirExists(path) {
			_ = os.MkdirAll(filepath.Dir(path), 0755)
		}
	}
	
	if file.Exist(path) {
		log.Fatal("config file already exists")
	}

	if err := os.WriteFile(path, data, 0655); err != nil {
		log.Fatal(err)
	}

	os.Exit(1)
}

func envPath(path string) (string) {
	if (path != "") {
		return path
	}

	out, err := exec.Command("go", "env", "GOMOD").Output()

	if (err != nil) {
		log.Fatal("failed to root project path: %v", err)
	}

	s := strings.Split(string(out), "/")
	
	if (len(s) < 2) {
		log.Fatal("needs to be absolute path: %v", err)
	}

	parts := s[:len(s) - 2]
	
	parts = append(parts, "infra", "api")

	path = strings.Join(parts, "/")	+ "/"	

	return path
}
