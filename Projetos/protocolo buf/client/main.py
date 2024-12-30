#python -m grpc_tools.protoc -I=./proto/ --python_out=. --grpc_python_out=. proto/*.proto
import school_pb2
import socket

client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client.connect(('localhost', 4000))
data = client.recv(1024)
professor = school_pb2.Professor()
professor.ParseFromString(data)
print(f"Recebido: {professor}")  # Isso deve 
client.close()
