#!/usr/bin/env python3

import base64

def main():
    ba = bytearray.fromhex("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
    b64 = base64.b64encode(ba)
    print("Text: {}".format(ba.decode("utf8")))
    print("Base64: {}".format(b64.decode("utf8")))

if __name__ == "__main__":
    main()
