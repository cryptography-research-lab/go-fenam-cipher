package fenam_cipher

import "errors"

// ErrEncryptText 要进行加密的文本必须只能包含英文字母
var ErrEncryptText = errors.New("text for encrypt only can contains letters")
