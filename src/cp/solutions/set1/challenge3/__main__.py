#!/usr/bin/env python3

import argparse
import binascii
import sys
import time

from cp import ansi, hexdump, langscorer, xorguess

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--frame-delay", type=float, default=0.025)
    args = parser.parse_args()

    enc_hex = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

    scorer = langscorer.english
    for state in xorguess.SingleByteKey(scorer, bytearray.fromhex(enc_hex)):
        sys.stdout.write(ansi.clear_screen)
        sys.stdout.write(ansi.cursor(0, 0))

        print(
            "   Decrypt: "
            "{score:.3f} "
            "[\033[1m{dec}\033[m] "
            "0x{xor_byte:02x}".format(
                score=state.guess.score,
                dec=hexdump.ascii(state.guess.dec),
                xor_byte=state.guess.xor_byte
            )
        )
        print()
        for index, guess in enumerate(reversed(state.sorted_guesses[-5:])):
            if index == 0:
                fmt = (
                    "Best guess: "
                    "\033[1m{score:.3f}\033[m "
                    "[\033[1m\033[33m{dec}\033[m] "
                    "\033[1m0x{xor_byte:02x}\033[m"
                )
            else:
                fmt = (
                    "            "
                    "{score:.3f} "
                    "[{dec}] "
                    "0x{xor_byte:02x}"
                )
            
            print(fmt.format(
                score=guess.score,
                dec=hexdump.ascii(guess.dec),
                xor_byte=guess.xor_byte
            ))

        time.sleep(args.frame_delay)

if __name__ == "__main__":
    main()
