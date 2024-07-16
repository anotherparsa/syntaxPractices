import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    debugShowCheckedModeBanner: false,
    home: Scaffold(
      appBar: AppBar(
        title: Text(
          "Contact Card",
          style: TextStyle(color: Colors.white),
        ),
        centerTitle: true,
        backgroundColor: Colors.black,
      ),
      body: Center(
        child: Column(
          children: [
            Padding(
              padding: EdgeInsets.only(top: 20),
              child: CircleAvatar(
                backgroundColor: Colors.white,
                radius: 50.0,
              ),
            ),
            Padding(
              padding: EdgeInsets.only(top: 20),
              child: Text(
                "Your name",
                style: TextStyle(color: Colors.white, fontSize: 30),
              ),
            ),
            Padding(
              padding: EdgeInsets.only(top: 20),
              child: Text(
                "Your occupation",
                style: TextStyle(color: Colors.white, fontSize: 30),
              ),
            ),
            Padding(
              padding: EdgeInsets.only(top: 20),
              child: Container(
                margin: EdgeInsets.only(left: 30, right: 30),
                color: Colors.white,
                child: Row(
                  children: [
                    Padding(
                      padding: EdgeInsets.only(top: 10, bottom: 10, left: 40),
                      child: Icon(
                        Icons.phone,
                        size: 40,
                        color: Colors.black,
                      ),
                    ),
                    Padding(
                      padding: EdgeInsets.only(
                          top: 10, bottom: 10, left: 20, right: 20),
                      child: Text(
                        "0123456789",
                        style: TextStyle(color: Colors.black, fontSize: 40),
                      ),
                    )
                  ],
                ),
              ),
            ),
            Padding(
              padding: EdgeInsets.only(top: 20),
              child: Container(
                margin: EdgeInsets.only(left: 30, right: 30),
                color: Colors.white,
                child: Row(
                  children: [
                    Padding(
                      padding: EdgeInsets.only(top: 10, bottom: 10, left: 40),
                      child: Icon(
                        Icons.email,
                        size: 40,
                        color: Colors.black,
                      ),
                    ),
                    Padding(
                      padding: EdgeInsets.only(
                          top: 10, bottom: 10, left: 20, right: 20),
                      child: Text(
                        "test@gmail.com",
                        style: TextStyle(color: Colors.black, fontSize: 40),
                      ),
                    ),
                  ],
                ),
              ),
            )
          ],
        ),
      ),
      backgroundColor: Colors.black,
    ),
  ));
}
