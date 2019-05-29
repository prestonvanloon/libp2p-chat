load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Download the rules_docker repository at release v0.7.0
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "aed1c249d4ec8f703edddf35cbe9dfaca0b5f5ea6e4cd9e83e99f3b0d1136c3d",
    strip_prefix = "rules_docker-0.7.0",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.7.0.tar.gz"],
)

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.5/rules_go-0.18.5.tar.gz"],
    sha256 = "a82a352bffae6bee4e95f68a8d80a70e87f42c4741e6a448bec11998fcc82329",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "org_golang_x_text",
    commit = "6c92c7dc7f53607809182301b96e4cc1975143f1",
    importpath = "golang.org/x/text",
)

go_repository(
    name = "com_github_mr_tron_base58",
    commit = "d724c80ecac7b49e4e562d58b2b4f4ee4ed8c312",
    importpath = "github.com/mr-tron/base58",
)

go_repository(
    name = "com_github_mr_tron_base58",
    commit = "d724c80ecac7b49e4e562d58b2b4f4ee4ed8c312",
    importpath = "github.com/mr-tron/base58",
)

go_repository(
    name = "com_github_whyrusleeping_go_logging",
    commit = "0457bb6b88fc1973573aaf6b5145d8d3ae972390",
    importpath = "github.com/whyrusleeping/go-logging",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_nat",
    commit = "a304452f6e87cb80430babce35f74ded88be8853",
    importpath = "github.com/libp2p/go-libp2p-nat",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    commit = "25a84ff92183e2f8ac018ba1db54f8a07b3c0e04",
    importpath = "github.com/opentracing/opentracing-go",
)

go_repository(
    name = "com_github_btcsuite_btcd",
    commit = "306aecffea325e97f513b3ff0cf7895a5310651d",
    importpath = "github.com/btcsuite/btcd",
)

go_repository(
    name = "com_github_coreos_go_semver",
    commit = "e214231b295a8ea9479f11b70b35d5acf3556d9b",
    importpath = "github.com/coreos/go-semver",
)

go_repository(
    name = "com_github_fd_go_nat",
    commit = "e3ba0d89e7d9f0a458bf08baae8db007eb7d242d",
    importpath = "github.com/fd/go-nat",
)

go_repository(
    name = "com_github_gogo_protobuf",
    commit = "ba06b47c162d49f2af050fb4c75bcbc86a159d5c",
    importpath = "github.com/gogo/protobuf",
)

go_repository(
    name = "com_github_google_uuid",
    commit = "9b3b1e0f5f99ae461456d768e7d301a7acdaa2d8",
    importpath = "github.com/google/uuid",
)

go_repository(
    name = "com_github_gxed_hashland",
    commit = "a72cc0875a1e95edd309d3134bc7c11bf2d7360b",
    importpath = "github.com/gxed/hashland",
)

go_repository(
    name = "com_github_huin_goupnp",
    commit = "656e61dfadd241c7cbdd22a023fa81ecb6860ea8",
    importpath = "github.com/huin/goupnp",
)

go_repository(
    name = "com_github_ipfs_go_log",
    commit = "8924f37936b0db0ffa60a1b0aba3ae973ed8c907",
    importpath = "github.com/ipfs/go-log",
)

go_repository(
    name = "com_github_jackpal_gateway",
    commit = "cbcf4e3f3baee7952fc386c8b2534af4d267c875",
    importpath = "github.com/jackpal/gateway",
)

go_repository(
    name = "com_github_jackpal_go_nat_pmp",
    commit = "d89d09f6f3329bc3c2479aa3cafd76a5aa93a35c",
    importpath = "github.com/jackpal/go-nat-pmp",
)

go_repository(
    name = "com_github_jbenet_goprocess",
    commit = "b497e2f366b8624394fb2e89c10ab607bebdde0b",
    importpath = "github.com/jbenet/goprocess",
)

