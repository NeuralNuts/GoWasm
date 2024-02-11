package main

import (
    "fmt"
    "net/http"
)


const HEADER string = "\033[95m"
const OKBLUE string = "\033[94m"
const OKCYAN string = "\033[96m"
const OKGREEN string = "\033[92m"
const WARNING string = "\033[93m"
const FAIL string = "\033[91m"
const ENDC string = "\033[0m"
const BOLD string = "\033[1m"
const UNDERLINE string = "\033[4m"

func add_wasm_test(a int, b int) int {
    return a + b
}

func setup_htttp_server(file_path string, server_prefix string, port_number string) {
    var base_local_host_url string = "http://localhost:"

    fmt.Println("SERVING AT: " + "[" + OKGREEN + base_local_host_url +
        port_number + "/" + file_path + ENDC + "]")

    http.Handle(server_prefix, http.FileServer(http.Dir(file_path)))
    http.ListenAndServe(":" + port_number, nil)
}
