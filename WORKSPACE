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
    commit = "b82aac8589e138824736b2a9d466981dbce6b0d4",
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
    commit = "e87afe3c25daf6fa321fed3705938e53672b0a3b",
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
    commit = "53fe9250f3341cda29af5069ae82ff903e0c449a",
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
    commit = "94b4c8b41eba1c82140fee9cca456ad13ae81706",
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
    commit = "d1252f9bfe49590ff7a6381e9f0d4e882b631e28",
    importpath = "github.com/libp2p/go-libp2p-crypto",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_host",
    commit = "e758e0ca4ae1107e480c43a22bf91272a5432105",
    importpath = "github.com/libp2p/go-libp2p-host",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_interface_connmgr",
    commit = "61a030e46d8f6c82500b0992360d7a5c701d367a",
    importpath = "github.com/libp2p/go-libp2p-interface-connmgr",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_loggables",
    commit = "2edffda90e410fab8ca3663511d33b59314d4b07",
    importpath = "github.com/libp2p/go-libp2p-loggables",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_metrics",
    commit = "20c0e3fed14ddf84ac8192038accfd393610ed82",
    importpath = "github.com/libp2p/go-libp2p-metrics",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_net",
    commit = "22c96766db92ab111e506ebcd9cc6511ed32e553",
    importpath = "github.com/libp2p/go-libp2p-net",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_peer",
    commit = "dd9b45c0649b38aebe65f98cb460676b4214a42c",
    importpath = "github.com/libp2p/go-libp2p-peer",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_peerstore",
    commit = "6295e61c9fd2f13ad159c6241be3b371918045e2",
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
    commit = "839f88f8de4d0f8300facdcdf7aa2124d020b2b6",
    importpath = "github.com/libp2p/go-libp2p-swarm",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_transport",
    commit = "e6d021be15cb2bfc73fb24d3b16848bc2825bbf6",
    importpath = "github.com/libp2p/go-libp2p-transport",
)

go_repository(
    name = "com_github_libp2p_go_libp2p",
    commit = "9356373d00ab1aef3e20c8202b682f93799acf78",
    importpath = "github.com/libp2p/go-libp2p",
)

go_repository(
    name = "com_github_libp2p_go_maddr_filter",
    commit = "7f7ca1e79c453741adb1cc10d8892b186952e9e1",
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
    commit = "393301199b4437234bd500b78bdf9555f6428c18",
    importpath = "github.com/multiformats/go-multiaddr",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr_dns",
    commit = "317a9bc842d426da4d878690625babb26d64e73b",
    importpath = "github.com/multiformats/go-multiaddr-dns",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr_net",
    commit = "5b127b903db6848e085c1c2078840d3401f928ef",
    importpath = "github.com/multiformats/go-multiaddr-net",
)

go_repository(
    name = "com_github_multiformats_go_multihash",
    commit = "97cdb562a04c6ef66d8ed40cd62f8fbcddd396d6",
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
    commit = "3b86bcbec8cbb09d205c1492e898ce3d0e81c4d5",
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
    commit = "c10e9556a7bc0e7c942242b606f0acf024ad5d6a",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "9b800f95dbbc54abff0acf7ee32d88ba4e328c89",
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
    commit = "b7192598fc6d96e4aa1d6a565387cd17ab667987",
    importpath = "github.com/libp2p/go-conn-security",
)

go_repository(
    name = "com_github_libp2p_go_conn_security_multistream",
    commit = "578125a681eee24cac960d84827db34541e3f707",
    importpath = "github.com/libp2p/go-conn-security-multistream",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_circuit",
    commit = "16eb677aaa62fd2e5e6d73d66388f8ac38bd6388",
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
    commit = "8f95e95b9fedc69b1367362a14f1ad3b5bd5bd46",
    importpath = "github.com/libp2p/go-libp2p-secio",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_transport_upgrader",
    commit = "8dde02b5e75342c09725bc601cf28e9e98f920c7",
    importpath = "github.com/libp2p/go-libp2p-transport-upgrader",
)

go_repository(
    name = "com_github_libp2p_go_mplex",
    commit = "0ef5fed5ba589e7e8776c274510cfe0d806bac9c",
    importpath = "github.com/libp2p/go-mplex",
)

go_repository(
    name = "com_github_libp2p_go_msgio",
    commit = "031a413e66129d593337a3f5948d9364e7fc6d43",
    importpath = "github.com/libp2p/go-msgio",
)

go_repository(
    name = "com_github_libp2p_go_reuseport",
    commit = "dd0c37d7767bc38280bd9813145b65f8bd560629",
    importpath = "github.com/libp2p/go-reuseport",
)

go_repository(
    name = "com_github_libp2p_go_reuseport_transport",
    commit = "58ea7103ffb4b5eb248c4421e60fdb30e9a56dad",
    importpath = "github.com/libp2p/go-reuseport-transport",
)

go_repository(
    name = "com_github_libp2p_go_sockaddr",
    commit = "a7494d4eefeb607c8bc491cf8850a6e8dbd41cab",
    importpath = "github.com/libp2p/go-sockaddr",
)

go_repository(
    name = "com_github_libp2p_go_tcp_transport",
    commit = "4a25127ad66b71ae4c91f1f42205b2ce679dd926",
    importpath = "github.com/libp2p/go-tcp-transport",
)

go_repository(
    name = "com_github_libp2p_go_ws_transport",
    commit = "246ec4b8bd9a5a539827eca50a6e6d4ce50bb056",
    importpath = "github.com/libp2p/go-ws-transport",
)

go_repository(
    name = "com_github_whyrusleeping_go_smux_multiplex",
    commit = "2b855d4f3a3051b0133f7783bffe06e4b7833d1e",
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
    commit = "ee77252da00fc7ea9bc3ba309d63cff1c13555be",
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
    commit = "56beff422ba25f782821152ce9a2aa4c15712938",
    importpath = "github.com/libp2p/go-libp2p-record",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_routing",
    commit = "c568217bd16dbdb16aaa3064f5d1f2dfa224b589",
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
    commit = "c9e99b39dba9c9d415e033d73e127121c527c8f8",
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
    commit = "6ac78c66a92e82b828ecaa49e86445ba081446d1",
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
