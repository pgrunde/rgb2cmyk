## Golang Channel Example Tutorial

The following tutorial is similar to an example found in [Programming in Go: Creating Applications for the 21st Century](http://www.amazon.com/Programming-Go-Creating-Applications-Developers/dp/0321774639) by Mark Summerfield, check out his book for more excellent Golang learning materials.

### Overview

The application we will be writing takes in [RGB](http://en.wikipedia.org/wiki/RGB_color_model#Numeric_representations) color values and returns equivalent [CMYK](http://en.wikipedia.org/wiki/CMYK_color_model) color values (don't worry we won't be getting [Neugebauer](http://en.wikipedia.org/wiki/Neugebauer_equations) fancy). The fancy bit will be writing it using go routines and channels! Please note that proper scholars would likely obstain from using channels for this task, but we're not proper scholars, we're here to learn. We're going to be using three packages "bufio", "fmt" and "os". Outside of our main function, we're going to write a function which creates a go routine for our happy color-computation channel. Then we make a function which handles input from the user by chuckin' it into an input channel, getting the result out of our output channel and printing it for all to see.

### The setup: what do we want?

We will need a prompt so the user knows what the heck this program is for, and we'll need a result string which contains the output of our CMYK channel.

`const result = "RGB R%v G%v B%v == C:%.03f M:%.03f Y:%03f K:%.03f\n"` 
