package fenam_cipher

import (
	"fmt"
	cycle_string "github.com/cryptography-research-lab/go-cycle-string"
	if_expression "github.com/golang-infrastructure/go-if-expression"
	stack "github.com/golang-infrastructure/go-stack"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
	"math"
	"strings"
)

// ------------------------------------------------- --------------------------------------------------------------------

// 当不指定秘钥时默认的秘钥
var defaultSecurityKey = "CC"

// Encrypt 对文本使用费娜姆加密
func Encrypt(asciiText string, securityKey ...string) (string, error) {
	securityKey = variable_parameter.SetDefaultParam(securityKey, defaultSecurityKey)
	plaintextBinaryString, err := convertAsciiStringToBinaryString(asciiText)
	if err != nil {
		return "", err
	}
	securityBinaryString, err := convertAsciiStringToBinaryString(securityKey[0])
	if err != nil {
		return "", err
	}
	return binaryStringXOR(plaintextBinaryString, securityBinaryString), nil
}

// 把ascii字符串转为二进制字符串
func convertAsciiStringToBinaryString(asciiText string) (string, error) {
	result := strings.Builder{}
	for _, char := range asciiText {
		if !(char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z') {
			return "", ErrEncryptText
		}
		binaryString := intToBinaryString(int(char))
		result.WriteString(binaryString)
	}
	return result.String(), nil
}

// 不使用内置库了，自己写一下
func intToBinaryString(number int) string {
	// 除2取余，逆序排列
	stack := stack.NewStack[int]()
	for number > 0 {
		stack.Push(number % 2)
		number /= 2
	}

	// 补前缀零
	for stack.Size() < 7 {
		stack.Push('0')
	}

	// 转为字符串
	result := strings.Builder{}
	for stack.IsNotEmpty() {
		result.WriteString(fmt.Sprintf("%d", stack.Pop()))
	}
	return result.String()
}

// 使用秘钥再进行一次加密
func binaryStringXOR(binaryStringA, binaryStringB string) string {
	cycleBinaryString := cycle_string.NewCycleString(binaryStringB)
	result := make([]rune, len(binaryStringA))
	for index, charA := range binaryStringA {
		charB := cycleBinaryString.RuneAt(index)
		// 异或运算
		if charA == charB {
			result[index] = '0'
		} else {
			result[index] = '1'
		}
	}
	return string(result)
}

// ------------------------------------------------- --------------------------------------------------------------------

// Decrypt 对文本进行费娜姆解密
func Decrypt(encryptBinaryText string, securityKey ...string) (string, error) {
	securityKey = variable_parameter.SetDefaultParam(securityKey, defaultSecurityKey)

	// 把秘钥也转为二进制字符串的形式
	securityKeyBinaryString, err := convertAsciiStringToBinaryString(securityKey[0])
	if err != nil {
		return "", err
	}

	// 然后xor解密
	plaintextBinaryString := binaryStringXOR(encryptBinaryText, securityKeyBinaryString)
	// 转为可读的ascii形式
	return convertBinaryStringToAsciiString(plaintextBinaryString)
}

func convertBinaryStringToAsciiString(binaryString string) (string, error) {
	result := strings.Builder{}
	index := 0
	for index < len(binaryString) {
		nextIndex := index + 7
		if nextIndex > len(binaryString) {
			return "", fmt.Errorf("binary string not mod 7: %s", binaryString)
		}
		groupBinaryString := binaryString[index:nextIndex]
		char := fromBinaryString(groupBinaryString)
		if !(char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z') {
			return "", ErrEncryptText
		}
		result.WriteRune(rune(char))
		index = nextIndex
	}
	return result.String(), nil
}

// 把二进制字符转为十进制数字
func fromBinaryString(binaryString string) int {
	result := 0
	weight := 0
	for index := len(binaryString) - 1; index >= 0; index-- {
		result += int(math.Pow(float64(2), float64(weight))) * if_expression.Return(binaryString[index] == '0', 0, 1)
		weight++
	}
	return result
}

// ------------------------------------------------- --------------------------------------------------------------------
