import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    debugShowCheckedModeBanner: false,
    home: Scaffold(
      appBar: AppBar(
        centerTitle: true,
        title: Text("Practice Box"),
        backgroundColor: Colors.blue,
      ),
      body: Container(
        color: Colors.green,
        child: Text(
          "This text is inside a container",
          style: TextStyle(fontSize: 50),
        ),
      ),
      backgroundColor: Colors.white,
    ),
  ));
}
