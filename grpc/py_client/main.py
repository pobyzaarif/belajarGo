from __future__ import print_function
import example_pb2
import example_pb2_grpc
import logging
import grpc
import sys


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('localhost:4040') as channel:
        stub = example_pb2_grpc.CalculateStub(channel)
        numb1, numb2 = 0, 0
        if len(sys.argv) > 2:
            try:
                numb1 = int(sys.argv[1])
                numb2 = int(sys.argv[2])
            except ValueError:
                print('invalid number: {0} or {1}'.format(sys.argv[1], sys.argv[2]))
                return
        response = stub.Add(example_pb2.Request(a=numb1, b=numb2))
        print(response.result)
        response = stub.Multiply(example_pb2.Request(a=numb1, b=numb2))
        print(response.result)

if __name__ == '__main__':
    logging.basicConfig()
    run() 