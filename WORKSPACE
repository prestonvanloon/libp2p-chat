load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "29d109605e0d6f9c892584f07275b8c9260803bf0c6fcb7de2623b2bedc910bd",
    strip_prefix = "rules_docker-0.5.1",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.5.1.tar.gz"],
)

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.16.1/rules_go-0.16.1.tar.gz"],
    sha256 = "f5127a8f911468cd0b2d7a141f17253db81177523e4429796e14d429f5444f5f",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.15.0/bazel-gazelle-0.15.0.tar.gz"],
    sha256 = "6e875ab4b6bf64a38c352887760f21203ab054676d9c1b274963907e0768740d",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

go_repository(
    name = "org_golang_x_text",
    commit = "6f44c5a2ea40ee3593d98cdcc905cc1fdaa660e2",
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
    commit = "fe236c4745a2536facac62dc99ec34c56d833d95",
    importpath = "github.com/libp2p/go-libp2p-nat",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    commit = "be550b025b433cdfa2f11efb962afa2ea3c4d967",
    importpath = "github.com/opentracing/opentracing-go",
)

go_repository(
    name = "com_github_btcsuite_btcd",
    commit = "67e573d211ace594f1366b4ce9d39726c4b19bd0",
    importpath = "github.com/btcsuite/btcd",
)

go_repository(
    name = "com_github_coreos_go_semver",
    commit = "e214231b295a8ea9479f11b70b35d5acf3556d9b",
    importpath = "github.com/coreos/go-semver",
)

go_repository(
    name = "com_github_fd_go_nat",
    commit = "bad65a492f32121a87197f4a085905c35e2a367e",
    importpath = "github.com/fd/go-nat",
)

go_repository(
    name = "com_github_gogo_protobuf",
    commit = "211a54c57504366b3bf4b490fb3c0a5f5b8d290b",
    importpath = "github.com/gogo/protobuf",
)

go_repository(
    name = "com_github_google_uuid",
    commit = "9b3b1e0f5f99ae461456d768e7d301a7acdaa2d8",
    importpath = "github.com/google/uuid",
)

go_repository(
    name = "com_github_gxed_hashland",
    commit = "d9f6b97f8db22dd1e090fd0bbbe98f09cc7dd0a8",
    importpath = "github.com/gxed/hashland",
)

go_repository(
    name = "com_github_huin_goupnp",
    commit = "656e61dfadd241c7cbdd22a023fa81ecb6860ea8",
    importpath = "github.com/huin/goupnp",
)

go_repository(
    name = "com_github_ipfs_go_log",
    commit = "de9a213953d6ec0e053b56e9d79800565c3fc9ca",
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
    commit = "6e2ee01c0ae21694bcc33aaa7d6ce214279134f6",
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
    commit = "3120e9f9526fe05f2d3905961a5e0701b85579d9",
    importpath = "github.com/libp2p/go-libp2p-crypto",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_host",
    commit = "598b2d31a677b46c40855423f88fd10c122f0309",
    importpath = "github.com/libp2p/go-libp2p-host",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_interface_connmgr",
    commit = "eb657584a141a641098794301c9ef3fc9f35561d",
    importpath = "github.com/libp2p/go-libp2p-interface-connmgr",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_loggables",
    commit = "024fce320cf7e720445c400721c6fa6d68c9aaa7",
    importpath = "github.com/libp2p/go-libp2p-loggables",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_metrics",
    commit = "2d5733beaa2a9fdd05ef696d7a734aa61549fb2a",
    importpath = "github.com/libp2p/go-libp2p-metrics",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_net",
    commit = "abc1a0a5489e49274b8115ddda5661a8a8270818",
    importpath = "github.com/libp2p/go-libp2p-net",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_peer",
    commit = "d3df4bca884d7a9c2d350c8120240db3c2b0f2ee",
    importpath = "github.com/libp2p/go-libp2p-peer",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_peerstore",
    commit = "545f62f160f44613034553f69024f91cef4d0de9",
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
    commit = "3d9a2ee0562310a6e6973203348ce3d5eea72c71",
    importpath = "github.com/libp2p/go-libp2p-swarm",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_transport",
    commit = "2f042b8302ef0c73d8722da9e283c77030c78827",
    importpath = "github.com/libp2p/go-libp2p-transport",
)

