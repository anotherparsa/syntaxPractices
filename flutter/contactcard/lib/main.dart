import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    home: Scaffold(
      body: Center(
          child: Column(
        children: [
          Container(
            width: 100,
            height: 100,
            child: Text("1ST"),
            color: Colors.red,
          ),
          Container(
            width: 100,
            height: 100,
            child: Text("2ND"),
            color: Colors.blue,
          )
        ],
      )),
      appBar: AppBar(
        title: Text("Contact Card"),
        backgroundColor: Colors.white,
        centerTitle: true,
      ),
      backgroundColor: Colors.white,
    ),
  ));
}
