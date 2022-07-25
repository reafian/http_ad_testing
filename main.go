package main

import (
    "encoding/json"
    "net/http"
    "bufio"
    "fmt"
    "log"
    "os"
)

type configuration struct {
    Prefix      string
    Suffix      string
}

func main() {
    var config string = "conf.json"
    var macs string = "mac_list"
    prefix, suffix := fixes(config)

    maclist, err := os.Open(macs)
    if err != nil {
        log.Fatal(err)
    }
    defer maclist.Close()

    scanner := bufio.NewScanner(maclist)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
        resp, err := http.Get(prefix+scanner.Text()+suffix)
        if err != nil {
            log.Fatalln(err)
        }

        fmt.Println(scanner.Text(), resp.StatusCode)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

}

func fixes(config string) (prefix, suffix string) {
    // open a file
    conf_file, _ := os.Open(config)

    //  Close file
    defer conf_file.Close()

    //NewDecoder creates a*decoder that reads and decodes the JSON object from File. The decoder has its own buffer and may read some JSON data in advance.
    decoder := json.NewDecoder(conf_file)

    conf := configuration{}
    //DECODE Read a JSON coding value from the input stream and save it in the V -Direct value
    err := decoder.Decode(&conf)
    if err != nil {
        fmt.Println("Error:", err)
    }
    prefix = conf.Prefix
    suffix = conf.Suffix
    return
}