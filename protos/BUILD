load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")

proto_library(
    name = "battleship_proto",
    srcs = ["battleship.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "battleship_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/anicolaspp/battleship/protos",
    proto = ":battleship_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "protos",
    embed = [":battleship_go_proto"],
    importpath = "github.com/anicolaspp/battleship/protos",
    visibility = ["//visibility:public"],
)
