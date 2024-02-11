package main

import (
    "fmt"
    "io/ioutil"
    "os/exec"
)

func run_wasm_build_command(wasm_file_names string) {
    shell_type := "bash"

    wasm_build_cmd := "GOARCH=wasm GOOS=js go build -o add.wasm " + wasm_file_names
    execute_wasm_build_cmd := exec.Command(shell_type, "-c", wasm_build_cmd)

    stdout, build_error := execute_wasm_build_cmd.Output()

    if build_error != nil {
        fmt.Println(build_error)
        return
    } else {
        fmt.Println(string(stdout))
    }
}

func build_wasm_binary(wasm_dir_file_path string) {
    wasm_dir, read_error := ioutil.ReadDir(wasm_dir_file_path)

    if read_error != nil {
        fmt.Println(read_error)
        return 
    } else {
        for _, wasm_files := range wasm_dir {
            fmt.Println(wasm_files.Name())

            run_wasm_build_command(wasm_dir_file_path + wasm_files.Name())
        }
    }
}
