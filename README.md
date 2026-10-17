# Tumai

Offline flashcard app with no limit of cards in single session. Made for learning definitions and vocabulary of foreign language. Remembers learned.

### Requirements

- go 1.25.0
- wails 2.15.0
- libwebkit2gtk-4.0-dev or `libwebkit2gtk-4.0.so`
- make

### Usage

Add excel file to the list. Can be MS spreadsheet or Libre.
Pick file and set option:
- Show: Name / Meaning
- Filter: `not implemented`
- Order: Random / As in the file

`<` - shows previous card.
`Next!` - marks as learned and moves to the next card.
`>` - marks as not learned and moves on.

Scoreboard shows number of learned / tested / all cards in the set.

The excel must contain a header (first row is not parsed to the deck) and 5 columns of which:
1. Name - the single word / term you want to memorize
2. Phonetics - proper pronunciation
3. Meaning - the definition of the term or meaning of the word
4. Notes - optional additional information
5. Status - this column is written by the program (boolean) and is preferred empty when opened for the first time.

Do not touch any spreadsheet opened in the app while it is running. Do not tamper with `~/.Tumai.list` as it may render your list unusable.

# Issues

The state of this application is `as is` and I do not provide any guaranty. If you find a bug then report it or whatever then pray I open github the same year.

If I decide I have too much spare time (or I'll be forced by the lack of sufficient learning tools) I'll add the `Filter` button, images on flashcards.
If you see this ported to windows, know my day sucked.
