go_library(
    name = "core",
    srcs = [
        "call.go",
        "error.go",
        "http.go",
        "request.go",
    ],
    visibility = ["//:all"],
    deps = [
        "//requests",
        "//third_party/go:typed",
    ],
)

go_test(
    name = "core_test",
    srcs = [
        "call_test.go",
        "error_test.go",
        "http_test.go",
        "request_test.go",
    ],
    deps = [
        ":core",
        "//requests",
        "//responses",
        "//third_party/go:gin",
        "//third_party/go:testify",
    ],
)
