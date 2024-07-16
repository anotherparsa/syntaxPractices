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
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text("Text 1", style: TextStyle(color: Colors.blue, fontSize: 50),),
          Text("Text 2", style: TextStyle(color: Colors.red, fontSize: 50),),
          Text("Text 3", style: TextStyle(color: Colors.green, fontSize: 50),),
          FlutterLogo(size: 250,)
        ],
      ),
      backgroundColor: Colors.white,
    ),
  ));
}
