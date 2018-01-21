import unittest

from cp import hamming

class Test(unittest.TestCase):
    def test_distance(self):
        self.assertEqual(0, hamming.distance(
            bytearray("same string", "utf8"),
            bytearray("same string", "utf8")
        ))

        self.assertEqual(37, hamming.distance(
            bytearray("this is a test", "utf8"),
            bytearray("wokka wokka!!!", "utf8")
        ))
