package main

import (
	// "strconv"
	"fmt"
)

func main() {
	fmt.Println("I will do it !");

	W := 'w';
	byteMat := byteBlockEncoder(W);

	fmt.Println(byteMat);
}

func convertToBinary(num int64) string {

	var binaryString string;
	for num > 0 {

		if (num & 1) > 0 {
			binaryString += "1"
		} else {
			binaryString += "0"
		}
		num = num >> 1;
	}

	// reversing the string
	runes := []rune(binaryString)
	s, e := 0, len(binaryString) - 1
	for s > e {
		runes[s], runes[e] = runes[e], runes[s]
		s++;
		e--;
		// binaryString[s], binaryString[e] = binaryString[e], binaryString[s]
	}

	return string(runes)
}

// writing the byte-block encoder
func byteBlockEncoder(ch rune) [4][2]int {

	/*
		For the version-2 QR code 
		we use a 4x2 matrix which gives us 8 bits (1 byte)
		to encode one character 
	*/

	// Get the binary representation of the character (rune here in go-lang)
	asciiCh := int(ch);
	fmt.Println(asciiCh);
	// Could have written the function to convert to binary myself but let's roll with the library for now.
	// binaryValue := strconv.FormatInt(int64(asciiCh), 2);
	binaryValue := convertToBinary(int64(asciiCh));
	fmt.Println(binaryValue);
	// rune of the binaryValue string
	runes := []rune(binaryValue)

	fmt.Println(runes);

	var byteMatrix[4][2] int;

	lsbPointer := len(binaryValue) - 1

	for i := 0; i < len(byteMatrix); i++ {
		for j := 0; j < len(byteMatrix[i]); j++ {

			if lsbPointer >= 0 {

				if runes[lsbPointer] == '1' {
					byteMatrix[i][j] = 1
				} else {
					byteMatrix[i][j] = 0
				}
				// byteMatrix[i][j] = int()
			} else {
				break;
			}

			lsbPointer--;
		}
	}

	return byteMatrix;
}