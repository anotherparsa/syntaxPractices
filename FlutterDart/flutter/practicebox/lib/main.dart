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
      body: Center(
        child: ElevatedButton(
          onPressed: () {},
          child: Text("Press Me"),
        ),
      ),
      backgroundColor: Colors.white,
    ),
  ));
}
