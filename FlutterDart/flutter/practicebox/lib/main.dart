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
          child: Column(
        children: [
          SizedBox(
            height: 30,
          ),
          ElevatedButton(
            onPressed: () {},
            child: Text("Press Me"),
          ),
          SizedBox(
            height: 30,
          ),
          FilledButton(
            onPressed: () {},
            child: Text("Press Me"),
          ),
          SizedBox(
            height: 30,
          ),
          FilledButton.tonal(
            onPressed: () {},
            child: Text("Press Me"),
          ),
          SizedBox(
            height: 30,
          ),
          OutlinedButton(
            onPressed: () {},
            child: Text("Press Me"),
          ),
          SizedBox(
            height: 30,
          ),
          TextButton(
            onPressed: () {},
            child: Text("Press Me"),
          )
        ],
      )),
      backgroundColor: Colors.white,
    ),
  ));
}
