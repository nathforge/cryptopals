#!/usr/bin/env python3

import binascii

from cp import xor

def main():
    ba = bytearray.fromhex("1c0111001f010100061a024b53535009181c")
    key = bytearray.fromhex("686974207468652062756c6c277320657965")
    enc = xor.repeating_bytes(ba, key)
    print("Output: {}".format(binascii.hexlify(enc).decode("utf8")))

if __name__ == "__main__":
    main()
