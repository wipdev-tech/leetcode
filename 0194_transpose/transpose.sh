# UNSOLVED - file can have 1 column or more than two
awk '{print $1}' file.txt | tr '\n' ' '
echo
awk '{print $2}' file.txt | tr '\n' ' '