go_repository(
    name = "com_github_libp2p_go_addr_util",
    commit = "73d4c93d8ab2cfebc0da96b04438693286340daf",
    importpath = "github.com/libp2p/go-addr-util",
)

go_repository(
    name = "com_github_libp2p_go_flow_metrics",
    commit = "7e5a55af485341567f98d6847a373eb5ddcdcd43",
    importpath = "github.com/libp2p/go-flow-metrics",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_crypto",
    build_file_proto_mode = "disable_global",
    commit = "e333f2201582e62b9f86d559aa9daa4ae3f41bec",
    importpath = "github.com/libp2p/go-libp2p-crypto",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_host",
    commit = "1a71c422ef2810530b32455426f1f92d81ece84c",
    importpath = "github.com/libp2p/go-libp2p-host",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_interface_connmgr",
    commit = "c8552ddb959e1ba3b1ba14c976335f968f46ebf9",
    importpath = "github.com/libp2p/go-libp2p-interface-connmgr",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_loggables",
    commit = "5b80b7ea4ee3f40a73cf09ea971038a1b597b6bc",
    importpath = "github.com/libp2p/go-libp2p-loggables",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_metrics",
    commit = "eb0033e81c5ef7b2609df5875af64a8a1bd54b74",
    importpath = "github.com/libp2p/go-libp2p-metrics",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_net",
    commit = "029d28261e49417116be58019c3edca41ba78847",
    importpath = "github.com/libp2p/go-libp2p-net",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_peer",
    commit = "9a193a66da547245154976338737b11eac2d583e",
    importpath = "github.com/libp2p/go-libp2p-peer",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_peerstore",
    commit = "985bc5953b50e6817231f6b821778962c3eaa4b9",
    importpath = "github.com/libp2p/go-libp2p-peerstore",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    commit = "efa589957cd060542a26d2dd7832fd6a6c6c3ade",
    importpath = "github.com/mattn/go-colorable",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_protocol",
    commit = "b29f3d97e3a2fb8b29c5d04290e6cb5c5018004b",
    importpath = "github.com/libp2p/go-libp2p-protocol",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_swarm",
    commit = "7a03ca822298635374de4ec9334436e2c858be99",
    importpath = "github.com/libp2p/go-libp2p-swarm",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_transport",
    commit = "d2bc1c17e0285d61493fabc8ca8b573d2ee0665f",
    importpath = "github.com/libp2p/go-libp2p-transport",
)

go_repository(
    name = "com_github_libp2p_go_libp2p",
    commit = "695d0ce801951bcb81dc68664d6a2bd2a8c09e13",
    importpath = "github.com/libp2p/go-libp2p",
)

go_repository(
    name = "com_github_libp2p_go_maddr_filter",
    commit = "666a1351c131174d1d4686d4bb2b2b48c87fd1fa",
    importpath = "github.com/libp2p/go-maddr-filter",
)

go_repository(
    name = "com_github_libp2p_go_stream_muxer",
    commit = "a3f82916c8ad4fb33755ab42eea6e03f8d754839",
    importpath = "github.com/libp2p/go-stream-muxer",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    commit = "3fb116b820352b7f0c281308a4d6250c22d94e27",
    importpath = "github.com/mattn/go-isatty",
)

go_repository(
    name = "com_github_minio_blake2b_simd",
    commit = "3f5f724cb5b182a5c278d6d3d55b40e7f8c2efb4",
    importpath = "github.com/minio/blake2b-simd",
)

go_repository(
    name = "com_github_minio_sha256_simd",
    commit = "2d45a736cd16732fe6a57563cc20d8b035193e58",
    importpath = "github.com/minio/sha256-simd",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr",
    commit = "19d7396df1055360753d180f17d498d20466716e",
    importpath = "github.com/multiformats/go-multiaddr",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr_dns",
    commit = "b3d6340f0777006db9e16e2817745c52fd64a500",
    importpath = "github.com/multiformats/go-multiaddr-dns",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr_net",
    commit = "c75d1cac17a0d84dbf8b2c53c61f0ebf0575183a",
    importpath = "github.com/multiformats/go-multiaddr-net",
)

