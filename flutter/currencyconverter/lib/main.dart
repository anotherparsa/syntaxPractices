import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    debugShowCheckedModeBanner: false,
    home: Scaffold(
      appBar: AppBar(
        title: Text(
          "Currency Converter",
          style: TextStyle(color: Colors.white, fontSize: 20),
        ),
        backgroundColor: Colors.black,
        centerTitle: true,
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              "0",
              style: TextStyle(color: Colors.white, fontSize: 30),
            ),
            TextField(
              style: TextStyle(color: Colors.black),
              decoration: InputDecoration(
                hintText: "Please enter the amount in IRR",
                hintStyle: TextStyle(color: Colors.black),
                filled: true,
                fillColor: Colors.white,
                prefixIcon: Icon(Icons.monetization_on),
                prefixIconColor: Colors.black,
              ),
              keyboardType: TextInputType.numberWithOptions(
                decimal: true,
              ),
            )
          ],
        ),
      ),
      backgroundColor: Colors.black,
    ),
  ));
}
