import grpc
import school_service_pb2
import school_service_pb2_grpc

def run():
    channel = grpc.insecure_channel('localhost:4000')
    stub = school_service_pb2_grpc.ClassServiceStub(channel)

    input = school_service_pb2.InputListAllStudents(code="XYZ123")
    students = stub.ListAllStudents(input)
    for student in students:
        print(f"Received: {student}")

if __name__ == '__main__':
    run()