go_repository(
    name = "com_github_multiformats_go_multihash",
    commit = "6b439b7c6e3c44c112171540500be697ba235235",
    importpath = "github.com/multiformats/go-multihash",
)

go_repository(
    name = "com_github_multiformats_go_multistream",
    commit = "0c61f185f3d6e16bcda416874e7a0fca4696e7e0",
    importpath = "github.com/multiformats/go-multistream",
)

go_repository(
    name = "com_github_spaolacci_murmur3",
    commit = "f09979ecbc725b9e6d41a297405f65e7e8804acc",
    importpath = "github.com/spaolacci/murmur3",
)

go_repository(
    name = "com_github_whyrusleeping_go_notifier",
    commit = "097c5d47330ff6a823f67e3515faa13566a62c6f",
    importpath = "github.com/whyrusleeping/go-notifier",
)

go_repository(
    name = "com_github_whyrusleeping_mafmt",
    commit = "d18651cae1dcc0938f12fd3da4fd2a7ed7c73126",
    importpath = "github.com/whyrusleeping/mafmt",
)

go_repository(
    name = "com_github_whyrusleeping_multiaddr_filter",
    commit = "e903e4adabd70b78bc9293b6ee4f359afb3f9f59",
    importpath = "github.com/whyrusleeping/multiaddr-filter",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "a4c6cb3142f211c99e4bf4cd769535b29a9b616f",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_net",
    commit = "3a22650c66bd7f4fb6d1e8072ffd7b75c8a27898",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "ec7b60b042fd2c54ad5ceaac0c891bd88843f4a2",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "com_github_gorilla_websocket",
    commit = "7c8e298727d149d7c329b4dec7e94e1932ac5c11",
    importpath = "github.com/gorilla/websocket",
)

go_repository(
    name = "com_github_gxed_goendian",
    commit = "0f5c6873267e5abf306ffcdfcfa4bf77517ef4a7",
    importpath = "github.com/gxed/GoEndian",
)

go_repository(
    name = "com_github_gxed_eventfd",
    commit = "80a92cca79a8041496ccc9dd773fcb52a57ec6f9",
    importpath = "github.com/gxed/eventfd",
)

go_repository(
    name = "com_github_jbenet_go_temp_err_catcher",
    commit = "aac704a3f4f27190b4ccc05f303a4931fd1241ff",
    importpath = "github.com/jbenet/go-temp-err-catcher",
)

go_repository(
    name = "com_github_libp2p_go_buffer_pool",
    commit = "a8d831235797607f736331e5f8c702e92550fb7b",
    importpath = "github.com/libp2p/go-buffer-pool",
)

go_repository(
    name = "com_github_libp2p_go_conn_security",
    commit = "3e30d86de3d7ded418e57afa91e9af41aa667d36",
    importpath = "github.com/libp2p/go-conn-security",
)

go_repository(
    name = "com_github_libp2p_go_conn_security_multistream",
    commit = "199bde7312d5a5c3c99b6663a1690a63a6cbd836",
    importpath = "github.com/libp2p/go-conn-security-multistream",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_circuit",
    commit = "646688e9f8abdb9de9e71a837a3b087cab44bdcb",
    importpath = "github.com/libp2p/go-libp2p-circuit",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_interface_pnet",
    commit = "d240acf619f63dfb776821a1d4d28a918f77edd5",
    importpath = "github.com/libp2p/go-libp2p-interface-pnet",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_secio",
    build_file_proto_mode = "disable_global",
    commit = "38f90b017ad167cb9ce84267b826fe9d3c801bad",
    importpath = "github.com/libp2p/go-libp2p-secio",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_transport_upgrader",
    commit = "68cf0192f1d454f0ca409ccea450f220b433d5c6",
    importpath = "github.com/libp2p/go-libp2p-transport-upgrader",
)

