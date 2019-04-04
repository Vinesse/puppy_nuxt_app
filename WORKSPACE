load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
http_archive(
    name = "build_bazel_rules_nodejs",
    sha256 = "1fbea5f33a8cbf8fd798f4cf7973a9b09e53bda872c75b5665a92a6d28fbbb13",
    urls = ["https://github.com/bazelbuild/rules_nodejs/releases/download/0.27.10/rules_nodejs-0.27.10.tar.gz"],
)

load("@build_bazel_rules_nodejs//:defs.bzl", "yarn_install")
yarn_install(
    name = "npm",
    package_json = "//:package.json",
    yarn_lock = "//:yarn.lock",
)

load("@npm//:install_bazel_dependencies.bzl", "install_bazel_dependencies")
install_bazel_dependencies()