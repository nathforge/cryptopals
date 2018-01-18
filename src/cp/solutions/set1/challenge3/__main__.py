#!/usr/bin/env python3

import argparse
import binascii
import os
import sys
import time

sys.path.insert(0,os.path.join(os.path.dirname(__file__),"../../../.."))

from cp import hexdump, langscorer, xorguess

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--frame-delay", type=float, default=0.025)
    args = parser.parse_args()

    enc_hex = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

    scorer = langscorer.english
    for state in xorguess.SingleByteKey(scorer, bytearray.fromhex(enc_hex)):
        sys.stdout.write("\033[2J") # Clear screen
        sys.stdout.write("\033[0;0H") # Move cursor to 0,0

        print("   Decrypt: {:.3f} [\033[1m{}\033[m] 0x{:02x}".format(
            state.guess.score,
            hexdump.ascii(state.guess.dec),
            state.guess.xor_byte
        ))
        print()
        for index, guess in enumerate(reversed(state.sorted_guesses[-5:])):
            if index == 0:
                print(
                    "Best guess: "
                    "\033[1m{:.3f}\033[m "
                    "[\033[1m\033[33m{}\033[m] "
                    "\033[1m0x{:02x}\033[m".format(
                        guess.score,
                        hexdump.ascii(guess.dec),
                        guess.xor_byte
                    )
                )
            else:
                print(
                    "            "
                    "{:.3f} "
                    "[{}] "
                    "0x{:02x}".format(
                        guess.score,
                        hexdump.ascii(guess.dec),
                        guess.xor_byte
                    )
                )

        sys.stdout.flush()
        time.sleep(args.frame_delay)

if __name__ == "__main__":
    main()
