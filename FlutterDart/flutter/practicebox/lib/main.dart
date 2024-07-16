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
        color: Colors.white,
        child: Column(
          children: [
            Padding(
              padding: EdgeInsets.only(top: 20, left: 200),
              child: Text("This is a text 1"),
            ),
            Padding(
              padding: EdgeInsets.only(top: 20, left: 200),
              child: Text("This is a text 2", textAlign: TextAlign.left,),
            ),
          ],
        ),
      ),
      backgroundColor: Colors.white,
    ),
  ));
}
