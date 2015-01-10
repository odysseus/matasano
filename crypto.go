package main

import (
  "encoding/hex"
  "encoding/base64"
  "fmt"
  "os"
  "errors"
  "strings"
  "sort"
)

func main() {
  fmt.Println(challenge1())
  fmt.Println(challenge2())
  fmt.Println(challenge3())
}

func chkErr(err error) {
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
}

func base64MimeEncoding() *base64.Encoding {
  alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
  e := base64.NewEncoding(alpha)
  return e
}

func base64MimeEncodedString(b []byte) string {
  e := base64MimeEncoding()
  return e.EncodeToString(b)
}

func fixedLenXOR(a, b []byte) ([]byte, error) {
  if len(a) != len(b) {
    return nil, errors.New("fixedLenXOR: Byte arrays must be the same length")
  }
  c := make([]byte, len(a))
  for i:=0; i<len(a); i++ {
    c[i] = a[i] ^ b[i]
  }
  return c, nil
}

func challenge1() bool {
  answer := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

  a := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
  b, _ := hex.DecodeString(a)
  e := base64MimeEncoding()
  c := e.EncodeToString(b)
  if c == answer { return true } else { return false }
}

func challenge2() bool {
  answer := "746865206b696420646f6e277420706c6179"
  a := "1c0111001f010100061a024b53535009181c"
  b := "686974207468652062756c6c277320657965"
  c, _ := hex.DecodeString(a)
  d, _ := hex.DecodeString(b)

  e, _ := fixedLenXOR(c, d)
  final := hex.EncodeToString(e)
  if final == answer { return true } else { return false }
}

func challenge3() string {
  a := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
  b, _ := hex.DecodeString(a)

  var final string
  maxIntersect := 0
  freq := "ETAONI"

  for i:=0; i<256; i++ {
    c := make([]byte, len(b))
    for j:=0; j<len(b); j++ {
      c[j] = byte(i)
    }
    d, _ := fixedLenXOR(b, c)
    sd := string(d)
    di := stringIntersection(freq, sd)
    if di > maxIntersect {
      final = sd
      maxIntersect = di
      fmt.Println(di)
    }
  }
  return final
}

func charCount(s string) map[rune]int {
  m := make(map[rune]int)
  for _, c := range s {
    m[c] += 1
  }
  return m
}

type charPair struct {
  char rune
  count int
}

func (cp charPair) String() string {
  return fmt.Sprintf("%v: %v", cp.char, cp.count)
}

type CharCount []charPair

func (cc CharCount) Len() int {
  return len(cc)
}

func (cc CharCount) Less(i, j int) bool {
  return cc[i].count < cc[j].count
}

func (cc CharCount) Swap(i, j int) {
  cc[i], cc[j] = cc[j], cc[i]
}

func frequencyString(s string) string {
  fs := make([]rune, 0)
  us := strings.ToUpper(s)

  // First strip the non alphabetic characters
  alphas := make([]rune, 0)
  alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
  for _, sc := range us {
    for _, ac := range alpha {
      if sc == ac {
        alphas = append(alphas, sc)
        break
      }
    }
  }
  as := string(alphas)

  m := charCount(as)
  vals := make(CharCount, 0)
  for k, v := range m {
    vals = append(vals, charPair{k, v})
  }
  sort.Sort(vals)
  for i:=len(vals)-1; i>=0; i-- {
    fs = append(fs, vals[i].char)
  }

  return string(fs)
}

func stringIntersection(a, b string) int {
  var l, count int

  fa, fb := frequencyString(a), frequencyString(b)
  la, lb := len(fa), len(fb)
  if la != lb {
    if la < lb {
      l = la
    } else {
      l = lb
    }
  }
  for i:=0; i<l; i++ {
    c := fa[i]
    for j:=0; j<l; j++ {
      if c == fb[j] {
        count++
        break
      }
    }
  }
  return count
}