go_repository(
    name = "com_github_libp2p_go_mplex",
    commit = "8ac902b6abdf9f65dbb520de1c9bde952da98469",
    importpath = "github.com/libp2p/go-mplex",
)

go_repository(
    name = "com_github_libp2p_go_msgio",
    commit = "f8aaa1f70c8b4d3bff511251db3d8fbac7ce2839",
    importpath = "github.com/libp2p/go-msgio",
)

go_repository(
    name = "com_github_libp2p_go_reuseport",
    commit = "6a75ffd916e5e70a5231da13d63ee5bba1109742",
    importpath = "github.com/libp2p/go-reuseport",
)

go_repository(
    name = "com_github_libp2p_go_reuseport_transport",
    commit = "7c623f705b2b7639b091d045d7ea1f973e24c5c0",
    importpath = "github.com/libp2p/go-reuseport-transport",
)

go_repository(
    name = "com_github_libp2p_go_sockaddr",
    commit = "a7494d4eefeb607c8bc491cf8850a6e8dbd41cab",
    importpath = "github.com/libp2p/go-sockaddr",
)

go_repository(
    name = "com_github_libp2p_go_tcp_transport",
    commit = "ddf3e5c50ef07635f5b011712938f21ec44d6caa",
    importpath = "github.com/libp2p/go-tcp-transport",
)

go_repository(
    name = "com_github_libp2p_go_ws_transport",
    commit = "ebb07cb6e3ab3604a8805c23399106e5f7e03051",
    importpath = "github.com/libp2p/go-ws-transport",
)

go_repository(
    name = "com_github_whyrusleeping_go_smux_multiplex",
    commit = "40e9838863a304ad821775427dda6f965803526f",
    importpath = "github.com/whyrusleeping/go-smux-multiplex",
)

go_repository(
    name = "com_github_whyrusleeping_go_smux_multistream",
    commit = "8e5c10881353c91b4214466d916ad8429508c056",
    importpath = "github.com/whyrusleeping/go-smux-multistream",
)

go_repository(
    name = "com_github_whyrusleeping_go_smux_yamux",
    commit = "28bea8f315d17495c5cad8547b661245486ebeee",
    importpath = "github.com/whyrusleeping/go-smux-yamux",
)

go_repository(
    name = "com_github_whyrusleeping_yamux",
    commit = "5364a42fe4b5efa5967c11c8f9b0f049cac0c4a9",
    importpath = "github.com/whyrusleeping/yamux",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_kad_dht",
    build_file_proto_mode = "disable_global",
    commit = "0e89ed9015a8068b4191cf06dee2873dee7233a2",
    importpath = "github.com/libp2p/go-libp2p-kad-dht",
)

go_repository(
    name = "com_github_ipfs_go_datastore",
    commit = "0578849277985097cdc0559a1466cd389763b8b1",
    importpath = "github.com/ipfs/go-datastore",
)

go_repository(
    name = "com_github_whyrusleeping_base32",
    commit = "c30ac30633ccdabefe87eb12465113f06f1bab75",
    importpath = "github.com/whyrusleeping/base32",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_record",
    build_file_proto_mode = "disable_global",
    commit = "4e8ffc3e248560c97071de01ba2efeadee59bff6",
    importpath = "github.com/libp2p/go-libp2p-record",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_routing",
    commit = "330243f4314893c62db38d05d6752e0a8579ed80",
    importpath = "github.com/libp2p/go-libp2p-routing",
)

go_repository(
    name = "com_github_ipfs_go_todocounter",
    commit = "1e832b829506383050e6eebd12e05ea41a451532",
    importpath = "github.com/ipfs/go-todocounter",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_util",
    commit = "10d786c5ed859afd22223df76a89bf57b24b2ee1",
    importpath = "github.com/ipfs/go-ipfs-util",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    commit = "20f1fb78b0740ba8c3cb143a61e86ba5c8669768",
    importpath = "github.com/hashicorp/golang-lru",
)

