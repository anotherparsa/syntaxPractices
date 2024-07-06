void main() {
  var intVariable = returnOne();
  print(intVariable);
  var TwoVariable = returnTwo();
  print(TwoVariable);
  var (number, string, boolean) = returnThree();
  print(number);
  print(string);
  print(boolean);
}

int returnOne() {
  return 12;
}

(int, String) returnTwo() {
  return (12, "!@");
}

(int, String, bool) returnThree() {
  return (12, "!@", false);
}
