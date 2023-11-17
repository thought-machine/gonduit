go_library(
    name = "util",
    srcs = [
        "strings.go",
        "time.go",
    ],
    visibility = [
        "//:all",
        "//entities:all",
        "//responses:all",
    ],
)

go_test(
    name = "util_test",
    srcs = [
        "strings_test.go",
        "time_test.go",
    ],
    deps = [
        ":util",
        "//third_party/go:testify",
    ],
)