go_repository(
    name = "com_github_libp2p_go_libp2p",
    commit = "3e2dc09212ad570c9173376ca9e3c54ae8a7d394",
    importpath = "github.com/libp2p/go-libp2p",
)

go_repository(
    name = "com_github_libp2p_go_maddr_filter",
    commit = "9e2c18e714b119399bf9ba45cd26805321bf22b5",
    importpath = "github.com/libp2p/go-maddr-filter",
)

go_repository(
    name = "com_github_libp2p_go_stream_muxer",
    commit = "9c6bd93eecbbab56630bb2688bb435d9fd1dfeb1",
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
    commit = "51976451ce1942acbb55707a983ed232fa027110",
    importpath = "github.com/minio/sha256-simd",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr",
    commit = "ec8630b6b7436b5d7f6c1c2366d3d7214d1b29e2",
    importpath = "github.com/multiformats/go-multiaddr",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr_dns",
    commit = "8ad4cb43a93f543bd1449006f12dd8aab1149a42",
    importpath = "github.com/multiformats/go-multiaddr-dns",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr_net",
    commit = "f0af4033635f1241179700537dacdc04f2803df8",
    importpath = "github.com/multiformats/go-multiaddr-net",
)

go_repository(
    name = "com_github_multiformats_go_multihash",
    commit = "a91e75d03bf4dba801af7b159c8b2aa7b5f47ea8",
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
    commit = "c75a64cef2f64e7d538e6d43a8c58449ba2ab735",
    importpath = "github.com/whyrusleeping/mafmt",
)

go_repository(
    name = "com_github_whyrusleeping_multiaddr_filter",
    commit = "e903e4adabd70b78bc9293b6ee4f359afb3f9f59",
    importpath = "github.com/whyrusleeping/multiaddr-filter",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "4d3f4d9ffa16a13f451c3b2999e9c49e9750bf06",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_net",
    commit = "10aee181995363b41f712a55844a0dd52ea04646",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "7155702f2d47d94b134229da97195d0130cab001",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "com_github_gorilla_websocket",
    commit = "483fb8d7c32fcb4b5636cd293a92e3935932e2f4",
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
    commit = "058210c5a0d042677367d923eb8a6dc072a15f7f",
    importpath = "github.com/libp2p/go-buffer-pool",
)

go_repository(
    name = "com_github_libp2p_go_conn_security",
    commit = "401ec034f8aaf5782b0dec08c56f1d17b9b1ed34",
    importpath = "github.com/libp2p/go-conn-security",
)

go_repository(
    name = "com_github_libp2p_go_conn_security_multistream",
    commit = "cc89dcc3eedb3115ac0bc5284258eed12f7b990c",
    importpath = "github.com/libp2p/go-conn-security-multistream",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_circuit",
    commit = "7983c3631e4532913a7780cb88cb5217cdc00f95",
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
    commit = "6af448ede395fc4a48467f78a0f67a40cf4071ee",
    importpath = "github.com/libp2p/go-libp2p-secio",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_transport_upgrader",
    commit = "e78edd6288169a5c4a83aef78dd1705dd17892c8",
    importpath = "github.com/libp2p/go-libp2p-transport-upgrader",
)

go_repository(
    name = "com_github_libp2p_go_mplex",
    commit = "f6e0e0f222d6bcd9638a5ca0ccce45fa8ce7a274",
    importpath = "github.com/libp2p/go-mplex",
)

go_repository(
    name = "com_github_libp2p_go_msgio",
    commit = "031a413e66129d593337a3f5948d9364e7fc6d43",
    importpath = "github.com/libp2p/go-msgio",
)

go_repository(
    name = "com_github_libp2p_go_reuseport",
    commit = "8cfd5f2973c8e2476813120b9a516d9a82eb7c7a",
    importpath = "github.com/libp2p/go-reuseport",
)

go_repository(
    name = "com_github_libp2p_go_reuseport_transport",
    commit = "a7633860456a881830bde6a33d0722a4c457a61c",
    importpath = "github.com/libp2p/go-reuseport-transport",
)

go_repository(
    name = "com_github_libp2p_go_sockaddr",
    commit = "a7494d4eefeb607c8bc491cf8850a6e8dbd41cab",
    importpath = "github.com/libp2p/go-sockaddr",
)

go_repository(
    name = "com_github_libp2p_go_tcp_transport",
    commit = "e5f699ed1abe3c4689510b66bc810faba000dd6b",
    importpath = "github.com/libp2p/go-tcp-transport",
)

