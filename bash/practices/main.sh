#/bash/bin

array=(one two three four five six seven eigth nine ten)
for element in "${array[@]}"; do
	echo "This is $element"
done
