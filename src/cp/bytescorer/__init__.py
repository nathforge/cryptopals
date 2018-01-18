class ByteScorer(dict):
    def __init__(self, ba):
        for byte in bytearray(ba):
            self[byte] = self.get(byte, 0) + 1
        for byte, score in self.items():
            self[byte] /= float(len(ba))

    def __call__(self, text):
        return (
            sum(self.get(byte, 0) for byte in bytearray(text)) /
            float(len(text))
        )