go_repository(
    name = "com_github_ipfs_go_cid",
    commit = "29a66d1820a34dcf38d57a7f6c7f707619fbedd3",
    importpath = "github.com/ipfs/go-cid",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_util",
    commit = "10d786c5ed859afd22223df76a89bf57b24b2ee1",
    importpath = "github.com/ipfs/go-ipfs-util",
)

go_repository(
    name = "com_github_jbenet_go_context",
    commit = "d14ea06fba99483203c19d92cfcd13ebe73135f4",
    importpath = "github.com/jbenet/go-context",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_kbucket",
    commit = "f98f2bd87bdf1b089e59aaa1e6622a61c0d71f81",
    importpath = "github.com/libp2p/go-libp2p-kbucket",
)

go_repository(
    name = "com_github_multiformats_go_multibase",
    commit = "f25b77813c0a9aac7c5169b9978277865da5e7c8",
    importpath = "github.com/multiformats/go-multibase",
)

go_repository(
    name = "com_github_whyrusleeping_go_keyspace",
    commit = "5b898ac5add1da7178a4a98e69cb7b9205c085ee",
    importpath = "github.com/whyrusleeping/go-keyspace",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_util",
    commit = "10d786c5ed859afd22223df76a89bf57b24b2ee1",
    importpath = "github.com/ipfs/go-ipfs-util",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    commit = "4c99dd66303a54cbf8559dd6110d5c30b1819e4c",
    importpath = "github.com/prometheus/client_golang",
)

go_repository(
    name = "com_github_prometheus_common",
    commit = "7a3416fd1f41341905f138746c51a5092e7ddf7a",
    importpath = "github.com/prometheus/common",
)

go_repository(
    name = "com_github_prestonvanloon_libp2p_chat",
    commit = "e2972ee2662e07d3dd1ea7d5ddb566a0766bde1b",
    importpath = "github.com/prestonvanloon/libp2p-chat",
)

go_repository(
    name = "com_github_prometheus_client_model",
    commit = "fd36f4220a901265f90734c3183c5f0c91daa0b8",
    importpath = "github.com/prometheus/client_model",
)

go_repository(
    name = "com_github_prometheus_procfs",
    commit = "e4d4a2206da023361ed100d85c5f2cf9c8364e9f",
    importpath = "github.com/prometheus/procfs",
)

go_repository(
    name = "com_github_beorn7_perks",
    commit = "3a771d992973f24aa725d07868b467d1ddfceafb",
    importpath = "github.com/beorn7/perks",
)

go_repository(
    name = "com_github_golang_protobuf",
    commit = "c823c79ea1570fb5ff454033735a8e68575d1d0f",
    importpath = "github.com/golang/protobuf",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    commit = "c182affec369e30f25d3eb8cd8a478dee585ae7d",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_discovery",
    commit = "5e0d40c7c8803e06c9eebe6dd3a0525a1d774a82",
    importpath = "github.com/libp2p/go-libp2p-discovery",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_autonat",
    commit = "842b9c4919f5be0df3baee15b0f97e493f6d3fb2",
    importpath = "github.com/libp2p/go-libp2p-autonat",
)

go_repository(
    name = "com_github_pkg_errors",
    commit = "856c240a51a2bf8fb8269ea7f3f9b046aadde36e",
    importpath = "github.com/pkg/errors",
)

go_repository(
    name = "com_github_uber_jaeger_client_go",
    commit = "64f57863bf63d3842dbe79cdc793d57baaff9ab5",
    importpath = "github.com/uber/jaeger-client-go",
)

go_repository(
    name = "com_github_uber_jaeger_lib",
    commit = "d036253de8f5b698150d81b922486f1e8e7628ec",
    importpath = "github.com/uber/jaeger-lib",
)
