def test_single_request_netversion(qubetics_cluster):
    qubetics_cluster.use_websocket()
    eth_ws = qubetics_cluster.w3.provider

    response = eth_ws.make_request("net_version", [])

    # net_version should be 9000
    assert response["result"] == "9000", "got " + response["result"] + ", expected 9000"


# note:
# batch requests still not implemented in web3.py
# todo: follow https://github.com/ethereum/web3.py/issues/832, add tests when complete

# eth_subscribe and eth_unsubscribe support still not implemented in web3.py
# todo: follow https://github.com/ethereum/web3.py/issues/1402, add tests when complete


def test_batch_request_netversion(qubetics):
    return


def test_ws_subscribe_log(qubetics):
    return


def test_ws_subscribe_newheads(qubetics):
    return
