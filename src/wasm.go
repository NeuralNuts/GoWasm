package main

import (
    "fmt"
    "io/ioutil"
)

func build_wasm_binary() {
    wasm_dir, read_error := ioutil.ReadDir("wasm_files")

    if read_error != nil {
        fmt.Errorf("Cannot read null dir: ", read_error)
        return 
    } else {
        for _, wasm_files := range wasm_dir {
            fmt.Println(wasm_files.Name())
        }
    }
}
