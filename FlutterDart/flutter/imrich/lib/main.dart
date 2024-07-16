import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    home: Scaffold(
      appBar: AppBar(
        title: Text(
          "I'm Rich",
          style: TextStyle(color: Colors.white),
        ),
        backgroundColor: Colors.black,
        centerTitle: true,
      ),
      body: Center(
        child: Text(
          "I'm Rich",
          style: TextStyle(color: Colors.white),
        ),
      ),
      backgroundColor: Colors.black,
    ),
  ));
}
