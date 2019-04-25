load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "build_bazel_rules_nodejs",
    urls = ["https://github.com/bazelbuild/rules_nodejs/archive/0.27.12.tar.gz"],
    strip_prefix = "rules_nodejs-0.27.12",
    sha256 = "25dbb063a8a1a2b279d55ba158992ad61eb5266c416c77eb82a7d33b4eac533d",
)

http_archive(
    name = "ecosia_bazel_rules_nodejs_contrib",
    sha256 = "2d3e8b145833e820289a51c5516b33ecf938d54f3d4fc18b3449de25efbb6635",
    strip_prefix = "bazel_rules_nodejs_contrib-bda1e765beda782f59d9e8f3a982d515e1f24634",
    url = "https://github.com/ecosia/bazel_rules_nodejs_contrib/archive/bda1e765beda782f59d9e8f3a982d515e1f24634.tar.gz",
)

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.1/rules_go-0.18.1.tar.gz"],
    sha256 = "77dfd303492f2634de7a660445ee2d3de2960cbd52f97d8c0dffa9362d3ddef9",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)

load("@build_bazel_rules_nodejs//:defs.bzl", "yarn_install")
yarn_install(
    name = "npm",
    package_json = "//:package.json",
    yarn_lock = "//:yarn.lock",
)

load("@npm//:install_bazel_dependencies.bzl", "install_bazel_dependencies")
install_bazel_dependencies()

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()
