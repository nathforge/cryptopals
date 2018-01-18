import bisect
import collections

from cp import xor

class SingleByteKey(object):
    Guess = collections.namedtuple("Guess", ("score", "xor_byte", "dec"))
    State = collections.namedtuple("State", ("guess", "sorted_guesses"))

    def __init__(self, scorer, bytes):
        self.scorer = scorer
        self.bytes = bytes

    def __iter__(self):
        sorted_guesses = []
        for xor_byte in range(256):
            dec = xor.single_byte(self.bytes, xor_byte)
            score = self.scorer(dec)
            guess = self.Guess(score, xor_byte, dec)
            bisect.insort_left(sorted_guesses, guess)
            yield self.State(guess, sorted_guesses)
