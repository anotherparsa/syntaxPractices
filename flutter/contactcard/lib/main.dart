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
            child: Text("1ST in column"),
            color: Colors.red,
          ),
          Container(
            width: 100,
            height: 100,
            child: Text("2ND in column"),
            color: Colors.blue,
          ),
          Row(
            children: [
              Container(
                width: 50,
                height: 50,
                child: Text("1ST in row"),
                color: Colors.pink,
              ),
              Container(
                width: 50,
                height: 50,
                child: Text("2ND in row"),
                color: Colors.purple,
              )
            ],
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
