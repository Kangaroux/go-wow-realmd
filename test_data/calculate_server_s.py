"""
Usage:
$ python3 calculate_server_s.py > calculate_server_s.csv
"""

from gen import random_string

for _ in range(100):
    row = [
        random_string("hex", 64).upper(), # client public key (32 bytes, little endian)
        random_string("hex", 64).upper(), # server private key (32 bytes, little endian)
        random_string("hex", 64).upper(), # verifier (32 bytes, little endian)
        random_string("hex", 64).upper(), # u (32 bytes, little endian)
        "REPLACE_ME_IN_CSV", # expected value (32 bytes, little endian)
    ]
    print(",".join(row))
