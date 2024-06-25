#/bin/bash
#validating age
echo "You're $1 Years old"
if [ $1 -ge 20 ]
then
	echo "You're in legal age"
else
	echo "You're not in legal age"
fi
