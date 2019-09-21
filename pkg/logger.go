package logger

import (
  "flag"
  "os"
  "log"
  "time"
  "fmt"
)

var (
  Log      *log.Logger
  logdir="/var/log/docker/machine/"
  logname="csdriver.log"
)

func create_logger() {
  t := time.Now()
  stamp := t.Format("2006-01-02")
  os.MkdirAll(logdir, os.ModePerm)
  flag.Parse()
  logpath := logdir + stamp
  var file, err = os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  if err != nil {
    panic(err)
  }

  Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
  Info("Starting new logfile : " + logpath)
}

func init() {

  fmt.Println("initialized Apache Cloudstack driver logger module!")
  if os.Getenv("DMCS_LOGGING") == "1" {
    fmt.Println("following current OS environment, I am allowed to write my own logs!")
    create_logger()
  }


}

func do_log(level string, message string){

  if os.Getenv("DMCS_LOGGING") == "1" {
    Log.Printf(level + " " + message)
  }
  // recreate logger  and logfile if entered new day for logging

}

func Debug(message string) {
  do_log("DEBUG", message)
}

func Info(message string) {
  do_log("INFO", message)
}

func Error(message string) {
    do_log("ERROR", message)
}

func Warn(message string) {
    do_log("WARN", message)
}

func Trace(message string) {
    do_log("TRACE", message)
}
