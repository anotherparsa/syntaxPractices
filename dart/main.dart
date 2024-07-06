void main() {
  var intVariable = returnOne();
  print(intVariable);
  var TwoVariable = returnTwo();
  print(TwoVariable);
}

int returnOne() {
  return 12;
}

(int, String) returnTwo() {
  return (12, "!@");
}
