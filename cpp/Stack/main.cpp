#include <iostream>

using namespace std;

class Node{
    public:
        int value;
        Node* next;

        Node(int value){
            this->value = value;
            this->next = nullptr;
        }
};

class Stack{
    private:
        Node* top;
        Node* bottom;
        int length;

    public:
        Stack(){
            this->top = nullptr;
            this->bottom = nullptr;
            this->length = 0;
        }

        void start(){
            while(true){
                cout << "Stack management center: " << endl << 
                "choose an option below: " << endl << 
                "1-print top     2-print bottom     3-print length     4-print stack " << endl << 
                "5-print specific Node value by index     6-set specific Node value     7-push a Node     8-pop a Node" << endl;
                int input;
                cin >> input;
                switch (input){
                    case 1 : this->print_head(); break;

                    default : cout << "No valid option"; break;

                }
                cout << "-----------------------------------------------------" << endl;
            }
        }

        bool is_stack_empty(){
            return (this->top == nullptr || this->length == 0)
        }

        void print_top(){
            if (this->is_stack_empty()){
                cout << "There are no Nodes in the stack" << endl;
            }else{
                cout << "Top address: " << this->top << " Top value: " << this->top->value << endl;
            }
        }

        void print_bottom(){
            if (this->is_stack_empty()){
                cout << "THere are no Nodes in the stack" << endl;
            }else{
                cout << "Bottom address: " << this->bottom << " Bottom value: " << this->bottom->value << endl;
            }
        }

        
};



