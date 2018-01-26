#!/usr/bin/env python

from cp import hexdump, pkcs7

def main():
    for size in range(0, 16 + 1):
        padded = pkcs7.pad(bytearray("X" * size, "utf8"), 16)
        print(len(padded), padded)

if __name__ == "__main__":
    main()
