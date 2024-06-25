#/bash/bin

read -p "Enter your age: " age
echo "You're ${age} years old"
if [ $age -ge 18 ]
then
	echo "You're in legal age"
else
	echo "You're not in legal age" 
fi
