#!/usr/bin/env python3

import binascii

from cp import xor

def main():
    inp = bytearray("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal", "utf8")
    enc = xor.repeating_bytes(inp, bytearray("ICE", "utf8"))
    out = binascii.hexlify(enc)
    print("Output: {}".format(out.decode("utf8")))

if __name__ == "__main__":
    main()
