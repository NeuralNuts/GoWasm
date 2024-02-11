package main

import "fmt"

func run_http_server() {
    var root_index_file_path string = "../"
    var server_prefix string = "/"
    var port_number string = "8888"

    setup_htttp_server(root_index_file_path, server_prefix, port_number)
}

func main() {
    //run_http_server()

    build_wasm_binary()

    fmt.Println("main() function HIT!")
}