go_repository(
    name = "com_github_libp2p_go_ws_transport",
    commit = "5babbaedef1ce4f082d4d4fc3b6b27eeabb64cdb",
    importpath = "github.com/libp2p/go-ws-transport",
)

go_repository(
    name = "com_github_whyrusleeping_go_smux_multiplex",
    commit = "c9680d872b8e73a2fc8d9bba3eacabf1e8add80d",
    importpath = "github.com/whyrusleeping/go-smux-multiplex",
)

go_repository(
    name = "com_github_whyrusleeping_go_smux_multistream",
    commit = "e799b10bc6eea2cc5ce18f7b7b4d8e1a0439476d",
    importpath = "github.com/whyrusleeping/go-smux-multistream",
)

go_repository(
    name = "com_github_whyrusleeping_go_smux_yamux",
    commit = "49458276a01f7fbc32ff62c8955fa3e852b8e772",
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
    commit = "b579e0352c7cc727d16d3df1bea6681fc27d8108",
    importpath = "github.com/libp2p/go-libp2p-kad-dht",
)

go_repository(
    name = "com_github_ipfs_go_datastore",
    commit = "0706d00f0907891cb658f82b58c96a195c0cd6f0",
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
    commit = "237ab9e10af172232eedad9e63ea8983c50859b1",
    importpath = "github.com/libp2p/go-libp2p-record",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_routing",
    commit = "76077861eb4ae107b4366acc1a1152b8e15d4300",
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
    commit = "033594dcd6201e8e8628659c8d584bb800d1734c",
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
    commit = "b181991db757142b92d913a2f93ea3ee92288aad",
    importpath = "github.com/libp2p/go-libp2p-kbucket",
)

go_repository(
    name = "com_github_multiformats_go_multibase",
    commit = "bb91b53e5695e699a86654d77d03db7bc7506d12",
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
    commit = "abad2d1bd44235a26707c172eab6bca5bf2dbad3",
    importpath = "github.com/prometheus/client_golang",
)

go_repository(
    name = "com_github_prometheus_common",
    commit = "7e9e6cabbd393fc208072eedef99188d0ce788b6",
    importpath = "github.com/prometheus/common",
)

go_repository(
    name = "com_github_prestonvanloon_libp2p_chat",
    commit = "5d013f36d11e3fc9f4ca326d4bf65e1f5c233bbd",
    importpath = "github.com/prestonvanloon/libp2p-chat",
)

go_repository(
    name = "com_github_prometheus_client_model",
    commit = "5c3871d89910bfb32f5fcab2aa4b9ec68e65a99f",
    importpath = "github.com/prometheus/client_model",
)

go_repository(
    name = "com_github_prometheus_procfs",
    commit = "185b4288413d2a0dd0806f78c90dde719829e5ae",
    importpath = "github.com/prometheus/procfs",
)

go_repository(
    name = "com_github_beorn7_perks",
    commit = "3a771d992973f24aa725d07868b467d1ddfceafb",
    importpath = "github.com/beorn7/perks",
)

go_repository(
    name = "com_github_golang_protobuf",
    commit = "75dceb112b174be156fa9952e66d3e99945572b4",
    importpath = "github.com/golang/protobuf",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    commit = "c12348ce28de40eed0136aa2b644d0ee0650e56c",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_discovery",
    commit = "a35c2b4e8fa0b65864648bfd260cf23acfe9b0b2",
    importpath = "github.com/libp2p/go-libp2p-discovery",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_autonat",
    commit = "e37e937c770fecddf7860a2a0f24ada9dbbd003d",
    importpath = "github.com/libp2p/go-libp2p-autonat",
)

go_repository(
    name = "com_github_pkg_errors",
    commit = "059132a15dd08d6704c67711dae0cf35ab991756",
    importpath = "github.com/pkg/errors",
)

go_repository(
    name = "com_github_uber_jaeger_client_go",
    commit = "713fb4b2e160d92f92dbeb95369bd73ddbd81d11",
    importpath = "github.com/uber/jaeger-client-go",
)

go_repository(
    name = "com_github_uber_jaeger_lib",
    commit = "1fc5c315e03c871d98a3df762234414d185115d7",
    importpath = "github.com/uber/jaeger-lib",
)
