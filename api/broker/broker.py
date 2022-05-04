# -*- coding: utf-8 -*-
import zmq
import logging


def start_broker():
    context = zmq.Context()
    router_socket = context.socket(zmq.ROUTER)
    dealer_socket = context.socket(zmq.DEALER)

    router_socket.bind("tcp://*:5556")
    dealer_socket.bind("tcp://*:5557")

    poller = zmq.Poller()
    poller.register(router_socket, zmq.POLLIN)
    poller.register(dealer_socket, zmq.POLLIN)

    logging.info("Start broker")

    while True:
        socks = dict(poller.poll())

        if socks.get(router_socket) == zmq.POLLIN:
            message = router_socket.recv_multipart()
            # logging.debug(f"ROUTER {message}")
            dealer_socket.send_multipart(message)

        if socks.get(dealer_socket) == zmq.POLLIN:
            message = dealer_socket.recv_multipart()
            # logging.debug(f"DEALER {message}")
            router_socket.send_multipart(message)

    router_socket.close()
    dealer_socket.close()
    context.destroy()


if __name__ == "__main__":
    start_broker()