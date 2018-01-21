import bisect
import collections
import itertools

from cp import blocks, hamming, xor

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

class KeySize(object):
    Guess = collections.namedtuple("Guess", ("distance", "key_size"))
    State = collections.namedtuple("State", ("guess", "reverse_sorted_guesses"))

    def __init__(self, ba, min_key_size=2, max_key_size=40, sample_block_count=4):
        self.ba = bytearray(ba)
        self.min_key_size = min_key_size
        self.max_key_size = max_key_size
        self.sample_block_count = sample_block_count

    def __iter__(self):
        reverse_sorted_guesses = []
        for key_size in range(self.min_key_size, self.max_key_size + 1):
            iter_positions = blocks.positions(len(self.ba), key_size)

            distances = []
            first_block = None
            prev_block = None
            for position in itertools.islice(iter_positions, self.sample_block_count):
                # Break if block size isn't as expected.
                # This occurs when data size isn't a multiple of the key size
                # and we've reached end of the data.
                if key_size != (position.end - position.start):
                    break

                block = list(position.slice(self.ba))
                if prev_block is None:
                    first_block = block
                else:
                    distances.append(hamming.distance(block, prev_block))
                prev_block = block

            if block != first_block:
                distances.append(hamming.distance(block, first_block))

            average_distance = (sum(distances) / float(len(distances))) / key_size

            guess = self.Guess(average_distance, key_size)
            bisect.insort_left(reverse_sorted_guesses, guess)
            yield self.State(guess, reverse_sorted_guesses)
