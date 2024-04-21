load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

package(
    default_visibility = ["//visibility:public"],
)

go_binary(
    name = "main",
    embed = [":battleship_lib"],
)

go_library(
    name = "battleship_lib",
    srcs = ["main.go"],
    importpath = "github.com/anicolaspp/battleship",
)
