def single_byte(bytes, xor_byte):
    return bytearray(byte ^ xor_byte for byte in bytearray(bytes))
