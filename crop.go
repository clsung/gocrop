package main

import (
  "fmt"
  "os"
  "flag"
  "image"
  _ "image/gif"
  "image/jpeg"
  _ "image/png"
  "code.google.com/p/graphics-go/graphics"
)

func main() {
  flag.Parse()
  args := flag.Args()
  if len(args) < 1 {
    fmt.Println("Input file is missing.");
    os.Exit(1);
  }
  if len(args) < 2 {
    fmt.Println("Output file is missing.");
    os.Exit(1);
  }
  fmt.Printf("opening %s\n", args[0])
  fSrc, _ := os.Open(args[0])
  defer fSrc.Close()
  src, _, _ := image.Decode(fSrc)
  dst := image.NewRGBA(image.Rect(0, 0, 80, 80))
  graphics.Thumbnail(dst, src)
  toimg, _ := os.Create(args[1])
  defer toimg.Close()

  jpeg.Encode(toimg, dst, &jpeg.Options{jpeg.DefaultQuality})
}

