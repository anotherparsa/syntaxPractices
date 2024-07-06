void main() {
  int age = 20;
  if (age >= 18) {
    print("you're in legal age");
  } else {
    print("you're not in legal age");
  }

  String boolean = 20 != 20 ? "true" : "false";
  print(boolean);

  int day_of_the_weeek = 34;
  switch (day_of_the_weeek) {
    case 1:
      print("Shanbe");
    case 2:
      print("1 shanbe");
    case 3:
      print("2 shanbe");
    case 4:
      print("3 shanbe");
    case 5:
      print("4 Shanbe");
    case 6:
      print("5 shanbe");
    case 7:
      print("Jome");
    default:
      print("invalid option");
  }

  for (int i = 0; i < 10; i++) {
    print("${i}, times");
  }
}
