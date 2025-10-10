Command to create a 100MB file
```
yes "The quick brown fox jumps over the lazy dog near the riverbank while searching for food" | nl -s " - " | head -c 104857600 > output.txt
```

Resrouce : https://stackoverflow.com/questions/1450551/buffered-vs-unbuffered-io

