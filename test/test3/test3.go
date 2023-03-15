package main

import (
	"archive/tar"
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

func main() {
	// defaultType()
	//testTar()
	//t := time.NewTimer(time.Second)

	testXOrBytes()
}

func defaultType() {
	a1 := 'a'               // byte        97 is in the set of byte values
	a2 := 97                // rune        rune is an alias for int32, and 97 is in the set of 32-bit integers
	a3 := "foo"             // string      "foo" is in the set of string values
	a4 := 10241111          // int16       1024 is in the set of 16-bit integers
	a5 := 42.0              // byte        42 is in the set of unsigned 8-bit integers
	a6 := 1e10              // uint64      10000000000 is in the set of unsigned 64-bit integers
	a7 := 2.718281828459045 // float32     2.718281828459045 rounds to 2.7182817 which is in the set of float32 values
	a8 := -1e-1000          // float64     -1e-1000 rounds to IEEE -0.0 which is further simplified to 0.0
	a9 := 0i                // int         0 is an integer value
	a10 := (42 + 0i)        // float32     42.0 (with zero imaginary part) is in the set of float32 values
	a11 := 0                // bool        0 is not in the set of boolean values
	a12 := -1               // uint16      -1 is not in the set of unsigned 16-bit integers
	a13 := 42i              // float32     (0 + 42i) is not in the set of float32 values

	fmt.Println(strconv.FormatFloat(a7, 'f', 20, 64))

	// 1100010111000111000000100111111111100001111100110101000111100011
	var b uint64 = 14251362294112145891
	fmt.Printf("1, %b %v \n", b, b)
	b1 := "c5c7027fe1f351e3"
	v1, _ := strconv.ParseUint(b1, 16, 64)
	fmt.Printf("2, %b %v \n", v1, v1)

	for i := 0; i < len(b1); i += 4 {
		v2, _ := strconv.ParseUint(b1[i:i+4], 16, 64)
		fmt.Printf("3, %b %v %s \n", v2, v2, b1[i:i+4])
	}

	fmt.Printf("%T, %T, %T, %T, %T, %T, %T, %T, %T, %T, %T, %T, %T \n", a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)

	const arr = 1 + 2
	fmt.Printf("%T\n", arr)

	s := []int{1, 2, 3, 4, 5}
	s1 := make(map[string]string, 3)
	fmt.Println(s, s1)

	//m := map[string]string{"a": "aa"}
	//fmt.Println(m)

	t := reflect.TypeOf(s1)
	fmt.Println(111, t)

	//b := 12345
	//fmt.Printf("%02d", b)
}

func testTar() {

	// Create and add some files to the archive.
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	// Open and iterate through the files in the archive.
	tr := tar.NewReader(&buf)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}

}

func testXOrBytes() {
	// === 1101 0111 0001 0011 1111 1101 0001 1000 0110 1010 0010 0001 1011 0111 1111 1111 15498009024013645823 d713fd186a21b7ff
	// === 1101 0111 0010 0011 1011 1101 0001 1000 0110 1010 0010 0001 1011 0111 1111 1111 15502442254896838655 d723bd186a21b7ff

	s1 := "d712a5186a29f607"
	s2 := "d712a5186a29f7a7"
	bytes1, _ := hex.DecodeString(s1)
	bytes2, _ := hex.DecodeString(s2)
	fmt.Printf("%08b %v\n", bytes1, bytes1)
	fmt.Printf("%08b %v\n", bytes2, bytes2)

	fmt.Printf("%v\n", time.Now().UnixMilli())

	// var r []byte
	// n := fastxor.Bytes(r, sourceBytes, bytes)
	// fmt.Printf("xor, %s, %s, n: %v", sourceBytes, bytes, n)

	// s3 := "d713fd186a21b7ff"
	// bytes3, _ := hex.DecodeString(s3)
	// // fmt.Printf("%b, %v\n", bytes3, math.Log2(float64(binary.BigEndian.Uint64(bytes3))))

	// i1 := binary.BigEndian.Uint64(bytes3)
	// fmt.Println(countBit1(i1))

	// a := 2011
	// b := 2010
	// v := a ^ b
	// fmt.Printf("%020b ^ %b = %b %v\n", a, b, v, v)

	// fmt.Printf("%8s\n", "abc")
}

func countBit1(n uint64) (count uint64) {
	for n != 0 {
		count += n & uint64(1)
		n >>= 1
	}
	return
}
