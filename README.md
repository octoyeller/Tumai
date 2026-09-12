# Tumai

1. App start
2. Read list file
    - if no file => create
    - if can't read => splash screen
3. Print list
    - if the are more than 5-7 files => button to show all
    - the entry on visible list is button to add more
4. Chose file / option
    - new file already on list => open that file
    - file non-existent => splash screen? & back to list screen
5. Open file for processing
    - 
    - 
6. Exit file => modify the date in the file
7. On exit save the file


ORDER - as goes / random 
SELECT - all / learned / not learned
WHAT - name / meaning
SAVE
EXIT

excel structure
Sheet1
col 1 - name
col 2 - phonetic
col 3 - meaning
col 4 - notes   (optional)
col 5 - learned / not learned   (optional - will be added on save)


Contents of the list file:
```
12314\t/absolute/path
12314\t/absolute/path
int seconds tab path
```
The integer is date of last registered access is seconds. The date is registered by the app when opening the file via this application.
