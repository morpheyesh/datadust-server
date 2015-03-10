package main


import (
  "fmt"
  "github.com/megamsys/"
  "github.com/tsuru/config"
  "log"
  "os"
  "path/filepath"
  )

const (
  version = "0.1.0"
  header = "DataDust API Server"

)

const defaultConfigPath = "conf/dd.conf"

func buildManager(name string) *cmd.Manager {
	m := cmd.BuildBaseManager(name, version, header)
	m.Register(&IOTStart{})
	return m
}

func main() {
  p, _ := filepath.Abs(defaultConfigPath)
	log.Println(fmt.Errorf("Conf: %s", p))
	config.ReadConfigFile(defaultConfigPath)
	name := cmd.ExtractProgramName(os.Args[0])
	manager := buildManager(name)
	manager.Run(os.Args[1:])


}
