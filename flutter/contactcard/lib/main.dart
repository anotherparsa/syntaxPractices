import 'package:flutter/material.dart';

void main() {
  runApp(
    MaterialApp(
      debugShowCheckedModeBanner: false,
      home: Scaffold(
        body: Center(
          child: Column(
            children: [
              CircleAvatar(
                radius: 50.0,
                backgroundColor: Colors.white,

              ),
            ],
          ),
        ),

        appBar: AppBar(
          title: Text("Contact Card"),
          backgroundColor: Colors.white,
          centerTitle: true,
        ),

        backgroundColor: Colors.black,
      ),
    ),
  );
}
