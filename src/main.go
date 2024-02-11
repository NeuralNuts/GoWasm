package main

import "fmt"

func run_http_server() {
    var root_index_file_path string = "../"
    var server_prefix string = "/"
    var port_number string = "8888"

    setup_htttp_server(root_index_file_path, server_prefix, port_number)
}

func main() {
    var wasm_dir_file_path string = "wasm_files/"

    build_wasm_binary(wasm_dir_file_path)
    run_http_server()

    fmt.Println("main() function HIT!")
}
