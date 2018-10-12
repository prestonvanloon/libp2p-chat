load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.15.4/rules_go-0.15.4.tar.gz"],
    sha256 = "7519e9e1c716ae3c05bd2d984a42c3b02e690c5df728dc0a84b23f90c355c5a1",
)
http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.14.0/bazel-gazelle-0.14.0.tar.gz"],
    sha256 = "c0a5739d12c6d05b6c1ad56f2200cb0b57c5a70e03ebd2f7b87ce88cabf09c7b",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
gazelle_dependencies()

go_repository(
    name = "org_golang_x_text",
    commit = "4d1c5fb19474adfe9562c9847ba425e7da817e81",
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
    commit = "fcc8db1a9963a3d489abbc632f2c0f404e4ff607",
    importpath = "github.com/libp2p/go-libp2p-nat",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    commit = "6aa6febac7b98f836100ecaea478c04f30b6dbd0",
    importpath = "github.com/opentracing/opentracing-go",
)

go_repository(
    name = "com_github_btcsuite_btcd",
    commit = "2a560b2036bee5e3679ec2133eb6520b2f195213",
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
    commit = "fd322a3c49630fe6d05737e2b7d9426e6680e28d",
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
    commit = "1395d1447324cbea88d249fbfcfd70ea878fdfca",
    importpath = "github.com/huin/goupnp",
)

go_repository(
    name = "com_github_ipfs_go_log",
    commit = "5e6883f800a4ca39748c54d467126a22f92f3523",
    importpath = "github.com/ipfs/go-log",
)

go_repository(
    name = "com_github_jackpal_gateway",
    commit = "cbcf4e3f3baee7952fc386c8b2534af4d267c875",
    importpath = "github.com/jackpal/gateway",
)

go_repository(
    name = "com_github_jackpal_go_nat_pmp",
    commit = "28a68d0c24adce1da43f8df6a57340909ecd7fdd",
    importpath = "github.com/jackpal/go-nat-pmp",
)

go_repository(
    name = "com_github_jbenet_goprocess",
    commit = "b497e2f366b8624394fb2e89c10ab607bebdde0b",
    importpath = "github.com/jbenet/goprocess",
)

go_repository(
    name = "com_github_libp2p_go_addr_util",
    commit = "2dc53609d9aa745f7bf5755bef6f54fb9ff77cd6",
    importpath = "github.com/libp2p/go-addr-util",
)

go_repository(
    name = "com_github_libp2p_go_flow_metrics",
    commit = "7e5a55af485341567f98d6847a373eb5ddcdcd43",
    importpath = "github.com/libp2p/go-flow-metrics",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_crypto",
    commit = "7240b40a3ddc47c4d17c15baabcbe45e5219171b",
    importpath = "github.com/libp2p/go-libp2p-crypto",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_host",
    commit = "c178cbca4e134d6e7d9875b4cb5ea24ed91d6794",
    importpath = "github.com/libp2p/go-libp2p-host",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_interface_connmgr",
    commit = "482a51b59dc90355e35ca7a3944d946e5f37f23d",
    importpath = "github.com/libp2p/go-libp2p-interface-connmgr",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_loggables",
    commit = "fc4c4dc5d90bcd55e0af5b5f744005519e070dc3",
    importpath = "github.com/libp2p/go-libp2p-loggables",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_metrics",
    commit = "a10ff6e75dae3c868023867e8caa534a04bdc624",
    importpath = "github.com/libp2p/go-libp2p-metrics",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_net",
    commit = "313c9d5e2ce669728e2cc95fc8f8a3b82bdb3156",
    importpath = "github.com/libp2p/go-libp2p-net",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_peer",
    commit = "993d742bc29dcf4894b7730ba610fd78900be76c",
    importpath = "github.com/libp2p/go-libp2p-peer",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_peerstore",
    commit = "b87bbd4e218a39cd1a3d2662c97895b0af6ae985",
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
    commit = "aebf02bb94e183b2dad9226d21c3b1879dc81ed0",
    importpath = "github.com/libp2p/go-libp2p-swarm",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_transport",
    commit = "419f57b92e2d5c5b5d4181e8d037fea08b2e09d1",
    importpath = "github.com/libp2p/go-libp2p-transport",
)

