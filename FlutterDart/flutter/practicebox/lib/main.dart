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
        child: Align(
          alignment: Alignment.bottomCenter,
          child: Text("Test Text", style: TextStyle(fontSize: 25),
        )),
      ),
      backgroundColor: Colors.white,
    ),
  ));
}
