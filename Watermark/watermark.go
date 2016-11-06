//this will impose a image a image on another image .. 
// for this I used watermark.png to impose it on image.jpg
package main

import (
    "image"
    "image/draw"
    "image/jpeg"
    "image/png"
    "os"
)

func main() {

    imgb, _ := os.Open("image.jpg")
    img, _ := jpeg.Decode(imgb)
    defer imgb.Close()

    wmb, _ := os.Open("watermark.png")
    watermark, _ := png.Decode(wmb)
    defer wmb.Close()

    offset := image.Pt(200, 200)
    b := img.Bounds()
    m := image.NewRGBA(b)
    draw.Draw(m, b, img, image.ZP, draw.Src)
    draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

    imgw, _ := os.Create("watermarked.jpg")
    jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
    defer imgw.Close()

    fmt.Println("DONE !!")
}