go_repository(
    name = "com_github_libp2p_go_libp2p",
    commit = "9750574e2eb77892de61e6cf494a35ebabf30060",
    importpath = "github.com/libp2p/go-libp2p",
)

go_repository(
    name = "com_github_libp2p_go_maddr_filter",
    commit = "57fd7e2ed649ba28b4f2c7bcab3a606e7cc4b12c",
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
    commit = "fe1c46f8be5af4aff4db286e08839295bd922efb",
    importpath = "github.com/multiformats/go-multiaddr",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr_dns",
    commit = "78f39e8892d4f8c5c9f18679cc06d0b40ecab8cf",
    importpath = "github.com/multiformats/go-multiaddr-dns",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr_net",
    commit = "1cb9a0e8a6de3c8a10f6cee60d01d793603c4f7e",
    importpath = "github.com/multiformats/go-multiaddr-net",
)

go_repository(
    name = "com_github_multiformats_go_multihash",
    commit = "97cdb562a04c6ef66d8ed40cd62f8fbcddd396d6",
    importpath = "github.com/multiformats/go-multihash",
)

go_repository(
    name = "com_github_multiformats_go_multistream",
    commit = "0e509f6b999d3be262c0702d25fe868c77de1565",
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
    commit = "1dc32401ee9fdd3f6cdb3405ec984d5dae877b2a",
    importpath = "github.com/whyrusleeping/mafmt",
)

go_repository(
    name = "com_github_whyrusleeping_multiaddr_filter",
    commit = "e903e4adabd70b78bc9293b6ee4f359afb3f9f59",
    importpath = "github.com/whyrusleeping/multiaddr-filter",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "7c1a557ab941a71c619514f229f0b27ccb0c27cf",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_net",
    commit = "49bb7cea24b1df9410e1712aa6433dae904ff66a",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "fa43e7bc11baaae89f3f902b2b4d832b68234844",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "com_github_gorilla_websocket",
    commit = "a51a35ae3232254685f26a0b6d995ca0e81e2248",
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
    commit = "d9a2f52687b63f94b2bb087bcdf22e5454eb09d8",
    importpath = "github.com/libp2p/go-conn-security",
)

go_repository(
    name = "com_github_libp2p_go_conn_security_multistream",
    commit = "42e375808ebbc786e791f019d6a91097d90274a4",
    importpath = "github.com/libp2p/go-conn-security-multistream",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_circuit",
    commit = "543688aca031b4a66c2744cd6f24b1c032a9d789",
    importpath = "github.com/libp2p/go-libp2p-circuit",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_interface_pnet",
    commit = "d240acf619f63dfb776821a1d4d28a918f77edd5",
    importpath = "github.com/libp2p/go-libp2p-interface-pnet",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_secio",
    commit = "be39cabff59e06419cadbdb205abb918d555b5fb",
    importpath = "github.com/libp2p/go-libp2p-secio",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_transport_upgrader",
    commit = "939dfe8cfd16ce4d3568165c1a029703b630550c",
    importpath = "github.com/libp2p/go-libp2p-transport-upgrader",
)

go_repository(
    name = "com_github_libp2p_go_mplex",
    commit = "1386e7e2261644453d5aecc59c364fcb25ce7a1e",
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
    commit = "5cdb097c8035e75fc59d12f22509aeb700a272d0",
    importpath = "github.com/libp2p/go-reuseport-transport",
)

go_repository(
    name = "com_github_libp2p_go_sockaddr",
    commit = "a7494d4eefeb607c8bc491cf8850a6e8dbd41cab",
    importpath = "github.com/libp2p/go-sockaddr",
)

go_repository(
    name = "com_github_libp2p_go_tcp_transport",
    commit = "31233dd1174fca68461cb3032eeeed81ba7337b3",
    importpath = "github.com/libp2p/go-tcp-transport",
)

go_repository(
    name = "com_github_libp2p_go_ws_transport",
    commit = "b2cd8129d1037bcb64047839ecc490c4a3c8be28",
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
    commit = "47e72272c5734e889733a2c97c16ac6e43e258fb",
    importpath = "github.com/whyrusleeping/go-smux-yamux",
)

go_repository(
    name = "com_github_whyrusleeping_yamux",
    commit = "98c0f1efe02b69cd8d8d45452d83663fa25c81e0",
    importpath = "github.com/whyrusleeping/yamux",
)
