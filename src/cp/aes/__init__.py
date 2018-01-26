import collections
import random

import oscrypto.symmetric

from cp import blocks, xor

BLOCK_SIZE = 16

# OSCrypto doesn't provide ECB mode (for obvious reasons),
# so we'll simulate by encrypting blocks in isolation with
# a null IV.

def encrypt_ecb(key, data):
    key = bytes(key)
    null_iv = bytes([0] * BLOCK_SIZE)
    result = bytearray()
    for position in blocks.positions(len(data), BLOCK_SIZE):
        block = bytes(position.slice(data))
        _, enc_block = oscrypto.symmetric.aes_cbc_no_padding_encrypt(key, block, null_iv)
        result.extend(enc_block)
    return result

def decrypt_ecb(key, data):
    key = bytes(key)
    null_iv = bytes([0] * BLOCK_SIZE)
    result = bytearray()
    for position in blocks.positions(len(data), BLOCK_SIZE):
        enc_block = bytes(position.slice(data))
        dec_block = oscrypto.symmetric.aes_cbc_no_padding_decrypt(key, enc_block, null_iv)
        result.extend(dec_block)
    return result

def random_iv():
    sr = random.SystemRandom()
    return bytearray(sr.randint(0,255) for _ in xrange(BLOCK_SIZE))

def encrypt_cbc(key, iv, data):
    key = bytes(key)
    null_iv = bytes([0] * BLOCK_SIZE)
    prev_enc_block = iv
    result = bytearray()
    for position in blocks.positions(len(data), BLOCK_SIZE):
        block = bytes(position.slice(data))
        block = bytes(xor.repeating_bytes(block, prev_enc_block))
        _, enc_block = oscrypto.symmetric.aes_cbc_no_padding_encrypt(key, block, null_iv)
        result.extend(enc_block)
        prev_enc_block = enc_block
    return result

def decrypt_cbc(key, iv, data):
    key = bytes(key)
    null_iv = bytes([0] * BLOCK_SIZE)
    prev_enc_block = bytes([0] * BLOCK_SIZE)
    result = bytearray()
    for position in blocks.positions(len(data), BLOCK_SIZE):
        enc_block = bytes(position.slice(data))
        dec_block = oscrypto.symmetric.aes_cbc_no_padding_decrypt(key, enc_block, null_iv)
        dec_block = xor.repeating_bytes(dec_block, prev_enc_block)
        result.extend(dec_block)
        prev_enc_block = enc_block
    return result

def count_blocks(data):
    return collections.Counter(
        bytes(position.slice(data))
        for position in blocks.positions(len(data), BLOCK_SIZE)
    )

def detect_ecb(data, min_repeat=2):
    return count_blocks(data).most_common(1)[0][1] > min_repeat
