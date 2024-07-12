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
              Text(
                "Your Name",
                style: TextStyle(color: Colors.white, fontSize: 25),
              ),
              Text(
                "Your Occupation",
                style: TextStyle(color: Colors.white, fontSize: 25),
              ),
              Row(
                children: [
                  Text(
                    "Test 1",
                    style: TextStyle(color: Colors.white),
                  ),
                  Icon(
                    Icons.phone,
                    color: Colors.white,
                  ),
                  Text(
                    "Test 2",
                    style: TextStyle(color: Colors.white),
                  ),
                  Icon(
                    Icons.email,
                    color: Colors.white,
                  )
                ],
              ),
            ],
          ),
        ),
        appBar: AppBar(
          title: Text(
            "Contact Card",
            style: TextStyle(color: Colors.white),
          ),
          backgroundColor: Colors.black,
          centerTitle: true,
        ),
        backgroundColor: Colors.black,
      ),
    ),
  );
}
