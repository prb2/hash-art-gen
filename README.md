# hash-art-gen

Ref: [The drunken bishop: An analysis of the OpenSSH fingerprint visualization algorithm](http://www.dirk-loss.de/sshvis/drunken_bishop.pdf)

**Usage**
```sh
> go run hash-art-gen.go -seed="alsdfasasdkjflaasdfasdf"
```

**Example output**
```sh
Generating random art with seed

        alsdfasasdkjflaasdfasdf

SHA256 hash

        887d468e4afc2635eeb351b74f753e6576ca0b11320c8758a3e9c0cbe22a516e

Augmentation runes

        [  . o + = * B O X @ % & # / ^]

Gen art

        +-------------------+
        |      ..+..        |
        |     . o.+ +       |
        |    o o o + o      |
        |   ..o + + o    .o |
        |   o. * = o +   o= |
        | ....+ B * + + o * |
        | ..o. + + o o o +  |
        | ..    . +         |
        | .     .+          |
        +-------------------+


```

**TODO**

- random seed

- arbitrary grid size

- customizable augmentation string
