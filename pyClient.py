import socket

UDP_IP = "54.152.6.245"
UDP_PORT = 8080

print("UDP target IP:", UDP_IP)
print("UDP target port:", UDP_PORT)

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM) # UDP


for i in range(100):
    MESSAGE = str(i)
    sock.sendto(bytes(MESSAGE, "utf-8"), (UDP_IP, UDP_PORT))
