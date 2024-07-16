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
        child: Text("This text is inside a container"),
      ),
      backgroundColor: Colors.white,
    ),
  ));
}
