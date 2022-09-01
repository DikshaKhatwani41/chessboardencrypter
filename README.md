# chessboardencrypter
<h2>Encrypted chessboard encoded with Go Programming Language</h2>

## Languages used:
Go, HTML, CSS

## Database:
Go's in-built Database

## Description/ Demonstration

<strong>
Are you a spy that needs to leave an urgent top secret message for an incoming foreign diplomat but the only thing in the consulate’s parlor is a chessboard? Chesscode has you covered.
Chesscode is a way of encoding messages using a chess board and the starting set of chess pieces. The message can be up to 23 alphanumeric characters also allowing spaces and periods.
</strong>
<br/>
<br/>

![demo](https://user-images.githubusercontent.com/45715538/185772308-7077d116-1eb7-41c9-8860-00e26a837e1d.gif)


<strong>Encoding</strong>

Since the set of possible pieces is fixed as a subset of the starting position, this infinitely limits the theoretical possibilities of 13^64 configurations. Another issue is piece distribution, there is one white king and eight white pawns for example. Because of these constraints, pieces signify the order of data and squares hold the data itself.

**Piece Order**
<ul>
<li>White King
<li>Black King
<li>White Queen
<li>Black Queen
<li>White Rooks
<li>Black Rooks
<li>White Bishops
<li>Black Bishops
<li>White Knights
<li>Black Knights
<li>White Pawns
<li>Black Pawns
</ul>
This pattern is easiest to see with the white king. The first example with the king in the A1 square encodes 0, while the king in the C2 encodes A. The square simply indexes a character in the charset. This gets us one character.

Charset (space at end):
0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ 

<strong>Decoding</strong>

Decoding simply reverses the encoding process, but there is a few caveats:
<ul>
  <li>Not every board with the starting pieces can be decoded</li>
  <li>It is ambiguous if the message intended to leave white space at the end of the message</li>
</ul>
