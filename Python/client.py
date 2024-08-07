from __future__ import print_function
import logging
import grpc
from pb import hello_pb2
from pb import hello_pb2_grpc

def run():
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = hello_pb2_grpc.GreeterStub(channel)
        response = stub.SayHello(hello_pb2.HelloRequest(name="israel on python client"))
    print(response.message)

if __name__ == "__main__":
    logging.basicConfig()
    run()
