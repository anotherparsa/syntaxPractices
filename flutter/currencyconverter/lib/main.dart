import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    debugShowCheckedModeBanner: false,
    home: Scaffold(
      appBar: AppBar(
        title: Text(
          "Currency Converter",
          style: TextStyle(color: Colors.white, fontSize: 20),
        ),
        backgroundColor: Colors.black,
        centerTitle: true,
      ),
      body: Center(
          child: ColoredBox(
        color: Colors.blue,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              "Test",
              style: TextStyle(color: Colors.white, fontSize: 30),
            ),
          ],
        ),
      )),
      backgroundColor: Colors.black,
    ),
  ));
}
