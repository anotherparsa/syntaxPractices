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
        width: double.infinity,
        height: 150,
        margin: EdgeInsets.all(30),
        child: Text(
          "This text is inside a container",
          style: TextStyle(fontSize:25),
        ),
      ),
      backgroundColor: Colors.white,
    ),
  ));
}
