# Before starting

## Go's purpose

- Built from other languages good ideas,
- Remove all features that could lead to code complexity

One unique 🎯 : write simple, expressive, robust and efficient code.

Notes:

- mainly inspired from C.
- Fits big projects with large teams

##==##

# Before starting

## A general language

- Like C, fits almost all programmation domains (web, IoT, Cloud, Data)
- The perfect fit for Cloud ☁️
- Already used for HMI's, CLI apps, mobile apps, machine learning, WASM, ...

##==##

# Before starting

Adopted by big companies

- Google
- Docker
- Kubernetes
- Dropbox
- Spotify
- Hashicorp
- SoundCloud
- etc...

##==##

# The language

## Back in 2009

- Born in 2009 at Google (after multi-core CPU) and OSS
- Self-contained binaries
- Object-oriented
- Garbage collected (sub millisecond for 17 Go of heap)
- Pointers 😱
- Goroutines

  - Like a thread
  - But **not** un thread ⇒ **a lighter abstraction**

- Channels
- **Do not communicate by sharing memory; share memory by communicating.**
  Synchronisation
  Multiplexing (**select**)

Notes:

- CSP (1977) Communicating sequential processes

##==##

# The language

## Few keywords

- **Dependencies :** import package
- **Conditions :** if else switch case fallthrough break default goto select
- **Iterations :** for range continue
- **Type :** var func interface struct chan const type map make
- **Misc :** defer go return panic recover

![h-350](assets/go-100/images/few_keywords.jpg)<!-- .element: class="special-Intro-01-le-but-de-go-bottom-image" -->
![h-350](assets/go-100/images/i_know.jpeg)<!-- .element: class="special-Intro-01-le-but-de-go-bottom-image" -->
Notes:
- less than 30 keywords
- public/private => upper case / lower case
